package main

import (
	"catweb_parser/models"
	"catweb_parser/parser"
	"github.com/dop251/goja"
	"os"
	"testing"
)

func TestFfi(t *testing.T) {
	buffer, err := os.ReadFile("/Users/simonxu/Project/flutter/catweb/test/ffi_request.json")
	if err != nil {
		t.Fatal(err)
	}
	html, err := os.ReadFile("./test_data/list.html")
	if err != nil {
		t.Fatal(err)
	}
	result, err := parser.Parse(models.ListParserType, string(html), buffer)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(result))
}

func TestJs(t *testing.T) {
	vm := goja.New()
	_ = vm.Set("$arg", "1+2")
	value, err := vm.RunString(`eval($arg)`)
	if err != nil {
		panic(err)
	}
	println(value.String())
}
