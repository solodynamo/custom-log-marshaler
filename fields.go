package main

// this code is highly inspired frrom https://github.com/muroon/zmlog, many thanks to the author.
import (
	"fmt"

	"github.com/iancoleman/strcase"
)

type fieldType int

const (
	normal fieldType = iota
	ptr
	slice
)

type field struct {
	key       string
	fieldName string
	pkgName   string
	typeName  string
	fieldType fieldType
	zapFunc   string
}

func (f field) allTypeName() string {
	n := f.typeName

	if f.pkgName != "" {
		n = fmt.Sprintf("%s.%s", f.pkgName, f.typeName)
	}

	switch f.fieldType {
	case ptr:
		n = fmt.Sprintf("*%s", n)
	case slice:
		n = "[]"
	}
	return n
}

func (f field) isEmbedded() bool {
	return f.key == "" && f.allTypeName() != ""
}

func (f field) getKey() string {
	key := ""
	if f.isEmbedded() {
		if f.pkgName != "" {
			key = fmt.Sprintf("%s.%s", f.pkgName, f.typeName)
		} else {
			key = f.typeName
		}
	} else {
		key = f.key
	}
	return strcase.ToSnake(key)
}

func (f field) getFieldName() string {
	if f.isEmbedded() {
		return f.typeName
	}
	return f.fieldName
}

func (f field) ParamValue() string {
	fieldName := fmt.Sprintf("l.%s", f.getFieldName())
	if f.fieldType == ptr {
		fieldName = fmt.Sprintf("*l.%s", f.getFieldName())
	}
	str := fmt.Sprintf("%s(\"%s\", %s)", f.zapFunc, f.getKey(), fieldName)
	return str
}

func (f field) FieldNameWithoutAestrix() string {
	fieldName := fmt.Sprintf("l.%s", f.getFieldName())
	return fieldName
}
