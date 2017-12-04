package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type nametype struct {
	name  string
	_type string
}

type constuctor struct {
	id        string
	predicate string
	params    []nametype
	_type     string
}

func normalize(s string) string {
	x := []byte(s)
	for i, r := range x {
		if r == '.' {
			x[i] = '_'
		}
	}
	y := string(x)
	if y == "type" {
		return "_Type"
	}
	return y
}

func main() {
	var err error
	var parsed interface{}

	// read json file from stdin
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}

	// parse json
	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()
	err = d.Decode(&parsed)
	if err != nil {
		fmt.Println(err)
		return
	}

	// process constructors
	_order := make([]string, 0, 1000)
	_cons := make(map[string]constuctor, 1000)
	_types := make(map[string][]string, 1000)

	parsefunc := func(data []interface{}, kind string) {
		for _, data := range data {
			data := data.(map[string]interface{})

			// id
			//idx, err := data["id"].(json.Number).Int64()
			idx, err := strconv.Atoi(data["id"].(string))
			if err != nil {
				fmt.Println(err)
				return
			}
			_id := fmt.Sprintf("0x%08x", uint32(idx))

			// predicate
			_predicate := normalize(data[kind].(string))

			if _predicate == "vector" {
				continue
			}

			// params
			_params := make([]nametype, 0, 16)
			params := data["params"].([]interface{})
			for _, params := range params {
				params := params.(map[string]interface{})
				_params = append(_params, nametype{normalize(params["name"].(string)), normalize(params["type"].(string))})
			}

			// type
			_type := normalize(data["type"].(string))

			_order = append(_order, _predicate)
			_cons[_predicate] = constuctor{_id, _predicate, _params, _type}
			if kind == "predicate" {
				_types[_type] = append(_types[_type], _predicate)
			}
		}
	}
	parsefunc(parsed.(map[string]interface{})["constructors"].([]interface{}), "predicate")
	parsefunc(parsed.(map[string]interface{})["methods"].([]interface{}), "method")

	// constants
	fmt.Print("package mtproto\nimport \"fmt\"\nconst (\n")
	for _, key := range _order {
		c := _cons[key]
		fmt.Printf("crc_%s = %s\n", c.predicate, c.id)
	}
	fmt.Print(")\n\n")

	// type structs
	for _, key := range _order {
		c := _cons[key]
		fmt.Printf("type TL_%s struct {\n", c.predicate)
		for _, t := range c.params {
			t.name = strings.Title(t.name)
			//fmt.Printf("%s\t", t.name)
			if strings.HasPrefix(t._type, "flags") {
				subType := string(t._type[strings.Index(t._type, "?") + 1:])
				switch subType {
				case "int":
					fmt.Printf("%s\tint32", t.name)
				//fmt.Print("int32")
				case "true":
					fmt.Printf("// %s\tbool // %s", t.name, t._type)
				case "long":
					fmt.Printf("%s\tint64", t.name)
				//fmt.Print("int64")
				case "string":
					fmt.Printf("%s\tstring", t.name)
				//fmt.Print("string")
				case "double":
					fmt.Printf("%s\tfloat64", t.name)
				//fmt.Print("float64")
				case "bytes":
					fmt.Printf("%s\t[]byte", t.name)
				//fmt.Print("[]byte")
				case "Vector<int>":
					fmt.Printf("%s\t[]int32", t.name)
				//fmt.Print("[]int32")
				case "Vector<long>":
					fmt.Printf("%s\t[]int64", t.name)
				//fmt.Print("[]int64")
				case "Vector<string>":
					fmt.Printf("%s\t[]string", t.name)
				//fmt.Print("[]string")
				case "Vector<double>":
					fmt.Printf("%s\t[]float64", t.name)
				//fmt.Print("[]float64")
				case "!X":
					fmt.Printf("%s\tTL", t.name)
				//fmt.Print("TL")
				default:
					var inner string
					n, _ := fmt.Sscanf(subType, "Vector<%s", &inner)
					if n == 1 {
						fmt.Printf("%s\t[]TL // %s", t.name, inner[:len(inner) - 1])
						//fmt.Printf("[]TL // %s", inner[:len(inner) - 1])
					} else {

						fmt.Printf("%s\tTL // %s", t.name, t._type)
					}
				}
			} else {
				switch t._type {
				case "int", "#":
					fmt.Printf("%s\tint32", t.name)
					//fmt.Print("int32")
				case "long":
					fmt.Printf("%s\tint64", t.name)
					//fmt.Print("int64")
				case "string":
					fmt.Printf("%s\tstring", t.name)
					//fmt.Print("string")
				case "double":
					fmt.Printf("%s\tfloat64", t.name)
					//fmt.Print("float64")
				case "bytes":
					fmt.Printf("%s\t[]byte", t.name)
					//fmt.Print("[]byte")
				case "Vector<int>":
					fmt.Printf("%s\t[]int32", t.name)
					//fmt.Print("[]int32")
				case "Vector<long>":
					fmt.Printf("%s\t[]int64", t.name)
					//fmt.Print("[]int64")
				case "Vector<string>":
					fmt.Printf("%s\t[]string", t.name)
					//fmt.Print("[]string")
				case "Vector<double>":
					fmt.Printf("%s\t[]float64", t.name)
					//fmt.Print("[]float64")
				case "!X":
					fmt.Printf("%s\tTL", t.name)
					//fmt.Print("TL")
				default:
					var inner string
					n, _ := fmt.Sscanf(t._type, "Vector<%s", &inner)
					if n == 1 {
						fmt.Printf("%s\t[]TL // %s", t.name, inner[:len(inner) - 1])
					} else {
						fmt.Printf("%s\tTL // %s", t.name, t._type)
					}
				}
			}

			fmt.Print("\n")
		}
		fmt.Print("}\n\n")
	}

	// encode funcs
	for _, key := range _order {
		c := _cons[key]
		fmt.Printf("func (e TL_%s) encode() []byte {\n", c.predicate)
		fmt.Print("x := NewEncodeBuf(512)\n")
		fmt.Printf("x.UInt(crc_%s)\n", c.predicate)
		for _, t := range c.params {
			t.name = strings.Title(t.name)
			if strings.HasPrefix(t._type, "flags") {
				subType := string(t._type[strings.Index(t._type, "?") + 1:])
				switch subType {
				case "true":
				case "int":
					fmt.Printf("x.Int(e.%s)\n", t.name)
				case "long":
					fmt.Printf("x.Long(e.%s)\n", t.name)
				case "string":
					fmt.Printf("x.String(e.%s)\n", t.name)
				case "double":
					fmt.Printf("x.Double(e.%s)\n", t.name)
				case "bytes":
					fmt.Printf("x.StringBytes(e.%s)\n", t.name)
				case "Vector<int>":
					fmt.Printf("x.VectorInt(e.%s)\n", t.name)
				case "Vector<long>":
					fmt.Printf("x.VectorLong(e.%s)\n", t.name)
				case "Vector<string>":
					fmt.Printf("x.VectorString(e.%s)\n", t.name)
				case "!X":
					fmt.Printf("x.Bytes(e.%s.encode())\n", t.name)
				case "Vector<double>":
					panic(fmt.Sprintf("Unsupported %s", subType))
				default:
					var inner string
					n, _ := fmt.Sscanf(subType, "Vector<%s", &inner)
					if n == 1 {
						fmt.Printf("x.Vector(e.%s)\n", t.name)
					} else {
						fmt.Printf("x.Bytes(e.%s.encode())\n", t.name)
					}
				}
			} else {
				switch t._type {
				case "int", "#":
					fmt.Printf("x.Int(e.%s)\n", t.name)
				case "long":
					fmt.Printf("x.Long(e.%s)\n", t.name)
				case "string":
					fmt.Printf("x.String(e.%s)\n", t.name)
				case "double":
					fmt.Printf("x.Double(e.%s)\n", t.name)
				case "bytes":
					fmt.Printf("x.StringBytes(e.%s)\n", t.name)
				case "Vector<int>":
					fmt.Printf("x.VectorInt(e.%s)\n", t.name)
				case "Vector<long>":
					fmt.Printf("x.VectorLong(e.%s)\n", t.name)
				case "Vector<string>":
					fmt.Printf("x.VectorString(e.%s)\n", t.name)
				case "!X":
					fmt.Printf("x.Bytes(e.%s.encode())\n", t.name)
				case "Vector<double>":
					panic(fmt.Sprintf("Unsupported %s", t._type))
				default:
					var inner string
					n, _ := fmt.Sscanf(t._type, "Vector<%s", &inner)
					if n == 1 {
						fmt.Printf("x.Vector(e.%s)\n", t.name)
					} else {
						fmt.Printf("x.Bytes(e.%s.encode())\n", t.name)
					}
				}
			}

		}
		fmt.Print("return x.buf\n")
		fmt.Print("}\n\n")

	}

	// decode funcs
	fmt.Println(`
func (m *DecodeBuf) ObjectGenerated(constructor uint32) (r TL) {
	switch constructor {`)

	for _, key := range _order {
		c := _cons[key]
		fmt.Printf("case crc_%s:\n", c.predicate)
		for _, t := range c.params {
			if t._type == "#" {
				fmt.Print("flags := m.Flags()\n")
				fmt.Print("_ = flags\n")
				break
			}
		}
		fmt.Printf("r = TL_%s{\n", c.predicate)
		for _, t := range c.params {
			t.name = strings.Title(t.name)
			if strings.HasPrefix(t._type, "flags") {
				flagBit, _ := strconv.Atoi(string(t._type[strings.Index(t._type, "_") + 1:strings.Index(t._type, "?")]))
				subType := string(t._type[strings.Index(t._type, "?") + 1:])
				switch subType {
				case "true":
				case "int":
					fmt.Print(fmt.Sprintf("m.FlaggedInt(flags, %d),\n", flagBit))
				case "long":
					fmt.Print(fmt.Sprintf("m.FlaggedLong(flags, %d),\n", flagBit))
				case "string":
					fmt.Print(fmt.Sprintf("m.FlaggedString(flags, %d),\n", flagBit))
				case "double":
					fmt.Print(fmt.Sprintf("m.FlaggedDouble(flags, %d),\n", flagBit))
				case "bytes":
					fmt.Print(fmt.Sprintf("m.FlaggedStringBytes(flags, %d),\n", flagBit))
				case "Vector<int>":
					fmt.Print(fmt.Sprintf("m.FlaggedVectorInt(flags, %d),\n", flagBit))
				case "Vector<long>":
					fmt.Print(fmt.Sprintf("m.FlaggedVectorLong(flags, %d),\n", flagBit))
				case "Vector<string>":
					fmt.Print(fmt.Sprintf("m.FlaggedVectorString(flags, %d),\n", flagBit))
					fmt.Print("m.VectorString(),\n")
				case "!X":
					fmt.Print(fmt.Sprintf("m.FlaggedObject(flags, %d),\n", flagBit))
				case "Vector<double>":
					panic(fmt.Sprintf("Unsupported %s", subType))
				default:
					var inner string
					n, _ := fmt.Sscanf(subType, "Vector<%s", &inner)
					if n == 1 {
						fmt.Print(fmt.Sprintf("m.FlaggedVector(flags, %d),\n", flagBit))
					} else {
						fmt.Print(fmt.Sprintf("m.FlaggedObject(flags, %d),\n", flagBit))
					}
				}
			} else {
				switch t._type {
				case "#":
					fmt.Print("flags,\n")
					//fmt.Print("m.Flags(),\n")
				case "int":
					fmt.Print("m.Int(),\n")
				case "long":
					fmt.Print("m.Long(),\n")
				case "string":
					fmt.Print("m.String(),\n")
				case "double":
					fmt.Print("m.Double(),\n")
				case "bytes":
					fmt.Print("m.StringBytes(),\n")
				case "Vector<int>":
					fmt.Print("m.VectorInt(),\n")
				case "Vector<long>":
					fmt.Print("m.VectorLong(),\n")
				case "Vector<string>":
					fmt.Print("m.VectorString(),\n")
				case "!X":
					fmt.Print("m.Object(),\n")
				case "Vector<double>":
					panic(fmt.Sprintf("Unsupported %s", t._type))
				default:
					var inner string
					n, _ := fmt.Sscanf(t._type, "Vector<%s", &inner)
					if n == 1 {
						fmt.Print("m.Vector(),\n")
					} else {
						fmt.Print("m.Object(),\n")
					}
				}
			}
		}
		fmt.Print("}\n\n")
	}

	fmt.Println(`
	default:
		m.err = fmt.Errorf("Unknown constructor: \u002508x", constructor)
		return nil

	}

	if m.err != nil {
		return nil
	}

	return
}`)

}
