package main

import "os"

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
