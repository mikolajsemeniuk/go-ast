package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
)

// https://eli.thegreenplace.net/2021/rewriting-go-source-code-with-ast-tooling/
// https://medium.com/swlh/cool-stuff-with-gos-ast-package-pt-1-981460cddcd7
// https://medium.com/swlh/cool-stuff-with-gos-ast-package-pt-2-e4d39ab7e9db
// https://gist.github.com/imantung/60d0c82b8b1641c0aa1c071e1cf77adf
// https://stackoverflow.com/questions/69545160/check-if-ast-expr-implements-interface-in-go
type Validator interface {
	Valid() bool
}

func main() {
	set := token.NewFileSet()
	f := os.Getenv("GOFILE")
	// f = "../main.go"
	fmt.Println("f:", f)
	file, err := parser.ParseFile(set, f, nil, 0)
	if err != nil {
		log.Fatal(err)
	}

	for _, node := range file.Decls {
		switch node.(type) {

		case *ast.GenDecl:
			genDecl := node.(*ast.GenDecl)
			for _, spec := range genDecl.Specs {
				switch spec.(type) {
				case *ast.TypeSpec:
					typeSpec := spec.(*ast.TypeSpec)

					fmt.Printf("Struct: name=%s\n", typeSpec.Name.Name)

					switch typeSpec.Type.(type) {
					case *ast.StructType:
						structType := typeSpec.Type.(*ast.StructType)
						for _, field := range structType.Fields.List {
							i := field.Type.(*ast.Ident)
							j := field.Type.(ast.Expr)
							_, ok := j.(Validator)
							if ok {
								fmt.Println("implements Validator interface")
							}

							fieldType := i.Name
							for _, name := range field.Names {
								fmt.Printf("\tField: name=%s type=%s tag=%s\n", name.Name, fieldType, field.Tag.Value)
							}
						}
					}
				}
			}
		}
	}
}
