package main

import "C"
import (
	"catweb_parser/models"
	"catweb_parser/parsers"
	"encoding/json"
)

//export ParseHtml
func ParseHtml(input *C.char) *C.char {
	inputStr := C.GoString(input)
	params := &models.Params{}
	err := json.Unmarshal([]byte(inputStr), &params)

	if err != nil {
		return C.CString(models.NewErrorResult(err))
	}

	result, err := parsers.From(params.ParserType, params.Parser, params.Data)
	if err != nil {
		return C.CString(models.NewErrorResult(err))
	} else {
		return C.CString(models.NewResult(string(result)))
	}
}

func main() {

}
