package main

func main() {
	compiler := new(Compiler)
	compiler.FileName = "./test/test.c"

	compiler.GetSource()
	compiler.GetCompiler()

	compiler.Run()
}
