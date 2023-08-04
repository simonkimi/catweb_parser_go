package main

import (
	"catweb_parser/models"
	"fmt"
	"os"
	"testing"
)

func TestFfi(t *testing.T) {
	buffer, err := os.ReadFile("test\\ffi_request.json")
	if err != nil {
		t.Fatal(err)
	}
	ffiRequest, err := ParseFfi(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", ffiRequest.Parser.(*models.ListViewParser).ImageCount.Selector)
}
