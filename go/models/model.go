package models

import "encoding/json"

type Params struct {
	Data       string `json:"data"`
	ParserType string `json:"parserType"`
	Parser     string `json:"parser"`
}

type Result struct {
	HasError bool   `json:"hasError"`
	Error    string `json:"error"`
	Data     string `json:"data"`
}

func (i *Result) String() string {
	ret, _ := json.Marshal(i)
	return string(ret)
}

func NewResult(data string) string {
	ret := &Result{
		HasError: false,
		Data:     data,
	}
	return ret.String()
}

func NewErrorResult(err error) string {
	ret := &Result{
		HasError: true,
		Error:    err.Error(),
	}
	return ret.String()
}
