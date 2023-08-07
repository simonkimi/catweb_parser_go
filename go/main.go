package main

import "C"
import "catweb_parser/parsers"

//export ParseHtml
func ParseHtml(context *C.char, parserType *C.char, parser *C.char) *C.char {
	contextStr := C.GoString(context)
	parserTypeStr := C.GoString(parserType)
	parserStr := C.GoString(parser)

	result, err := parsers.From(parserTypeStr, parserStr, contextStr)
	if err != nil {
		return C.CString(err.Error())
	} else {
		return C.CString(string(result))
	}
}

func main() {

}
