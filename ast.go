package main

import (
	"go/ast"
	"strings"

	"github.com/iancoleman/strcase"
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
				for _, fi := range structType.Fields.List {
					fie := field{}
					isLoggable := true
					if fi.Tag != nil && fi.Tag.Value != "" && strings.Contains(fi.Tag.Value, "notloggable") {
						isLoggable = false
						break
					}
					if !isLoggable {
						// respect the tag to not include in final logging
						continue
					}
					for _, ind := range fi.Names {
						fie.key = strcase.ToSnake(ind.Name)
						fie.fieldName = ind.Name
					}

					// Type
					if fi.Type != nil {
						if ty, ok := fi.Type.(*ast.Ident); ok {
							fie.typeName = ty.Name
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
						}

						if ty, ok := fi.Type.(*ast.ArrayType); ok {
							fie.fieldType = slice
							if x, ok := ty.Elt.(*ast.Ident); ok {
								fie.typeName = x.Name
							}

							if cty, ok := ty.Elt.(*ast.SelectorExpr); ok {
								if x, ok := cty.X.(*ast.Ident); ok {
									fie.pkgName = x.Name
								}

								fie.typeName = cty.Sel.Name
							}
						}

						if ty, ok := fi.Type.(*ast.SelectorExpr); ok {
							if x, ok := ty.X.(*ast.Ident); ok {
								fie.pkgName = x.Name
							}

							fie.typeName = ty.Sel.Name
						}

						if _, ok := fi.Type.(*ast.MapType); ok {
							// force reflection
							fie.typeName = "reflection"
						}

						if _, ok := fi.Type.(*ast.SliceExpr); ok {
							// force reflection
							fie.typeName = "reflection"
						}
					}
					fie.zapFunc = loglib.GetLibFunc(fie.allTypeName())
					fields = append(fields, fie)
				}
				structFields[structName] = fields
			}
		}

		return true
	})

	return structs, structFields
}
