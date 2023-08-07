package parser

import (
	"catweb_parser/models"
	"encoding/json"
	"errors"
)

func parserBuilder(parserType string, data []byte) (any, error) {
	switch parserType {
	case models.ListParserType:
		var parser models.ListViewParser
		err := json.Unmarshal(data, &parser)
		if err != nil {
			return nil, err
		}
		return &parser, nil
	case models.DetailParserType:
		var parser models.DetailParser
		err := json.Unmarshal(data, &parser)
		if err != nil {
			return nil, err
		}
		return &parser, nil
	case models.ImageReaderParserType:
		var parser models.ImageReaderParser
		err := json.Unmarshal(data, &parser)
		if err != nil {
			return nil, err
		}
		return &parser, nil
	case models.AutoCompleteParserType:
		var parser models.AutoCompleteParser
		err := json.Unmarshal(data, &parser)
		if err != nil {
			return nil, err
		}
		return &parser, nil
	default:
		return nil, errors.New("Unknown parser type" + parserType)
	}
}

func parserExec(parserType string, content string, parser any) (any, error) {
	switch parserType {
	case models.ListParserType:
		return ListParser(content, parser.(*models.ListViewParser))
	case models.DetailParserType:
		return DetailParser(content, parser.(*models.DetailParser))
	case models.ImageReaderParserType:
		return ImageReaderParser(content, parser.(*models.ImageReaderParser))
	case models.AutoCompleteParserType:
		return AutoCompleteParser(content, parser.(*models.AutoCompleteParser))
	default:
		return nil, errors.New("Unknown parser type" + parserType)
	}
}

func Parse(parserType string, content string, data []byte) ([]byte, error) {
	parser, err := parserBuilder(parserType, data)
	if err != nil {
		return nil, err
	}
	result, err := parserExec(parserType, content, parser)
	if err != nil {
		return nil, err
	}
	return json.Marshal(result)
}
