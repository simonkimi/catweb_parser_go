package main

const (
	ListParser = iota
	DetailParser
	ImageReaderParser
	AutoCompleteParser
)

type FfiRequest struct {
	Context    string `json:"context"`
	ParserType int    `json:"parserType"`
	Parser     any    `json:"parser"`
}
