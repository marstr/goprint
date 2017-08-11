package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"log"
	"os"
)

type astPrinter struct {
	io.Writer
	indent string
}

func main() {
	errLog := log.New(os.Stderr, "", 0)
	for _, item := range os.Args[1:] {
		info, err := os.Stat(item)
		if err != nil {
			errLog.Printf("Unable to find \"%s\": %v", item, err)
			continue
		}

		var result ast.Node

		var files token.FileSet

		if info.IsDir() {
			var packages map[string]*ast.Package
			packages, err = parser.ParseDir(&files, item, nil, parser.ParseComments)

			for _, v := range packages {
				result = v
			}
		} else {
			var file *ast.File
			file, err = parser.ParseFile(&files, item, nil, parser.ParseComments)
			result = file
		}
		if err != nil {
			errLog.Printf("Unable to parse \"%s\": %v", item, err)
			continue
		}

		ast.Print(&files, result)
	}
}
