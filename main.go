package main

import "C"

//export ParseHtml
func ParseHtml(context *C.char, parserType *C.char, parser *C.char) *C.char {
	contextStr := C.GoString(context)
	parserTypeStr := C.GoString(parserType)
	parserStr := C.GoString(parser)
}
