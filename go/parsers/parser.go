package parsers

import "encoding/json"

const (
	ListParserType         = "ListParser"
	DetailParserType       = "DetailParser"
	ImageReaderParserType  = "ImageReaderParser"
	AutoCompleteParserType = "AutoCompleteParser"
)

func From(parserType string, parserData string, content string) (buffer []byte, err error) {
	var result any
	switch parserType {
	case ListParserType:
		parser := ListViewParser{}
		err := json.Unmarshal([]byte(parserData), &parser)
		if err != nil {
			return nil, err
		}
		result, err = parser.Parse(content)
	case DetailParserType:
		parser := DetailParser{}
		err := json.Unmarshal([]byte(parserData), &parser)
		if err != nil {
			return nil, err
		}
		result, err = parser.Parse(content)
	case ImageReaderParserType:
		parser := ImageReaderParser{}
		err := json.Unmarshal([]byte(parserData), &parser)
		if err != nil {
			return nil, err
		}
		result, err = parser.Parse(content)
	case AutoCompleteParserType:
		parser := AutoCompleteParser{}
		err := json.Unmarshal([]byte(parserData), &parser)
		if err != nil {
			return nil, err
		}
		result, err = parser.Parse(content)
	}
	return json.Marshal(result)
}
