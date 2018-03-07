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
		if r == '.' {
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
	var _typeOrder, _predOrder, _methodOrder []string

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
					_typeOrder = append(_typeOrder, _type)
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
package proto;
import "google/protobuf/any.proto";`)


	// types
	fmt.Fprintf(toProto, "\n\n// Types\n")

	var allTypes []string
	for key, _ := range _nonPredTypes {
		allTypes = append(allTypes, key)
	}
	allTypes = append(_typeOrder, allTypes...)
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
	printMessage := func(title string, c constructor) {
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
					//fmt.Fprintf(toProto, "\trepeated ")
					name := t.name
					if name == subType {
						name = "values"
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
			printMessage(fmt.Sprintf("Pred%s", strings.Title(c.predicate)), c)
		} else {
			fmt.Fprintf(os.Stderr, "no predicate:", key)
		}
		fmt.Fprint(toProto, "}\n\n")
	}

	fmt.Fprintf(toProto, "\n\n// Requests\n")
	for _, key := range _methodOrder {
		c, ok := _methods[key]
		if ok {
			printMessage(fmt.Sprintf("Req%s", strings.Title(c.predicate)), c)
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
	fmt.Fprint(toGo, "package proto\nimport \"fmt\"\nconst (\n")
	for _, key := range _predOrder {
		c := _preds[key]
		fmt.Fprintf(toGo, "crc_%s = %s\n", c.predicate, c.id)
	}
	fmt.Fprint(toGo, ")\n\n")

	// encode funcs for types
	fmt.Fprintf(toGo, "\n\n// Encode funcs for types\n")

	// XXX: treat only types from predicates (ignore vector types)
	for _, key := range _typeOrder {
		// exclude Vector t
		if key == "Vector t" {
			continue
		}
		preds := _predTypes[key]
		fmt.Fprintf(toGo, "func (e Type%s) encode() []byte {\n", strings.Title(key))
		switch len(preds) {
		case 0:
		case 1:
			fmt.Fprintf(toGo, "\treturn e.GetValue().encode()\n")
		default:
			fmt.Fprintf(toGo, "\tif e.GetValue() != nil {\n\t\treturn e.Get%s().encode()\n", strings.Title(preds[0]))
			for _, p := range preds {
				fmt.Fprintf(toGo, "\t} else if e.GetValue() != nil {\n\t\treturn e.Get%s().encode()\n", strings.Title(p))
			}
			fmt.Fprintln(toGo, "\t}\n\treturn nil\n")
		}
		fmt.Fprintln(toGo, "}")
	}

	// encode funcs for predicates
	fmt.Fprintf(toGo, "\n\n// Encode funcs for predicates\n")
	for _, key := range _predOrder {
		c := _preds[key]
		fmt.Fprintf(toGo, "func (e Pred%s) encode() []byte {\n", strings.Title(c.predicate))
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
					fmt.Fprintf(toGo, "x.Bytes(e.%s.encode())\n", strings.Title(t.name))
				case "Vector<double>":
					panic(fmt.Sprintf("Unsupported %s", subType))
				default:
					var inner string
					n, _ := fmt.Sscanf(subType, "Vector<%s", &inner)
					if n == 1 {
						fmt.Fprintf(toGo, "x.Vector(e.%s)\n", strings.Title(t.name))
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
					fmt.Fprintf(toGo, "x.Bytes(e.%s.encode())\n", strings.Title(t.name))
				case "Vector<double>":
					panic(fmt.Sprintf("Unsupported %s", t._type))
				default:
					var inner string
					n, _ := fmt.Sscanf(t._type, "Vector<%s", &inner)
					if n == 1 {
						fmt.Fprintf(toGo, "x.Vector(e.%s)\n", strings.Title(t.name))
					} else {
						fmt.Fprintf(toGo, "x.Bytes(e.%s.encode())\n", strings.Title(t.name))
					}
				}
			}

		}
		fmt.Fprint(toGo, "return x.buf\n")
		fmt.Fprint(toGo, "}\n\n")
	}



	//TOOD: for now, it does NOT support decoding to Vector
	// decode funcs
//	fmt.Println(`
//func (m *DecodeBuf) ObjectGenerated(constructor uint32) (r TL) {
//	switch constructor {`)
//
//	for _, key := range _predOrder {
//		c := _preds[key]
//		fmt.Printf("case crc_%s:\n", c.predicate)
//		for _, t := range c.params {
//			if t._type == "#" {
//				fmt.Print("flags := m.Flags()\n")
//				fmt.Print("_ = flags\n")
//				break
//			}
//		}
//		fmt.Printf("r = TL_%s{\n", c.predicate)
//		for _, t := range c.params {
//			t.name = strings.Title(t.name)
//			if strings.HasPrefix(t._type, "flags") {
//				flagBit, _ := strconv.Atoi(string(t._type[strings.Index(t._type, "_") + 1:strings.Index(t._type, "?")]))
//				subType := string(t._type[strings.Index(t._type, "?") + 1:])
//				switch subType {
//				case "true":
//				case "int":
//					fmt.Print(fmt.Sprintf("m.FlaggedInt(flags, %d),\n", flagBit))
//				case "long":
//					fmt.Print(fmt.Sprintf("m.FlaggedLong(flags, %d),\n", flagBit))
//				case "string":
//					fmt.Print(fmt.Sprintf("m.FlaggedString(flags, %d),\n", flagBit))
//				case "double":
//					fmt.Print(fmt.Sprintf("m.FlaggedDouble(flags, %d),\n", flagBit))
//				case "bytes":
//					fmt.Print(fmt.Sprintf("m.FlaggedStringBytes(flags, %d),\n", flagBit))
//				case "Vector<int>":
//					fmt.Print(fmt.Sprintf("m.FlaggedVectorInt(flags, %d),\n", flagBit))
//				case "Vector<long>":
//					fmt.Print(fmt.Sprintf("m.FlaggedVectorLong(flags, %d),\n", flagBit))
//				case "Vector<string>":
//					fmt.Print(fmt.Sprintf("m.FlaggedVectorString(flags, %d),\n", flagBit))
//					fmt.Print("m.VectorString(),\n")
//				case "!X":
//					fmt.Print(fmt.Sprintf("m.FlaggedObject(flags, %d),\n", flagBit))
//				case "Vector<double>":
//					panic(fmt.Sprintf("Unsupported %s", subType))
//				default:
//					var inner string
//					n, _ := fmt.Sscanf(subType, "Vector<%s", &inner)
//					if n == 1 {
//						fmt.Print(fmt.Sprintf("m.FlaggedVector(flags, %d),\n", flagBit))
//					} else {
//						fmt.Print(fmt.Sprintf("m.FlaggedObject(flags, %d),\n", flagBit))
//					}
//				}
//			} else {
//				switch t._type {
//				case "#":
//					fmt.Print("flags,\n")
//					//fmt.Print("m.Flags(),\n")
//				case "int":
//					fmt.Print("m.Int(),\n")
//				case "long":
//					fmt.Print("m.Long(),\n")
//				case "string":
//					fmt.Print("m.String(),\n")
//				case "double":
//					fmt.Print("m.Double(),\n")
//				case "bytes":
//					fmt.Print("m.StringBytes(),\n")
//				case "Vector<int>":
//					fmt.Print("m.VectorInt(),\n")
//				case "Vector<long>":
//					fmt.Print("m.VectorLong(),\n")
//				case "Vector<string>":
//					fmt.Print("m.VectorString(),\n")
//				case "!X":
//					fmt.Print("m.Object(),\n")
//				case "Vector<double>":
//					panic(fmt.Sprintf("Unsupported %s", t._type))
//				default:
//					var inner string
//					n, _ := fmt.Sscanf(t._type, "Vector<%s", &inner)
//					if n == 1 {
//						fmt.Print("m.Vector(),\n")
//					} else {
//						fmt.Print("m.Object(),\n")
//					}
//				}
//			}
//		}
//		fmt.Print("}\n\n")
//	}

	toGo.Flush()
	goFp.Close()

//	fmt.Println(`
//	default:
//		m.err = fmt.Errorf("Unknown constructor: \u002508x", constructor)
//		return nil
//
//	}
//
//	if m.err != nil {
//		return nil
//	}
//
//	return
//}`)

}
