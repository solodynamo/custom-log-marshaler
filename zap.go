package main

import (
	"fmt"
	"strings"
)

type UberZap struct {
}

var _ PIIMarshaler = &UberZap{}

func (uz *UberZap) Generate(structs []string, structFields map[string][]field) string {
	if len(structs) == 0 {
		return ""
	}

	contents := make([]string, 0, len(structs))
	for _, st := range structs {
		content := uz.generateStructData(st, structFields)
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

func (uz *UberZap) generateStructData(structName string, structFields map[string][]field) string {
	if len(structFields[structName]) == 0 {
		return ""
	}

	format := `// MarshalLogObject ...
func (l %s) MarshalLogObject(enc zapcore.ObjectEncoder) error {
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

// https://github.com/uber-go/zap/blob/master/zapcore/encoder.go#L349
func (uz *UberZap) GetLibFunc(typeName string) string {
	switch typeName {
	case "zapcore.ObjectMarshaler":
		return "AddObject"
	case "zapcore.ArrayMarshaler":
		return "AddArray"
	case "bool":
		return "AddBool"
	case "*bool":
		return "AddBool"
	case "complex128":
		return "AddComplex128"
	case "*complex128":
		return "AddComplex128"
	case "complex64":
		return "AddComplex64"
	case "*complex64":
		return "AddComplex64"
	case "float64":
		return "AddFloat64"
	case "*float64":
		return "AddFloat64"
	case "float32":
		return "AddFloat32"
	case "*float32":
		return "AddFloat32"
	case "int":
		return "AddInt"
	case "*int":
		return "AddInt"
	case "int64":
		return "AddInt64"
	case "*int64":
		return "AddInt64"
	case "int32":
		return "AddInt32"
	case "*int32":
		return "AddInt32"
	case "int16":
		return "AddInt16"
	case "*int16":
		return "AddInt16"
	case "int8":
		return "AddInt8"
	case "*int8":
		return "AddInt8"
	case "string":
		return "AddString"
	case "*string":
		return "AddString"
	case "uint":
		return "AddUint"
	case "*uint":
		return "AddUint"
	case "uint64":
		return "AddUint64"
	case "*uint64":
		return "AddUint64"
	case "uint32":
		return "AddUint32"
	case "*uint32":
		return "AddUint32"
	case "uint16":
		return "AddUint16"
	case "*uint16":
		return "AddUint16"
	case "uint8":
		return "AddUint8"
	case "*uint8":
		return "AddUint8"
	case "[]byte":
		return "AddBinary"
	case "uintptr":
		return "AddUintptr"
	case "*uintptr":
		return "AddUintptr"
	case "time.Time":
		return "AddTime"
	case "*time.Time":
		return "AddTime"
	case "time.Duration":
		return "AddDuration"
	case "*time.Duration":
		return "AddDuration"
	default:
		return "AddReflected"
	}
}
