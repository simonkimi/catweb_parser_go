package main

import (
	lua "github.com/yuin/gopher-lua"
	"testing"
)

func TestParser(t *testing.T) {
	l := lua.NewState()
	defer l.Close()
	l.SetGlobal("_ARG", lua.LString("10|L"))
	err := l.DoString("local r, c = string.match(_ARG, \"(%d+)|([LN])\") _RESULT = tonumber(r) * ({L = 5, N = 10})[c]")
	if err != nil {
		t.Fatal(err)
	}
	println(l.GetGlobal("_RESULT").String())
}
