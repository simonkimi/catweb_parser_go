package main

import (
	"catweb_parser/models"
	"encoding/json"
)

type FfiRequest struct {
	Context    string `json:"context"`
	ParserType string `json:"parserType"`
	Parser     any    `json:"parser"`
}

func ParseFfi(data []byte) (*FfiRequest, error) {
	ffiRequest := &FfiRequest{}
	err := json.Unmarshal(data, ffiRequest)
	if err != nil {
		return nil, err
	}
	parserJson, err := json.Marshal(ffiRequest.Parser)
	if err != nil {
		return nil, err
	}
	switch ffiRequest.ParserType {
	case ListParser:
		var parser models.ListViewParser
		err = json.Unmarshal(parserJson, &parser)
		if err != nil {
			return nil, err
		}
		ffiRequest.Parser = &parser
	case DetailParser:
		var parser models.DetailParser
		err = json.Unmarshal(parserJson, &parser)
		if err != nil {
			return nil, err
		}
		ffiRequest.Parser = &parser
	case ImageReaderParser:
		var parser models.ImageReaderParser
		err = json.Unmarshal(parserJson, &parser)
		if err != nil {
			return nil, err
		}
		ffiRequest.Parser = &parser
	case AutoCompleteParser:
		var parser models.AutoCompleteParser
		err = json.Unmarshal(parserJson, &parser)
		if err != nil {
			return nil, err
		}
		ffiRequest.Parser = &parser
	}

	return ffiRequest, nil
}
