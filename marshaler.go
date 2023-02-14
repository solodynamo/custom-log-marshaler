package main

import (
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"os"
)

type PIIMarshaler interface {
	Generate([]string, map[string][]field) string
	GetLibFunc(string) string
}

func generate(path string, loglib PIIMarshaler) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, parser.Mode(0))
	if err != nil {
		panic(err)
	}

	structs, structFields := getFields(f, loglib)

	if len(structs) == 0 {
		panic(fmt.Errorf("cannot generate no suitable struct exists target=%s", path))
	}

	data := loglib.Generate(structs, structFields)
	if data == "" {
		panic(fmt.Errorf("cannot generate file no suitable struct exists target=%s", path))
	}

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Write([]byte(data))
	if err != nil {
		panic(err)
	}
}

func main() {
	path := flag.String("f", "", "target file path")
	lib := flag.String("lib", "", "zap/zerolog?")
	flag.Parse()

	var loglib PIIMarshaler
	switch *lib {
	case "zerolog":
		loglib = &ZeroLog{}
	default:
		loglib = &UberZap{}
	}
	generate(*path, loglib)
}
