package main

import (
	"fmt"
	"strings"
)

type ZeroLog struct {
}

var _ PIIMarshaler = &ZeroLog{}

func (zl *ZeroLog) Generate(structs []string, structFields map[string][]field) string {
	if len(structs) == 0 {
		return ""
	}

	contents := make([]string, 0, len(structs))
	for _, st := range structs {
		content := zl.generateStructData(st, structFields)
		if content == "" {
			continue
		}

		contents = append(contents, content)
	}
	if len(contents) == 0 {
		return ""
	}

	return strings.Join(contents, "\n")
}

func (zl *ZeroLog) generateStructData(structName string, structFields map[string][]field) string {
	if len(structFields[structName]) == 0 {
		return ""
	}

	format := `// MarshalZerologObject ...
func (l %s) MarshalZerologObject(enc *zerolog.Event) error {
%s	
}
`
	indent := "\t"
	newLine := "\n"

	targetFields := make([]string, 0, len(structFields[structName])+2)

	for _, fie := range structFields[structName] {
		exe := fmt.Sprintf("enc.%s", fie.ParamValue())
		val := fmt.Sprintf("%s%s%s%s", indent, indent, exe, newLine)
		targetFields = append(targetFields, val)
	}
	suffix := fmt.Sprintf("%s%sreturn nil", indent, indent)
	targetFields = append(targetFields, suffix)

	return fmt.Sprintf(format, structName, strings.Join(targetFields, ""))
}

// https://github.com/rs/zerolog/blob/762546b5c64e03f3d23f867213e80aa45906aaf7/array.go
func (zl *ZeroLog) GetLibFunc(typeName string) string {
	switch typeName {
	case "zerolog.LogObjectMarshaler":
		return "Object"
	case "bool":
		return "Bool"
	case "*bool":
		return "Bool"
	case "float64":
		return "AddFloat64"
	case "*float64":
		return "Float64"
	case "float32":
		return "Float32"
	case "*float32":
		return "Float32"
	case "int":
		return "Int"
	case "*int":
		return "Int"
	case "int64":
		return "Int64"
	case "*int64":
		return "Int64"
	case "int32":
		return "Int32"
	case "*int32":
		return "Int32"
	case "int16":
		return "Int16"
	case "*int16":
		return "Int16"
	case "int8":
		return "Int8"
	case "*int8":
		return "Int8"
	case "string":
		return "Str"
	case "*string":
		return "Str"
	case "uint":
		return "Uint"
	case "*uint":
		return "Uint"
	case "uint64":
		return "Uint64"
	case "*uint64":
		return "Uint64"
	case "uint32":
		return "Uint32"
	case "*uint32":
		return "Uint32"
	case "uint16":
		return "Uint16"
	case "*uint16":
		return "Uint16"
	case "uint8":
		return "Uint8"
	case "*uint8":
		return "Uint8"
	case "[]byte":
		return "Bytes"
	case "time.Time":
		return "Time"
	case "*time.Time":
		return "Time"
	case "time.Duration":
		return "Dur"
	case "*time.Duration":
		return "Dur"
	default:
		return "Interface"

	}
}
