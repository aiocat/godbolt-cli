package main

import "os"

const GODBOLT_API = "https://godbolt.org/api/"

func main() {
	if len(os.Args) < 2 {
		panic("File must be given")
	}

	fileName := os.Args[1]

	compiler := new(Compiler)
	compiler.FileName = fileName

	compiler.GetSource()
	compiler.GetCompiler()

	compiler.Run()
}
