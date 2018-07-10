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

const (
	typesFileName = "types.tl.proto"
	convsFileName = "convs.tl.go"
	procsFilename = "procs.tl.go"
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

// slice field names which needs manual field name correction
var vectorFieldNameExceptions = map[string]string {
	"BotInfo": "BotInfos",
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

	//convsFp, err := os.Create(fmt.Sprintf("%s.go", filename))
	convsFp, err := os.OpenFile(convsFileName, os.O_CREATE | os.O_WRONLY, 0666)
	if err != nil {
		fmt.Fprint(os.Stderr, "file open failure:", err)
	}
	defer convsFp.Close()
	//typesFp, err := os.Create(fmt.Sprintf("%s.proto", filename))
	typesFp, err := os.OpenFile(typesFileName, os.O_CREATE | os.O_WRONLY, 0666)
	if err != nil {
		fmt.Fprint(os.Stderr, "file open failure:", err)
	}
	defer typesFp.Close()

	procsFp, err := os.OpenFile(procsFilename, os.O_CREATE | os.O_WRONLY, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "file open failure; %v", err)
	}
	defer procsFp.Close()

	convsWriter := bufio.NewWriter(convsFp)
	typesWriter := bufio.NewWriter(typesFp)
	procsWriter := bufio.NewWriter(procsFp)

	defer func() {
		convsWriter.Flush()
		typesWriter.Flush()
		procsWriter.Flush()
		convsFp.Close()
		typesFp.Close()
		procsFp.Close()
	}()
	//typesWriter.

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
	fmt.Fprintln(typesWriter, `syntax = "proto3";
package mtproto;
import "google/protobuf/any.proto";`)


	// types in gRPC messages
	fmt.Fprintf(typesWriter, "\n\n// Types\n")
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
			fmt.Fprintf(typesWriter, "message TypeVector%s {\n", strings.Title(inner))
			switch inner {
			case "int":
				fmt.Fprintf(typesWriter, "\trepeated int32 values = 1;\n")
			case"long":
				fmt.Fprintf(typesWriter, "\trepeated int64 values = 1;\n")
			case "string":
				fmt.Fprintf(typesWriter, "\trepeated string values = 1;\n")
			case "double":
				fmt.Fprintf(typesWriter, "\trepeated double values = 1;\n")
			default:
				fmt.Fprintf(typesWriter, "\trepeated Type%s %s = 1;\n", inner, inner)
			}
		} else {
			fmt.Fprintf(typesWriter, "message Type%s {\n", strings.Title(key))
			switch len(preds) {
			case 0:
			case 1:
				fmt.Fprintf(typesWriter, "\tPred%s Value = 1;\n", strings.Title(preds[0]))
			default:
				fmt.Fprintf(typesWriter, "\toneof Value {\n")
				for i, p := range preds {
					fmt.Fprintf(typesWriter, "\t\tPred%s %s = %d;\n", strings.Title(p), strings.Title(p), i+1)
				}
				fmt.Fprintln(typesWriter, "\t}")
			}
		}
		fmt.Fprintln(typesWriter, "}")
	}

	// predicates & methods in gRPC messages
	generateGRPCmessages := func(title string, c constructor) {
		fmt.Fprintf(typesWriter, "message %s {\n", title)
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
					fmt.Fprintf(typesWriter, "\tint32 %s = %d;", strings.Title(t.name), index)
				} else {
					fmt.Fprintf(typesWriter, "//fatal: type # param is flaged: %s", t.name)
				}
			case "int":
				fmt.Fprintf(typesWriter, "\tint32 %s = %d;", strings.Title(t.name), index)
			case "true":
				if isFlag {
					fmt.Fprintf(typesWriter, "// %s\tbool // %s", strings.Title(t.name), t._type)
				} else {
					fmt.Fprintf(typesWriter, "//fatal: type true param is non-flaged: %s", t.name)
				}
			case "long":
				fmt.Fprintf(typesWriter, "\tint64 %s = %d;", strings.Title(t.name), index)
			case "string":
				fmt.Fprintf(typesWriter, "\tstring %s = %d;", strings.Title(t.name), index)
			case "double":
				fmt.Fprintf(typesWriter, "\tdouble %s = %d;", strings.Title(t.name), index)
			case "bytes":
				fmt.Fprintf(typesWriter, "\tbytes %s = %d;", strings.Title(t.name), index)
			case "Vector<int>":
				fmt.Fprintf(typesWriter, "\trepeated int32 %s = %d;", strings.Title(t.name), index)
			case "Vector<long>":
				fmt.Fprintf(typesWriter, "\trepeated int64 %s = %d;", strings.Title(t.name), index)
			case "Vector<string>":
				fmt.Fprintf(typesWriter, "\trepeated string %s = %d;", strings.Title(t.name), index)
			case "Vector<double>":
				fmt.Fprintf(typesWriter, "\trepeated double %s = %d;", strings.Title(t.name), index)
			case "!X":
				fmt.Fprintf(typesWriter, "\tgoogle.protobuf.Any %s = %d;", strings.Title(t.name), index)
			default:
				fmt.Fprintf(typesWriter, "// default: %s\n", _type)
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
						fmt.Fprintf(typesWriter, "\trepeated int32 %s = %d;", strings.Title(name), index)
					case"long":
						fmt.Fprintf(typesWriter, "\trepeated int64 %s = %d;", strings.Title(name), index)
					case "string":
						fmt.Fprintf(typesWriter, "\trepeated string %s = %d;", strings.Title(name), index)
					case "double":
						fmt.Fprintf(typesWriter, "\trepeated double %s = %d;", strings.Title(name), index)
					default:
						fmt.Fprintf(typesWriter, "\trepeated Type%s %s = %d;", strings.Title(subType), strings.Title(name), index)
					}
				} else {
					fmt.Fprintf(typesWriter, "\tType%s %s = %d;", strings.Title(_type), strings.Title(t.name), index)
				}
				if isFlag {
					fmt.Fprintf(typesWriter, "// %s", t._type)
				}
			}
			fmt.Fprint(typesWriter, "\n")
		}
	}

	fmt.Fprintf(typesWriter, "\n\n// Predicates\n")
	for _, key := range _predOrder {
		// exclude vanilla vector
		if key == "vector" {
			continue
		}
		c, ok := _preds[key]
		if ok {
			generateGRPCmessages(fmt.Sprintf("Pred%s", strings.Title(c.predicate)), c)
		} else {
			fmt.Fprintf(os.Stderr, "no predicate; %v", key)
		}
		fmt.Fprint(typesWriter, "}\n\n")
	}

	fmt.Fprintf(typesWriter, "\n\n// Requests\n")
	for _, key := range _methodOrder {
		c, ok := _methods[key]
		if ok {
			generateGRPCmessages(fmt.Sprintf("Req%s", strings.Title(c.predicate)), c)
		} else {
			fmt.Fprintf(os.Stderr, "no method; %v", key)
		}
		fmt.Fprint(typesWriter, "}\n\n")
	}

	// methods in gRPC procedures
	fmt.Fprintf(typesWriter, "\n\n// Procedures\n")
	fmt.Fprintln(typesWriter, `service Mtproto {`)
	for _, key := range _methodOrder {
		c, ok:= _methods[key]
		if ok {
			switch c._type {
			case "X":
				fmt.Fprintf(typesWriter, "\trpc %s (Req%s) returns (google.protobuf.Any) {}\n", strings.Title(c.predicate), strings.Title(c.predicate))
			default:
				var inner string
				n, _ := fmt.Sscanf(c._type, "Vector<%s", &inner)
				if n == 1 {
					inner = inner[:len(inner) - 1]
					fmt.Fprintf(typesWriter, "\trpc %s (Req%s) returns (TypeVector%s) {}\n", strings.Title(c.predicate), strings.Title(c.predicate), strings.Title(inner))
				} else {
					fmt.Fprintf(typesWriter, "\trpc %s (Req%s) returns (Type%s) {}\n", strings.Title(c.predicate), strings.Title(c.predicate), strings.Title(c._type))
				}
			}

		} else {
			fmt.Fprintf(os.Stderr, "no methods; %v", key)
		}
	}
	fmt.Fprintln(typesWriter, "}")

	// Generate convs.tl.go
	// constants
	fmt.Fprintln(convsWriter, `package mtproto
import (
"fmt"
"strings"
"github.com/golang/protobuf/ptypes"
"github.com/golang/protobuf/ptypes/any"
)
`)
	fmt.Fprintf(convsWriter, "// predicates crc\nconst (\n")
	for _, key := range _predOrder {
		c := _preds[key]
		fmt.Fprintf(convsWriter, "crc_%s = %s\n", c.predicate, c.id)
	}
	fmt.Fprint(convsWriter, ")\n\n")
	fmt.Fprintf(convsWriter, "// methods crc\nconst (\n")
	for _, key := range _methodOrder {
		c := _methods[key]
		fmt.Fprintf(convsWriter, "crc_%s = %s\n", c.predicate, c.id)
	}
	fmt.Fprint(convsWriter, ")\n\n")

	// encode funcs for types
	fmt.Fprintf(convsWriter, "\n\n// Encode funcs for types\n")

	// XXX: treat only types from predicates (ignore vector types)
	for _, key := range _predTypeOrder {
		// exclude Vector t
		if key == "Vector t" {
			continue
		}
		preds := _predTypes[key]
		fmt.Fprintf(convsWriter, "func (e *Type%s) encode() []byte {\n", strings.Title(key))
		switch len(preds) {
		case 0:
		case 1:
			fmt.Fprintf(convsWriter, "\treturn e.GetValue().encode()\n")
		default:
			fmt.Fprintf(convsWriter, "switch x := e.GetValue().(type) {\n")
			//fmt.Fprintf(convsWriter, "\tif e.GetValue() != nil {\n\t\treturn e.Get%s().encode()\n", strings.Title(preds[0]))
			for _, p := range preds {
				fmt.Fprintf(convsWriter, "case *Type%s_%s:\nreturn x.%s.encode()\n", strings.Title(key), strings.Title(p), strings.Title(p))
				//fmt.Fprintf(convsWriter, "\t} else if e.GetValue() != nil {\n\t\treturn e.Get%s().encode()\n", strings.Title(p))
			}
			fmt.Fprintf(convsWriter, "}\nreturn nil\n")
			//fmt.Fprintln(convsWriter, "\t}\n\treturn nil")
		}
		fmt.Fprintln(convsWriter, "}")
	}

	generateEncoder := func (constructorName string, c constructor) {
		fmt.Fprintf(convsWriter, "func (e *%s) encode() []byte {\n", constructorName)
		fmt.Fprint(convsWriter, "x := NewEncodeBuf(512)\n")
		fmt.Fprintf(convsWriter, "x.UInt(crc_%s)\n", c.predicate)
		for _, t := range c.params {
			t.name = strings.Title(t.name)
			if strings.HasPrefix(t._type, "flags") {
				flagBit, _ := strconv.Atoi(string(t._type[strings.Index(t._type, ".") + 1:strings.Index(t._type, "?")]))
				subType := string(t._type[strings.Index(t._type, "?") + 1:])
				switch subType {
				case "true":
				case "int":
					fmt.Fprintf(convsWriter, "x.FlaggedInt(e.Flags, %d, e.%s)\n", flagBit, strings.Title(t.name))
				case "long":
					fmt.Fprintf(convsWriter, "x.FlaggedLong(e.Flags, %d, e.%s)\n", flagBit, strings.Title(t.name))
				case "string":
					fmt.Fprintf(convsWriter, "x.FlaggedString(e.Flags, %d, e.%s)\n", flagBit, strings.Title(t.name))
				case "double":
					fmt.Fprintf(convsWriter, "x.FlaggedDouble(e.Flags, %d, e.%s)\n", flagBit, strings.Title(t.name))
				case "bytes":
					fmt.Fprintf(convsWriter, "x.FlaggedStringBytes(e.Flags, %d, e.%s)\n", flagBit, strings.Title(t.name))
				case "Vector<int>":
					fmt.Fprintf(convsWriter, "x.FlaggedVectorInt(e.Flags, %d, e.%s)\n", flagBit, strings.Title(t.name))
				case "Vector<long>":
					fmt.Fprintf(convsWriter, "x.FlaggedVectorLong(e.Flags, %d, e.%s)\n", flagBit, strings.Title(t.name))
				case "Vector<string>":
					fmt.Fprintf(convsWriter, "x.FlaggedVectorString(e.Flags, %d, e.%s)\n", flagBit, strings.Title(t.name))
				case "!X":
					fmt.Fprintf(convsWriter, "unpacked%s := unpack(e.%s)\nif unpacked%s != nil {\nx.FlaggedObject(e.Flags, %d, unpacked%s)\n}\n", strings.Title(t.name), strings.Title(t.name), strings.Title(t.name), flagBit, strings.Title(t.name))
					//fmt.Fprintf(convsWriter, "x.Bytes(unpack(e.%s).encode())\n", strings.Title(t.name))
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
						fmt.Fprintf(convsWriter, "x.FlaggedVector(e.Flags, %d, toTLslice(e.%s))\n", flagBit, strings.Title(name))
					} else {
						fmt.Fprintf(convsWriter, "x.FlaggedObject(e.Flags, %d, e.%s)\n", flagBit, strings.Title(t.name))
					}
				}
			} else {
				switch t._type {
				case "int", "#":
					fmt.Fprintf(convsWriter, "x.Int(e.%s)\n", strings.Title(t.name))
				case "long":
					fmt.Fprintf(convsWriter, "x.Long(e.%s)\n", strings.Title(t.name))
				case "string":
					fmt.Fprintf(convsWriter, "x.String(e.%s)\n", strings.Title(t.name))
				case "double":
					fmt.Fprintf(convsWriter, "x.Double(e.%s)\n", strings.Title(t.name))
				case "bytes":
					fmt.Fprintf(convsWriter, "x.StringBytes(e.%s)\n", strings.Title(t.name))
				case "Vector<int>":
					fmt.Fprintf(convsWriter, "x.VectorInt(e.%s)\n", strings.Title(t.name))
				case "Vector<long>":
					fmt.Fprintf(convsWriter, "x.VectorLong(e.%s)\n", strings.Title(t.name))
				case "Vector<string>":
					fmt.Fprintf(convsWriter, "x.VectorString(e.%s)\n", strings.Title(t.name))
				case "!X":
					//fmt.Fprintln(convsWriter, `
//splits := strings.Split(x.TypeUrl, ".")
//if len(splits) < 1 {
//	return nil
//}
//typeString := splits[len(splits) - 1]
//`)
					fmt.Fprintf(convsWriter, "unpacked%s := unpack(e.%s)\nif unpacked%s != nil {\nx.Bytes(unpacked%s.encode())\n}\n", strings.Title(t.name), strings.Title(t.name), strings.Title(t.name), strings.Title(t.name))
					//fmt.Fprintf(convsWriter, "x.Bytes(unpack(e.%s).encode())\n", strings.Title(t.name))
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
						fmt.Fprintf(convsWriter, "x.Vector(toTLslice(e.%s))\n", strings.Title(name))
					} else {
						fmt.Fprintf(convsWriter, "x.Bytes(e.%s.encode())\n", strings.Title(t.name))
					}
				}
			}

		}
		fmt.Fprint(convsWriter, "return x.buf\n")
		fmt.Fprint(convsWriter, "}\n")
	}

	// predicate encoders
	fmt.Fprintf(convsWriter, "\n\n// Encode funcs for predicates\n")
	for _, key := range _predOrder {
		if key == "vector" {
			continue
		}
		c := _preds[key]
		generateEncoder(fmt.Sprintf("Pred%s", strings.Title(c.predicate)), c)
	}

	// method encoders
	fmt.Fprintf(convsWriter, "\n\n// Encode funcs for methods\n")
	for _, key := range _methodOrder {
		c := _methods[key]
		generateEncoder(fmt.Sprintf("Req%s", strings.Title(c.predicate)), c)
	}

	// TL converters to Type
	fmt.Fprintf(convsWriter, "\n\n// TL converters to a Type\n")
	for _, key := range _predTypeOrder {
		if key == "Vector t" {
			continue
		}
		preds := _predTypes[key]
		if len(preds) == 0 {
			continue
		}
		fmt.Fprintf(convsWriter, "func toType%s(tl TL) *Type%s {\n", strings.Title(key), strings.Title(key))
		fmt.Fprintf(convsWriter, "switch x := tl.(type) {\n")
		for _, p := range preds {
			c := _preds[p]
			if len(preds) == 1 {
				fmt.Fprintf(convsWriter, "case *Pred%s:\nreturn &Type%s{Value: x}\n", strings.Title(c.predicate), strings.Title(key))
			} else {
				fmt.Fprintf(convsWriter, "case *Pred%s:\nreturn &Type%s{Value: &Type%s_%s{x}}\n", strings.Title(c.predicate), strings.Title(key), strings.Title(key), strings.Title(c.predicate))
			}
		}
		fmt.Fprintf(convsWriter, "}\nreturn nil\n}\n")
	}

	// TL slice converters to Type slice
	fmt.Fprintf(convsWriter, "\n\n// TL slice converters to a Type slice\n")
	for _, key := range _predTypeOrder {
		// exclude Vector t
		if key == "Vector t" {
			continue
		}
		preds := _predTypes[key]
		if len(preds) == 0 {
			continue
		}
		fmt.Fprintf(convsWriter, "func toType%sSlice(tlslice []TL) (converted []*Type%s) {\n", strings.Title(key), strings.Title(key))
		//fmt.Fprintf(convsWriter, "if len(tlslice) < 1 {\nreturn nil\n}\n")
		fmt.Fprintf(convsWriter, "for _, tl := range tlslice {\nswitch x := tl.(type) {\n")
		if len(preds) == 1 {
			fmt.Fprintf(convsWriter, "	case *Pred%s:\n", strings.Title(preds[0]))
			fmt.Fprintf(convsWriter, "converted = append(converted, &Type%s{Value: x})\n", strings.Title(key))
		} else {
			for _, p := range preds {
				fmt.Fprintf(convsWriter, "	case *Pred%s:\n", strings.Title(p))
				fmt.Fprintf(convsWriter, "converted = append(converted, &Type%s{Value: &Type%s_%s{x}})\n", strings.Title(key), strings.Title(key), strings.Title(p))
			}
		}
		fmt.Fprintf(convsWriter, "default:\n// invalid predicate\n}	\n}\nreturn converted\n}\n")
	}

	// predicate converters to Type
	fmt.Fprintf(convsWriter, "\n\n// predicate converters to a Type\n")
	for _, key := range _predOrder {
		if key == "vector" {
			continue
		}
		c := _preds[key]
		//fmt.Fprintf(convsWriter, "func (p *Pred%s) ToType() *Type%s {\n", strings.Title(key), strings.Title(c._type))
		fmt.Fprintf(convsWriter, "func (p *Pred%s) ToType() TL {\n", strings.Title(key))
		if len(_predTypes[c._type]) == 1 {
			fmt.Fprintf(convsWriter, "return &Type%s{Value: p}\n}\n", strings.Title(c._type))
		} else {
			fmt.Fprintf(convsWriter, "return &Type%s{Value: &Type%s_%s{p}}\n}\n", strings.Title(c._type), strings.Title(c._type), strings.Title(key))
		}
	}

	//TODO: for now, it does NOT support decoding to TypeVector... structs
	// decode funcs
	fmt.Fprintln(convsWriter, `
func (m *DecodeBuf) ObjectGenerated(constructor uint32) (r TL) {
	switch constructor {`)

	generateDecoder := func(constructorName string, c constructor) {
		fmt.Fprintf(convsWriter, "case crc_%s:\n", c.predicate)
		for _, t := range c.params {
			if t._type == "#" {
				fmt.Fprint(convsWriter, "flags := m.Flags()\n")
				fmt.Fprint(convsWriter, "_ = flags\n")
				break
			}
		}
		fmt.Fprintf(convsWriter, "r = &%s{\n", constructorName)
		for _, t := range c.params {
			t.name = strings.Title(t.name)
			if strings.HasPrefix(t._type, "flags") {
				flagBit, _ := strconv.Atoi(string(t._type[strings.Index(t._type, ".") + 1:strings.Index(t._type, "?")]))
				subType := string(t._type[strings.Index(t._type, "?") + 1:])
				switch subType {
				case "true":
				case "int":
					fmt.Fprintf(convsWriter, fmt.Sprintf("%s: m.FlaggedInt(flags, %d),\n", t.name, flagBit))
				case "long":
					fmt.Fprintf(convsWriter, fmt.Sprintf("%s: m.FlaggedLong(flags, %d),\n", t.name, flagBit))
				case "string":
					fmt.Fprintf(convsWriter, fmt.Sprintf("%s: m.FlaggedString(flags, %d),\n", t.name, flagBit))
				case "double":
					fmt.Fprintf(convsWriter, fmt.Sprintf("%s: m.FlaggedDouble(flags, %d),\n", t.name, flagBit))
				case "bytes":
					fmt.Fprintf(convsWriter, fmt.Sprintf("%s: m.FlaggedStringBytes(flags, %d),\n", t.name, flagBit))
				case "Vector<int>":
					fmt.Fprintf(convsWriter, fmt.Sprintf("%s: m.FlaggedVectorInt(flags, %d),\n", t.name, flagBit))
				case "Vector<long>":
					fmt.Fprintf(convsWriter, fmt.Sprintf("%s: m.FlaggedVectorLong(flags, %d),\n", t.name, flagBit))
				case "Vector<string>":
					fmt.Fprintf(convsWriter, fmt.Sprintf("%s: m.FlaggedVectorString(flags, %d),\n", t.name, flagBit))
					//fmt.Fprintf(convsWriter, "m.VectorString(),\n")
				case "!X":
					fmt.Fprintf(convsWriter, fmt.Sprintf("Query: Pack(m.FlaggedObject(flags, %d)),\n", flagBit))
				case "Vector<double>":
					panic(fmt.Sprintf("Unsupported %s", subType))
				default:
					var inner string
					n, _ := fmt.Sscanf(subType, "Vector<%s", &inner)
					if n == 1 {
						inner = inner[:len(inner) - 1]
						fieldName := t.name
						if corrected, ok := vectorFieldNameExceptions[fieldName]; ok {
							fieldName = corrected
						}
						fmt.Fprint(convsWriter, fmt.Sprintf("%s: toType%sSlice(m.FlaggedVector(flags, %d)),\n", fieldName, strings.Title(inner), flagBit))
					} else {
						fmt.Fprint(convsWriter, fmt.Sprintf("%s: toType%s(m.FlaggedObject(flags, %d)),\n", t.name, strings.Title(subType), flagBit))
					}
				}
			} else {
				switch t._type {
				case "#":
					fmt.Fprint(convsWriter, "Flags: flags,\n")
					//fmt.Fprint(convsWriter, "m.Flags(),\n")
				case "int":
					fmt.Fprintf(convsWriter, "%s: m.Int(),\n", t.name)
				case "long":
					fmt.Fprintf(convsWriter, "%s: m.Long(),\n", t.name)
				case "string":
					fmt.Fprintf(convsWriter, "%s: m.String(),\n", t.name)
				case "double":
					fmt.Fprintf(convsWriter, "%s: m.Double(),\n", t.name)
				case "bytes":
					fmt.Fprintf(convsWriter, "%s: m.StringBytes(),\n", t.name)
				case "Vector<int>":
					fmt.Fprintf(convsWriter, "%s: m.VectorInt(),\n", t.name)
				case "Vector<long>":
					fmt.Fprintf(convsWriter, "%s: m.VectorLong(),\n", t.name)
				case "Vector<string>":
					fmt.Fprintf(convsWriter, "%s: m.VectorString(),\n", t.name)
				case "!X":
					fmt.Fprintf(convsWriter, "Query: Pack(m.Object()),\n")
				case "Vector<double>":
					panic(fmt.Sprintf("Unsupported %s", t._type))
				default:
					var inner string
					n, _ := fmt.Sscanf(t._type, "Vector<%s", &inner)
					if n == 1 {
						inner = inner[:len(inner) - 1]
						fieldName := t.name
						if corrected, ok := vectorFieldNameExceptions[fieldName]; ok {
							fieldName = corrected
						}
						fmt.Fprintf(convsWriter, "%s: toType%sSlice(m.Vector()),\n", fieldName, strings.Title(inner))
					} else {
						fmt.Fprintf(convsWriter, "%s: toType%s(m.Object()),\n", t.name, strings.Title(t._type))
					}
				}
			}
		}
		fmt.Fprint(convsWriter, "}\n\n")

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

	fmt.Fprintln(convsWriter, `
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
	fmt.Fprintln(convsWriter, "// Packer from TypeXXX to gRPC Any")
	fmt.Fprintln(convsWriter, `func Pack(tl TL) *any.Any {
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
		fmt.Fprintf(convsWriter, "case *Type%s:\nmarshaled, err = ptypes.MarshalAny(x)\n", strings.Title(key))
	}
	fmt.Fprintf(convsWriter, "\n// methods\n")
	for _, key := range _methodOrder{
		fmt.Fprintf(convsWriter, "case *Req%s:\nmarshaled, err = ptypes.MarshalAny(x)\n", strings.Title(key))
	}
	fmt.Fprintf(convsWriter, "\n// predicates\n")
	for _, key := range _predOrder{
		if key == "vector" {
			continue
		}
		fmt.Fprintf(convsWriter, "case *Pred%s:\nmarshaled, err = ptypes.MarshalAny(x)\n", strings.Title(key))
	}
	fmt.Fprintln(convsWriter, `	}
	if err != nil {
		return nil
	}
	return marshaled
}`)

	// unpacker from gRPC Any to ReqXXX
	fmt.Fprintln(convsWriter, "// Unpacker from gRPC Any to ReqXXX")
	fmt.Fprintf(convsWriter, "func unpack(x *any.Any) TL {\n")
	fmt.Fprintln(convsWriter, `if x == nil {
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
		fmt.Fprintf(convsWriter, "case \"Req%s\":\nu := &Req%s{}\n", strings.Title(c.predicate), strings.Title(c.predicate))
		fmt.Fprintln(convsWriter, `err := ptypes.UnmarshalAny(x, u)
if err != nil {
	return nil
}
return u`)
	}
	fmt.Fprintf(convsWriter, "}\nreturn nil\n}\n")


	// gRPC precedure implementations
	fmt.Fprintln(procsWriter, `package mtproto
import (
	"golang.org/x/net/context"
	"github.com/golang/protobuf/ptypes/any"
	"fmt"
)`)
	fmt.Fprintf(procsWriter, "\n\n// Procedures\n")
	for _, key := range _methodOrder {
		c, ok:= _methods[key]
		if ok {
			switch c._type {
			case "X":
				fmt.Fprintf(procsWriter, "func (caller RPCaller) %s(ctx context.Context, req *Req%s) (*any.Any, error) {\n", strings.Title(c.predicate), strings.Title(c.predicate))
				fmt.Fprintf(procsWriter, `resp, err := caller.RPC.InvokeBlocked(req)
if err != nil {
return nil, err
}
if pred, ok := resp.(Predicate); ok {
`)
				fmt.Fprintln(procsWriter, `packed := Pack(pred.ToType())
if packed != nil {
return packed, nil
}
}
return &any.Any{}, fmt.Errorf("unexpected return: %T", resp)
}`)
			default:
				var inner string
				n, _ := fmt.Sscanf(c._type, "Vector<%s", &inner)
				if n == 1 {
					inner = inner[:len(inner) - 1]
					fmt.Fprintf(procsWriter, "func (caller RPCaller) %s(ctx context.Context, req *Req%s) (*TypeVector%s, error) {\n", strings.Title(c.predicate), strings.Title(c.predicate), strings.Title(inner))
					fmt.Fprintln(procsWriter, `resp, err := caller.RPC.InvokeBlocked(req)
if err != nil {
return nil, err
}`)
					switch inner {
					case "int":
						fmt.Fprintf(procsWriter, "if resp, ok := resp.([]int32); ok {\nreturn &TypeVectorInt{Values: resp}, nil\n}\nreturn &TypeVectorInt{}, ")
					case "long":
						fmt.Fprintf(procsWriter, "if resp, ok := resp.([]int64); ok {\nreturn &TypeVectorLong{Values: resp}, nil\n}\nreturn &TypeVectorLong{}, ")
					default:
						fmt.Fprintf(procsWriter, "if resp, ok := resp.([]TL); ok {\nv := &TypeVector%s{}\nv.%s = make([]*Type%s, len(resp))\nfor i, tl := range resp {\nswitch x := tl.(type) {\n", strings.Title(inner), strings.Title(inner), strings.Title(inner))
						for _, p := range _predTypes[inner] {
							fmt.Fprintf(procsWriter, "case *Pred%s:\nv.%s[i] = x.ToType().(*Type%s)\n", strings.Title(p), strings.Title(inner), strings.Title(inner))
						}
						fmt.Fprintf(procsWriter, "default:\nreturn nil, fmt.Errorf(\"invalid %s vector element type:", strings.Title(inner))
						fmt.Fprintln(procsWriter, `%T", resp)
}
}
return v, nil
}
`)
						fmt.Fprintf(procsWriter, "return &TypeVector%s{}, ", strings.Title(inner))

					}
				} else {
					fmt.Fprintf(procsWriter, "func (caller RPCaller) %s(ctx context.Context, req *Req%s) (*Type%s, error) {\n", strings.Title(c.predicate), strings.Title(c.predicate), strings.Title(c._type))
					fmt.Fprintln(procsWriter, `resp, err := caller.RPC.InvokeBlocked(req)
if err != nil {
return nil, err
}
switch x := resp.(type) {`)
					for _, p := range _predTypes[c._type] {
						fmt.Fprintf(procsWriter, "case *Pred%s:\nreturn x.ToType().(*Type%s), nil\n", strings.Title(p), strings.Title(c._type))
					}
					fmt.Fprintf(procsWriter, "}\nreturn &Type%s{}, ", strings.Title(c._type))
				}
				fmt.Fprintln(procsWriter, `fmt.Errorf("unexpected return: %T", resp)
}`)
			}

		} else {
			fmt.Fprintf(os.Stderr, "no methods; %v", key)
		}
	}
}
