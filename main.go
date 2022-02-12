package main

import (
	"fmt"
	"os"
)

const GODBOLT_API = "https://godbolt.org/api/"

func main() {
	if len(os.Args) < 2 {
		panic("File must be given")
	}

	fileName := os.Args[1]
	var outputFile string

	if len(os.Args) > 2 {
		outputFile = os.Args[2]
	} else {
		outputFile = ""
	}

	compiler := new(Compiler)
	compiler.FileName = fileName

	compiler.GetSource()
	compiler.GetCompiler()

	body := compiler.Run()
	if outputFile == "" {
		fmt.Printf("%s\n", body)
	} else {
		os.WriteFile(outputFile, body, os.ModeAppend)
	}
}
