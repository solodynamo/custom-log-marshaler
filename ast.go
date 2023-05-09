package main

import (
	"go/ast"
	"go/token"
	"strings"

	"github.com/iancoleman/strcase"
)

const (
	cannotDetermine     = "cannot_determine"
	customStruct        = "custom_struct"
	customMap           = "custom_map"
	customPrimitiveType = "custom_primitive_type"
)

var (
	primitiveTypes = []string{"bool", "byte", "complex64", "complex128", "error", "float32", "float64", "int", "int8", "int16", "int32", "int64", "rune", "string", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr"}
)

func getFields(f *ast.File, loglib PIIMarshaler) ([]string, map[string][]field) {
	structs := make([]string, 0)
	structFields := make(map[string][]field)

	ast.Inspect(f, func(n ast.Node) bool {
		if v, ok := n.(*ast.GenDecl); ok {
			for _, s := range v.Specs {
				typeSpec, ok := s.(*ast.TypeSpec)
				if !ok {
					continue
				}

				structType, ok := typeSpec.Type.(*ast.StructType)
				if !ok {
					continue
				}
				structName := typeSpec.Name.String()
				structs = append(structs, structName)
				fields := make([]field, 0, len(structType.Fields.List))

				for i := 0; i < len(structType.Fields.List); i++ {
					fi := structType.Fields.List[i]
					fie := field{}
					if fi.Tag != nil && fi.Tag.Value != "" && strings.Contains(fi.Tag.Value, "notloggable") {
						continue

					}

					for _, ind := range fi.Names {
						fie.key = strcase.ToSnake(ind.Name)
						fie.fieldName = ind.Name
					}

					if len(fi.Names) == 0 {
						// This is an embedded field
						ident, ok := fi.Type.(*ast.Ident)
						if ok {
							fie.fieldName = ident.Name
						}
					}

					// Type
					if fi.Type != nil {
						fie.typeName = cannotDetermine

						if ty, ok := fi.Type.(*ast.Ident); ok {
							if ok, goType := isPrimitiveType(ty); ok {
								fie.typeName = goType
								goto register
							}

							if isCustomMap(f, ty.Name) {
								fie.typeName = customMap
								goto register
							}

							if ok, _ := isCustomPrimitive(f, ty.Name); ok {
								// force reflect custom defined primitive types
								fie.typeName = cannotDetermine
								goto register
							}

							if isCustomStruct(n, ty.Name) {
								fie.typeName = customStruct
								// make sure string in log fn doesn't cuse custom_struct string
								// uses the Actual field/struct name
								fie.key = ty.Name
								goto register
							}

						}

						if ty, ok := fi.Type.(*ast.StarExpr); ok {
							fie.fieldType = ptr
							if x, ok := ty.X.(*ast.Ident); ok {
								fie.typeName = x.Name
							}

							if cty, ok := ty.X.(*ast.SelectorExpr); ok {
								if x, ok := cty.X.(*ast.Ident); ok {
									fie.pkgName = x.Name
								}

								fie.typeName = cty.Sel.Name
							}
							goto register
						}

						if ty, ok := fi.Type.(*ast.SelectorExpr); ok {
							if x, ok := ty.X.(*ast.Ident); ok {
								fie.pkgName = x.Name
							}

							fie.typeName = ty.Sel.Name
						}
					}
				register:
					fie.libFunc = loglib.GetLibFunc(fie.allTypeName())
					fields = append(fields, fie)
				}
				structFields[structName] = fields
			}
		}

		return true
	})

	return structs, structFields
}

func isCustomStruct(n ast.Node, typeName string) bool {
	found := false
	switch x := n.(type) {
	case *ast.TypeSpec:
		if x.Name.Name == "User" {
			if _, ok := x.Type.(*ast.StructType); ok {
				found = true
			}
		}
	case *ast.StructType:
		for _, field := range x.Fields.List {
			if ident, ok := field.Type.(*ast.Ident); ok && ident.Name == "User" {
				found = true
			}
		}
	}
	return !found
}

func isCustomMap(f *ast.File, typeName string) bool {
	for _, decl := range f.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			continue
		}

		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if ok && typeSpec.Name.Name == typeName {
				_, isMap := typeSpec.Type.(*ast.MapType)
				return isMap
			}
		}
	}

	return false
}

func isCustomPrimitive(f *ast.File, typeName string) (bool, string) {

	for _, decl := range f.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			continue
		}

		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if ok && typeSpec.Name.Name == typeName {
				ident, isIdent := typeSpec.Type.(*ast.Ident)
				if isIdent {
					for _, primitiveType := range primitiveTypes {
						if ident.Name == primitiveType {
							return true, primitiveType
						}
					}
				}
				break
			}
		}
	}

	return false, ""
}

func isPrimitiveType(ident *ast.Ident) (bool, string) {
	for _, t := range primitiveTypes {
		if ident.Name == t {
			return true, t
		}
	}

	return false, ""
}
