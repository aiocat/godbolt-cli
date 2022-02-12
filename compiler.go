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

const GODBOLT_API = "https://godbolt.org/api/compiler/"

var COMPILER_SOURCES = map[string]string{
	"c": "cg112",
}

type Compiler struct {
	Compiler, Source, FileName, Language string
}

func (c *Compiler) GetCompiler() {
	extension := filepath.Ext(c.FileName)
	c.Language = strings.ToLower(extension)[1:]

	if value, res := COMPILER_SOURCES[c.Language]; res {
		c.Compiler = value
	} else {
		panic("Language is not supported")
	}
}

func (c *Compiler) GetSource() {
	bytes, err := os.ReadFile(c.FileName)

	if err != nil {
		panic(err)
	}

	c.Source = string(bytes)
}

func (c *Compiler) Run() {
	jsonFormat := []byte("{\"source\": " + fmt.Sprintf("%q", c.Source) + ",\"compiler\": \"" + c.Compiler + "\",\"options\": {\"userArguments\": \"\",\"executeParameters\": {\"args\": \"\",\"stdin\": \"\"},\"compilerOptions\": {\"executorRequest\": true,\"skipAsm\": true},\"filters\": {\"execute\": true},\"tools\": [],\"libraries\": []},\"lang\": \"" + c.Language + "\",\"allowStoreCodeDebug\": true}")

	request, err := http.NewRequest("POST", GODBOLT_API+c.Compiler+"/compile", bytes.NewBuffer(jsonFormat))

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
	fmt.Println(string(body))
}
