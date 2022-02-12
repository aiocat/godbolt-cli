package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var COMPILER_SOURCES = map[string]string{
	"c":        "cg112",
	"csharp":   "dotnet601csharp",
	"assembly": "nasm21402",
	"c++":      "g112",
	"crystal":  "crystal131",
	"d":        "ldc1_27",
	"dart":     "dart2144",
	"erlang":   "erl2416",
	"fsharp":   "dotnet601fsharp",
	"go":       "gl1170",
	"haskell":  "ghc921",
	"java":     "java1700",
	"kotlin":   "kotlinc1610",
	"llvm":     "irclangtrunk",
	"nim":      "nim160",
	"python":   "python310",
	"ruby":     "ruby302",
	"rust":     "r1580",
	"scala":    "scalac2136",
	"swift":    "swift55",
	"zig":      "z090",
}

type Compiler struct {
	Compiler, Source, FileName, Language string
}

func (c *Compiler) GetCompiler() {
	extension := filepath.Ext(c.FileName)
	c.Language = strings.ToLower(extension)[1:]

	switch c.Language {
	case "h":
		c.Language = "c"
	case "hs":
		c.Language = "haskell"
	case "cpp":
		c.Language = "c++"
	case "cxx":
		c.Language = "c++"
	case "hpp":
		c.Language = "c++"
	case "hxx":
		c.Language = "c++"
	case "ixx":
		c.Language = "c++"
	case "cs":
		c.Language = "csharp"
	case "fs":
		c.Language = "fsharp"
	case "asm":
		c.Language = "assembly"
	case "s":
		c.Language = "assembly"
	case "cr":
		c.Language = "crystal"
	case "erl":
		c.Language = "erlang"
	case "hrl":
		c.Language = "erlang"
	case "ll":
		c.Language = "llvm"
	case "py":
		c.Language = "python"
	case "rb":
		c.Language = "ruby"
	case "rs":
		c.Language = "rust"
	}

	if value, res := COMPILER_SOURCES[c.Language]; res {
		c.Compiler = value
	} else {
		panic("Language is not supported")
	}
}

func (c *Compiler) GetSource() {
	if strings.HasPrefix(c.FileName, "http://") || strings.HasPrefix(c.FileName, "https://") {
		response, err := http.Get(c.FileName)
		if err != nil {
			panic(err)
		}

		defer response.Body.Close()

		bytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}

		c.Source = string(bytes)
	} else {
		bytes, err := os.ReadFile(c.FileName)

		if err != nil {
			panic(err)
		}

		c.Source = string(bytes)
	}
}

func (c *Compiler) Run() []byte {
	jsonFormat := []byte("{\"source\": " + fmt.Sprintf("%q", c.Source) + ",\"compiler\": \"" + c.Compiler + "\",\"options\": {\"userArguments\": \"\",\"executeParameters\": {\"args\": \"\",\"stdin\": \"\"},\"compilerOptions\": {\"executorRequest\": true,\"skipAsm\": true},\"filters\": {\"execute\": true},\"tools\": [],\"libraries\": []},\"lang\": \"" + c.Language + "\",\"allowStoreCodeDebug\": true}")

	request, err := http.NewRequest("POST", GODBOLT_API+"compiler/"+c.Compiler+"/compile", bytes.NewBuffer(jsonFormat))

	if err != nil {
		panic(err)
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	return body
}
