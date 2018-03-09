package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"bufio"
)

type nametype struct {
	name  string
	_type string
}

type constructor struct {
	id        string
	predicate string
	params    []nametype
	_type     string
}

func normalize(s string) string {
	var prefix string
	var x []byte
	if strings.HasPrefix(s, "flags") {
		splits := strings.Split(s, "?")
		switch len(splits) {
		case 0:
			fmt.Fprintln(os.Stderr, "no param:", s)
		case 1:
			return s
		case 2:
			prefix = fmt.Sprintf("%s?", splits[0])
			x = []byte(splits[1])
		default:
			fmt.Fprintln(os.Stderr, "too many ? in a flaged param:", s)
		}
	} else {
		x = []byte(s)
	}

	until := len(x)
	for i := 0; i < until; i++ {
		r := x[i]
		if r == '.' || r == '_' {
			if until - i > 1 {
				x = append(x[:i], strings.Title(string(x[i+1:]))...)
				until--
			} else {
				x = x[:i]
			}
		}
	}
	return fmt.Sprintf("%s%s", prefix, string(x))
}

func main() {
	var err error
	var parsed interface{}


	// set outs
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "no args")
		os.Exit(128)
	} else if len(os.Args) > 2 {
		fmt.Fprintf(os.Stderr, "too many args")
		os.Exit(128)
	}
	filename := os.Args[1]
	//goFp, err := os.Create(fmt.Sprintf("%s.go", filename))
	goFp, err := os.OpenFile(fmt.Sprintf("%s.go", filename), os.O_CREATE | os.O_WRONLY, 0666)
	if err != nil {
		fmt.Fprint(os.Stderr, "file open failure:", err)
	}
	defer goFp.Close()
	//protoFp, err := os.Create(fmt.Sprintf("%s.proto", filename))
	protoFp, err := os.OpenFile(fmt.Sprintf("%s.proto", filename), os.O_CREATE | os.O_WRONLY, 0666)
	if err != nil {
		fmt.Fprint(os.Stderr, "file open failure:", err)
	}
	defer protoFp.Close()

	toGo := bufio.NewWriter(goFp)
	toProto := bufio.NewWriter(protoFp)
	//toProto.

	// read json file from stdin
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	// parse json
	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()
	err = d.Decode(&parsed)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	// process constructors
	var _predTypeOrder, _predOrder, _methodOrder []string

	_preds := make(map[string]constructor)
	_methods := make(map[string]constructor)
	_predTypes := make(map[string][]string)
	_nonPredTypes := make(map[string]bool)

	parsefunc := func(data []interface{}, kind string) {
		for _, data := range data {
			data := data.(map[string]interface{})

			// id
			//idx, err := data["id"].(json.Number).Int64()
			idx, err := strconv.Atoi(data["id"].(string))
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			_id := fmt.Sprintf("0x%08x", uint32(idx))

			// predicate
			_predicate := normalize(data[kind].(string))

			/*if _predicate == "vector" {
				continue
			}*/

			// params
			_params := make([]nametype, 0, 16)
			params := data["params"].([]interface{})
			for _, params := range params {
				params := params.(map[string]interface{})
				_params = append(_params, nametype{normalize(params["name"].(string)), normalize(params["type"].(string))})
			}

			// type
			_type := normalize(data["type"].(string))

			switch kind {
			case "predicate":
				_predOrder = append(_predOrder, _predicate)
				_preds[_predicate] = constructor{_id, _predicate, _params, _type}
				if _, ok := _predTypes[_type]; !ok {
					_predTypeOrder = append(_predTypeOrder, _type)
				}
				_predTypes[_type] = append(_predTypes[_type], _predicate)
			case "method":
				_methodOrder = append(_methodOrder, _predicate)
				_methods[_predicate] = constructor{_id, _predicate, _params, _type}
				if _, ok := _predTypes[_type]; !ok {
					if _, ok := _nonPredTypes[_type]; !ok && _type != "X" {
						_nonPredTypes[_type] = true
					}
				}
			default:
				fmt.Fprintln(os.Stderr, "unknown kind:", kind)
			}
		}
	}
	parsefunc(parsed.(map[string]interface{})["constructors"].([]interface{}), "predicate")
	parsefunc(parsed.(map[string]interface{})["methods"].([]interface{}), "method")

	// proto meta
	fmt.Fprintln(toProto, `syntax = "proto3";
package mtp;
import "google/protobuf/any.proto";`)


	// types in gRPC messages
	fmt.Fprintf(toProto, "\n\n// Types\n")
	var allTypes []string
	for key, _ := range _nonPredTypes {
		allTypes = append(allTypes, key)
	}
	allTypes = append(_predTypeOrder, allTypes...)
	for _, key := range allTypes {
		// exclude Vector t
		if key == "Vector t" {
			continue
		}
		preds := _predTypes[key]
		var inner string
		n, _ := fmt.Sscanf(key, "Vector<%s", &inner)
		if n == 1 {
			inner = inner[:len(inner)-1]
			fmt.Fprintf(toProto, "message TypeVector%s {\n", strings.Title(inner))
			switch inner {
			case "int":
				fmt.Fprintf(toProto, "\trepeated int32 values = 1;\n")
			case"long":
				fmt.Fprintf(toProto, "\trepeated int64 values = 1;\n")
			case "string":
				fmt.Fprintf(toProto, "\trepeated string values = 1;\n")
			case "double":
				fmt.Fprintf(toProto, "\trepeated double values = 1;\n")
			default:
				fmt.Fprintf(toProto, "\trepeated Type%s %s = 1;\n", inner, inner)
			}
		} else {
			fmt.Fprintf(toProto, "message Type%s {\n", strings.Title(key))
			switch len(preds) {
			case 0:
			case 1:
				fmt.Fprintf(toProto, "\tPred%s Value = 1;\n", strings.Title(preds[0]))
			default:
				fmt.Fprintf(toProto, "\toneof Value {\n")
				for i, p := range preds {
					fmt.Fprintf(toProto, "\t\tPred%s %s = %d;\n", strings.Title(p), strings.Title(p), i+1)
				}
				fmt.Fprintln(toProto, "\t}")
			}
		}
		fmt.Fprintln(toProto, "}")
	}

	// predicates & methods in gRPC messages
	generateGRPCmessages := func(title string, c constructor) {
		fmt.Fprintf(toProto, "message %s {\n", title)
		for i, t := range c.params {
			t.name = strings.Title(t.name)
			isFlag := false
			var _type string
			if strings.HasPrefix(t._type, "flags") {
				isFlag = true
				_type = string(t._type[strings.Index(t._type, "?")+1:])
			} else {
				_type = t._type
			}

			var index= i + 1
			switch _type {
			case "#":
				if !isFlag {
					fmt.Fprintf(toProto, "\tint32 %s = %d;", strings.Title(t.name), index)
				} else {
					fmt.Fprintf(toProto, "//fatal: type # param is flaged: %s", t.name)
				}
			case "int":
				fmt.Fprintf(toProto, "\tint32 %s = %d;", strings.Title(t.name), index)
			case "true":
				if isFlag {
					fmt.Fprintf(toProto, "// %s\tbool // %s", strings.Title(t.name), t._type)
				} else {
					fmt.Fprintf(toProto, "//fatal: type true param is non-flaged: %s", t.name)
				}
			case "long":
				fmt.Fprintf(toProto, "\tint64 %s = %d;", strings.Title(t.name), index)
			case "string":
				fmt.Fprintf(toProto, "\tstring %s = %d;", strings.Title(t.name), index)
			case "double":
				fmt.Fprintf(toProto, "\tdouble %s = %d;", strings.Title(t.name), index)
			case "bytes":
				fmt.Fprintf(toProto, "\tbytes %s = %d;", strings.Title(t.name), index)
			case "Vector<int>":
				fmt.Fprintf(toProto, "\trepeated int32 %s = %d;", strings.Title(t.name), index)
			case "Vector<long>":
				fmt.Fprintf(toProto, "\trepeated int64 %s = %d;", strings.Title(t.name), index)
			case "Vector<string>":
				fmt.Fprintf(toProto, "\trepeated string %s = %d;", strings.Title(t.name), index)
			case "Vector<double>":
				fmt.Fprintf(toProto, "\trepeated double %s = %d;", strings.Title(t.name), index)
			case "!X":
				fmt.Fprintf(toProto, "\tgoogle.protobuf.Any %s = %d;", strings.Title(t.name), index)
			default:
				fmt.Fprintf(toProto, "// default: %s\n", _type)
				var inner string
				n, _ := fmt.Sscanf(_type, "Vector<%s", &inner)
				if n == 1 {
					subType := inner[:len(inner)-1]
					name := t.name
					if name == subType {
						//name = "values"
						name = fmt.Sprintf("%ss", name)
					}
					switch subType {
					case "int":
						fmt.Fprintf(toProto, "\trepeated int32 %s = %d;", strings.Title(name), index)
					case"long":
						fmt.Fprintf(toProto, "\trepeated int64 %s = %d;", strings.Title(name), index)
					case "string":
						fmt.Fprintf(toProto, "\trepeated string %s = %d;", strings.Title(name), index)
					case "double":
						fmt.Fprintf(toProto, "\trepeated double %s = %d;", strings.Title(name), index)
					default:
						fmt.Fprintf(toProto, "\trepeated Type%s %s = %d;", strings.Title(subType), strings.Title(name), index)
					}
				} else {
					fmt.Fprintf(toProto, "\tType%s %s = %d;", strings.Title(_type), strings.Title(t.name), index)
				}
				if isFlag {
					fmt.Fprintf(toProto, "// %s", t._type)
				}
			}
			fmt.Fprint(toProto, "\n")
		}
	}

	fmt.Fprintf(toProto, "\n\n// Predicates\n")
	for _, key := range _predOrder {
		// exclude vanilla vector
		if key == "vector" {
			continue
		}
		c, ok := _preds[key]
		if ok {
			generateGRPCmessages(fmt.Sprintf("Pred%s", strings.Title(c.predicate)), c)
		} else {
			fmt.Fprintf(os.Stderr, "no predicate:", key)
		}
		fmt.Fprint(toProto, "}\n\n")
	}

	fmt.Fprintf(toProto, "\n\n// Requests\n")
	for _, key := range _methodOrder {
		c, ok := _methods[key]
		if ok {
			generateGRPCmessages(fmt.Sprintf("Req%s", strings.Title(c.predicate)), c)
		} else {
			fmt.Fprintf(os.Stderr, "no method:", key)
		}
		fmt.Fprint(toProto, "}\n\n")
	}

	// methods in gRPC procedures
	fmt.Fprintf(toProto, "\n\n// Procedures\n")
	fmt.Fprintln(toProto, `service Mtproto {`)
	for _, key := range _methodOrder {
		c, ok:= _methods[key]
		if ok {
			switch c._type {
			case "X":
				fmt.Fprintf(toProto, "\trpc %s (Req%s) returns (google.protobuf.Any) {}\n", strings.Title(c.predicate), strings.Title(c.predicate))
			default:
				var inner string
				n, _ := fmt.Sscanf(c._type, "Vector<%s", &inner)
				if n == 1 {
					inner = inner[:len(inner) - 1]
					fmt.Fprintf(toProto, "\trpc %s (Req%s) returns (TypeVector%s) {}\n", strings.Title(c.predicate), strings.Title(c.predicate), strings.Title(inner))
				} else {
					fmt.Fprintf(toProto, "\trpc %s (Req%s) returns (Type%s) {}\n", strings.Title(c.predicate), strings.Title(c.predicate), strings.Title(c._type))
				}
			}

		} else {
			fmt.Fprintf(os.Stderr, "no methods:", key)
		}
	}
	fmt.Fprintln(toProto, "}")
	toProto.Flush()
	protoFp.Close()









	// Generate Go file
	// constants
	fmt.Fprintln(toGo, `package mtp
import (
"fmt"
"strings"
"github.com/golang/protobuf/ptypes"
"github.com/golang/protobuf/ptypes/any"
)
`)
	fmt.Fprintf(toGo, "// predicates crc\nconst (\n")
	for _, key := range _predOrder {
		c := _preds[key]
		fmt.Fprintf(toGo, "crc_%s = %s\n", c.predicate, c.id)
	}
	fmt.Fprint(toGo, ")\n\n")
	fmt.Fprintf(toGo, "// methods crc\nconst (\n")
	for _, key := range _methodOrder {
		c := _methods[key]
		fmt.Fprintf(toGo, "crc_%s = %s\n", c.predicate, c.id)
	}
	fmt.Fprint(toGo, ")\n\n")

	// encode funcs for types
	fmt.Fprintf(toGo, "\n\n// Encode funcs for types\n")

	// XXX: treat only types from predicates (ignore vector types)
	for _, key := range _predTypeOrder {
		// exclude Vector t
		if key == "Vector t" {
			continue
		}
		preds := _predTypes[key]
		fmt.Fprintf(toGo, "func (e *Type%s) encode() []byte {\n", strings.Title(key))
		switch len(preds) {
		case 0:
		case 1:
			fmt.Fprintf(toGo, "\treturn e.GetValue().encode()\n")
		default:
			fmt.Fprintf(toGo, "switch x := e.GetValue().(type) {\n")
			//fmt.Fprintf(toGo, "\tif e.GetValue() != nil {\n\t\treturn e.Get%s().encode()\n", strings.Title(preds[0]))
			for _, p := range preds {
				fmt.Fprintf(toGo, "case *Type%s_%s:\nreturn x.%s.encode()\n", strings.Title(key), strings.Title(p), strings.Title(p))
				//fmt.Fprintf(toGo, "\t} else if e.GetValue() != nil {\n\t\treturn e.Get%s().encode()\n", strings.Title(p))
			}
			fmt.Fprintf(toGo, "}\nreturn nil\n")
			//fmt.Fprintln(toGo, "\t}\n\treturn nil")
		}
		fmt.Fprintln(toGo, "}")
	}

	generateEncoder := func (constructorName string, c constructor) {
		fmt.Fprintf(toGo, "func (e *%s) encode() []byte {\n", constructorName)
		fmt.Fprint(toGo, "x := NewEncodeBuf(512)\n")
		fmt.Fprintf(toGo, "x.UInt(crc_%s)\n", c.predicate)
		for _, t := range c.params {
			t.name = strings.Title(t.name)
			if strings.HasPrefix(t._type, "flags") {
				subType := string(t._type[strings.Index(t._type, "?") + 1:])
				switch subType {
				case "true":
				case "int":
					fmt.Fprintf(toGo, "x.Int(e.%s)\n", strings.Title(t.name))
				case "long":
					fmt.Fprintf(toGo, "x.Long(e.%s)\n", strings.Title(t.name))
				case "string":
					fmt.Fprintf(toGo, "x.String(e.%s)\n", strings.Title(t.name))
				case "double":
					fmt.Fprintf(toGo, "x.Double(e.%s)\n", strings.Title(t.name))
				case "bytes":
					fmt.Fprintf(toGo, "x.StringBytes(e.%s)\n", strings.Title(t.name))
				case "Vector<int>":
					fmt.Fprintf(toGo, "x.VectorInt(e.%s)\n", strings.Title(t.name))
				case "Vector<long>":
					fmt.Fprintf(toGo, "x.VectorLong(e.%s)\n", strings.Title(t.name))
				case "Vector<string>":
					fmt.Fprintf(toGo, "x.VectorString(e.%s)\n", strings.Title(t.name))
				case "!X":
					fmt.Fprintf(toGo, "unpacked%s := unpack(e.%s)\nif unpacked%s != nil {\nx.Bytes(unpacked%s.encode())\n}\n", strings.Title(t.name), strings.Title(t.name), strings.Title(t.name), strings.Title(t.name))
					//fmt.Fprintf(toGo, "x.Bytes(unpack(e.%s).encode())\n", strings.Title(t.name))
				case "Vector<double>":
					panic(fmt.Sprintf("Unsupported %s", subType))
				default:
					var inner string
					n, _ := fmt.Sscanf(subType, "Vector<%s", &inner)
					if n == 1 {
						inner = inner[:len(inner) - 1]
						name := t.name
						if name == inner {
							name = fmt.Sprintf("%ss", name)
						}
						fmt.Fprintf(toGo, "x.Vector(toTLslice(e.%s))\n", strings.Title(name))
					} else {
						fmt.Fprintf(toGo, "x.Bytes(e.%s.encode())\n", strings.Title(t.name))
					}
				}
			} else {
				switch t._type {
				case "int", "#":
					fmt.Fprintf(toGo, "x.Int(e.%s)\n", strings.Title(t.name))
				case "long":
					fmt.Fprintf(toGo, "x.Long(e.%s)\n", strings.Title(t.name))
				case "string":
					fmt.Fprintf(toGo, "x.String(e.%s)\n", strings.Title(t.name))
				case "double":
					fmt.Fprintf(toGo, "x.Double(e.%s)\n", strings.Title(t.name))
				case "bytes":
					fmt.Fprintf(toGo, "x.StringBytes(e.%s)\n", strings.Title(t.name))
				case "Vector<int>":
					fmt.Fprintf(toGo, "x.VectorInt(e.%s)\n", strings.Title(t.name))
				case "Vector<long>":
					fmt.Fprintf(toGo, "x.VectorLong(e.%s)\n", strings.Title(t.name))
				case "Vector<string>":
					fmt.Fprintf(toGo, "x.VectorString(e.%s)\n", strings.Title(t.name))
				case "!X":
					//fmt.Fprintln(toGo, `
//splits := strings.Split(x.TypeUrl, ".")
//if len(splits) < 1 {
//	return nil
//}
//typeString := splits[len(splits) - 1]
//`)
					fmt.Fprintf(toGo, "unpacked%s := unpack(e.%s)\nif unpacked%s != nil {\nx.Bytes(unpacked%s.encode())\n}\n", strings.Title(t.name), strings.Title(t.name), strings.Title(t.name), strings.Title(t.name))
					//fmt.Fprintf(toGo, "x.Bytes(unpack(e.%s).encode())\n", strings.Title(t.name))
				case "Vector<double>":
					panic(fmt.Sprintf("Unsupported %s", t._type))
				default:
					var inner string
					n, _ := fmt.Sscanf(t._type, "Vector<%s", &inner)
					if n == 1 {
						inner = inner[:len(inner) - 1]
						name := t.name
						if name == inner {
							name = fmt.Sprintf("%ss", name)
						}
						fmt.Fprintf(toGo, "x.Vector(toTLslice(e.%s))\n", strings.Title(name))
					} else {
						fmt.Fprintf(toGo, "x.Bytes(e.%s.encode())\n", strings.Title(t.name))
					}
				}
			}

		}
		fmt.Fprint(toGo, "return x.buf\n")
		fmt.Fprint(toGo, "}\n")
	}

	// predicate encoders
	fmt.Fprintf(toGo, "\n\n// Encode funcs for predicates\n")
	for _, key := range _predOrder {
		if key == "vector" {
			continue
		}
		c := _preds[key]
		generateEncoder(fmt.Sprintf("Pred%s", strings.Title(c.predicate)), c)
	}

	// method encoders
	fmt.Fprintf(toGo, "\n\n// Encode funcs for methods\n")
	for _, key := range _methodOrder {
		c := _methods[key]
		generateEncoder(fmt.Sprintf("Req%s", strings.Title(c.predicate)), c)
	}

	// TL converters to Type
	fmt.Fprintf(toGo, "\n\n// TL converters to a Type\n")
	for _, key := range _predTypeOrder {
		if key == "Vector t" {
			continue
		}
		preds := _predTypes[key]
		if len(preds) == 0 {
			continue
		}
		fmt.Fprintf(toGo, "func toType%s(tl TL) *Type%s {\n", strings.Title(key), strings.Title(key))
		fmt.Fprintf(toGo, "switch x := tl.(type) {\n")
		for _, p := range preds {
			c := _preds[p]
			if len(preds) == 1 {
				fmt.Fprintf(toGo, "case *Pred%s:\nreturn &Type%s{x}\n", strings.Title(c.predicate), strings.Title(key))
			} else {
				fmt.Fprintf(toGo, "case *Pred%s:\nreturn &Type%s{&Type%s_%s{x}}\n", strings.Title(c.predicate), strings.Title(key), strings.Title(key), strings.Title(c.predicate))
			}
		}
		fmt.Fprintf(toGo, "}\nreturn nil\n}\n")
	}

	// TL slice converters to Type slice
	fmt.Fprintf(toGo, "\n\n// TL slice converters to a Type slice\n")
	for _, key := range _predTypeOrder {
		// exclude Vector t
		if key == "Vector t" {
			continue
		}
		preds := _predTypes[key]
		if len(preds) == 0 {
			continue
		}
		fmt.Fprintf(toGo, "func toType%sSlice(tlslice []TL) (converted []*Type%s) {\n", strings.Title(key), strings.Title(key))
		//fmt.Fprintf(toGo, "if len(tlslice) < 1 {\nreturn nil\n}\n")
		fmt.Fprintf(toGo, "for _, tl := range tlslice {\nswitch x := tl.(type) {\n")
		if len(preds) == 1 {
			fmt.Fprintf(toGo, "	case *Pred%s:\n", strings.Title(preds[0]))
			fmt.Fprintf(toGo, "converted = append(converted, &Type%s{x})\n", strings.Title(key))
		} else {
			for _, p := range preds {
				fmt.Fprintf(toGo, "	case *Pred%s:\n", strings.Title(p))
				fmt.Fprintf(toGo, "converted = append(converted, &Type%s{&Type%s_%s{x}})\n", strings.Title(key), strings.Title(key), strings.Title(p))
			}
		}
		fmt.Fprintf(toGo, "default:\n// invalid predicate\n}	\n}\nreturn converted\n}\n")
	}


	//TODO: for now, it does NOT support decoding to TypeVector... structs
	// decode funcs
	fmt.Fprintln(toGo, `
func (m *DecodeBuf) ObjectGenerated(constructor uint32) (r TL) {
	switch constructor {`)

	generateDecoder := func(constructorName string, c constructor) {
		fmt.Fprintf(toGo, "case crc_%s:\n", c.predicate)
		for _, t := range c.params {
			if t._type == "#" {
				fmt.Fprint(toGo, "flags := m.Flags()\n")
				fmt.Fprint(toGo, "_ = flags\n")
				break
			}
		}
		fmt.Fprintf(toGo, "r = &%s{\n", constructorName)
		for _, t := range c.params {
			t.name = strings.Title(t.name)
			if strings.HasPrefix(t._type, "flags") {
				flagBit, _ := strconv.Atoi(string(t._type[strings.Index(t._type, ".") + 1:strings.Index(t._type, "?")]))
				subType := string(t._type[strings.Index(t._type, "?") + 1:])
				switch subType {
				case "true":
				case "int":
					fmt.Fprint(toGo, fmt.Sprintf("m.FlaggedInt(flags, %d),\n", flagBit))
				case "long":
					fmt.Fprint(toGo, fmt.Sprintf("m.FlaggedLong(flags, %d),\n", flagBit))
				case "string":
					fmt.Fprint(toGo, fmt.Sprintf("m.FlaggedString(flags, %d),\n", flagBit))
				case "double":
					fmt.Fprint(toGo, fmt.Sprintf("m.FlaggedDouble(flags, %d),\n", flagBit))
				case "bytes":
					fmt.Fprint(toGo, fmt.Sprintf("m.FlaggedStringBytes(flags, %d),\n", flagBit))
				case "Vector<int>":
					fmt.Fprint(toGo, fmt.Sprintf("m.FlaggedVectorInt(flags, %d),\n", flagBit))
				case "Vector<long>":
					fmt.Fprint(toGo, fmt.Sprintf("m.FlaggedVectorLong(flags, %d),\n", flagBit))
				case "Vector<string>":
					fmt.Fprint(toGo, fmt.Sprintf("m.FlaggedVectorString(flags, %d),\n", flagBit))
					fmt.Fprint(toGo, "m.VectorString(),\n")
				case "!X":
					fmt.Fprint(toGo, fmt.Sprintf("pack(m.FlaggedObject(flags, %d)),\n", flagBit))
				case "Vector<double>":
					panic(fmt.Sprintf("Unsupported %s", subType))
				default:
					var inner string
					n, _ := fmt.Sscanf(subType, "Vector<%s", &inner)
					if n == 1 {
						inner = inner[:len(inner) - 1]
						fmt.Fprint(toGo, fmt.Sprintf("toType%sSlice(m.FlaggedVector(flags, %d)),\n", strings.Title(inner), flagBit))
					} else {
						fmt.Fprint(toGo, fmt.Sprintf("toType%s(m.FlaggedObject(flags, %d)),\n", strings.Title(subType), flagBit))
					}
				}
			} else {
				switch t._type {
				case "#":
					fmt.Fprint(toGo, "flags,\n")
					//fmt.Fprint(toGo, "m.Flags(),\n")
				case "int":
					fmt.Fprint(toGo, "m.Int(),\n")
				case "long":
					fmt.Fprint(toGo, "m.Long(),\n")
				case "string":
					fmt.Fprint(toGo, "m.String(),\n")
				case "double":
					fmt.Fprint(toGo, "m.Double(),\n")
				case "bytes":
					fmt.Fprint(toGo, "m.StringBytes(),\n")
				case "Vector<int>":
					fmt.Fprint(toGo, "m.VectorInt(),\n")
				case "Vector<long>":
					fmt.Fprint(toGo, "m.VectorLong(),\n")
				case "Vector<string>":
					fmt.Fprint(toGo, "m.VectorString(),\n")
				case "!X":
					fmt.Fprint(toGo, "pack(m.Object()),\n")
				case "Vector<double>":
					panic(fmt.Sprintf("Unsupported %s", t._type))
				default:
					var inner string
					n, _ := fmt.Sscanf(t._type, "Vector<%s", &inner)
					if n == 1 {
						inner = inner[:len(inner) - 1]
						fmt.Fprintf(toGo, "toType%sSlice(m.Vector()),\n", strings.Title(inner))
					} else {
						fmt.Fprintf(toGo, "toType%s(m.Object()),\n", strings.Title(t._type))
					}
				}
			}
		}
		fmt.Fprint(toGo, "}\n\n")

	}

	// predicate decodes
	for _, key := range _predOrder {
		if key == "vector" {
			continue
		}
		c := _preds[key]
		generateDecoder(fmt.Sprintf("Pred%s", strings.Title(c.predicate)), c)
	}

	// method decoders
	for _, key := range _methodOrder {
		if key == "vector" {
			continue
		}
		c := _methods[key]
		generateDecoder(fmt.Sprintf("Req%s", strings.Title(c.predicate)), c)
	}

	fmt.Fprintln(toGo, `
	default:
		m.err = fmt.Errorf("Unknown constructor: \u002508x", constructor)
		return nil

	}

	if m.err != nil {
		return nil
	}

	return
}`)

	// packers from types(TypeXXX) to gRPC Any
	// TODO: for now, it ignores TypeVectorXXX
	fmt.Fprintln(toGo, "// Packer from TypeXXX to gRPC Any")
	fmt.Fprintln(toGo, `func pack(tl TL) *any.Any {
	var marshaled *any.Any
	var err error
	if tl == nil {
		return nil
	}
	switch x := tl.(type) {
	// types`)
	for _, key := range _predTypeOrder {
		if key == "Vector t" {
			continue
		}
		fmt.Fprintf(toGo, "case *Type%s:\nmarshaled, err = ptypes.MarshalAny(x)\n", strings.Title(key))
	}
	fmt.Fprintf(toGo, "\n// methods\n")
	for _, key := range _methodOrder{
		fmt.Fprintf(toGo, "case *Req%s:\nmarshaled, err = ptypes.MarshalAny(x)\n", strings.Title(key))
	}
	fmt.Fprintln(toGo, `	}
	if err != nil {
		return nil
	}
	return marshaled
}`)
//	for _, key := range _predTypeOrder {
//		// exclude Vector t
//		if key == "Vector t" {
//			continue
//		}
//		//fmt.Fprintf(toGo, "func (e *Type%s) pack() *any.Any {\n", strings.Title(key))
//
//		fmt.Fprintf(toGo, `x, err := ptypes.MarshalAny(e)
//if err != nil {
//	return nil
//}
//return x
//}
//`)
//	}

	// unpacker from gRPC Any to ReqXXX
	fmt.Fprintln(toGo, "// Unpacker from gRPC Any to ReqXXX")
	fmt.Fprintf(toGo, "func unpack(x *any.Any) TL {\n")
	fmt.Fprintln(toGo, `if x == nil {
return nil
}
splits := strings.Split(x.TypeUrl, ".")
if len(splits) < 1 {
	return nil
}
typeString := splits[len(splits) - 1]
switch typeString {`)
	for _, key := range _methodOrder {
		c := _methods[key]
		fmt.Fprintf(toGo, "case \"Req%s\":\nu := &Req%s{}\n", strings.Title(c.predicate), strings.Title(c.predicate))
		fmt.Fprintln(toGo, `err := ptypes.UnmarshalAny(x, u)
if err != nil {
	return nil
}
return u`)
	}
	fmt.Fprintf(toGo, "}\nreturn nil\n}\n")

	toGo.Flush()
	goFp.Close()

}
