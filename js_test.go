package main

import (
	"github.com/dop251/goja"
	"testing"
)

func TestJs(t *testing.T) {
	vm := goja.New()
	_ = vm.Set("$arg", "4|L")
	value, err := vm.RunString(`
	let[r,c]=$arg.split('|');parseInt(r)*{'L':5,'N':10}[c];
	`)
	if err != nil {
		panic(err)
	}
	println(value.String())
}
