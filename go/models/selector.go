package models

import (
	"catweb_parser/utils"
)

const (
	SelectorTypeSelf     = "self"
	SelectorTypeCss      = "css"
	SelectorTypeXpath    = "xpath"
	SelectorTypeJsonPath = "jsonPath"

	SelectorFunctionText = "text"
	SelectorFunctionAttr = "attr"
	SelectorFunctionRaw  = "raw"

	ScriptOutput  = "output"
	ScriptReplace = "replace"
	ScriptJs      = "js"
)

type Selector struct {
	Selector     string            `json:"selector"`
	Type         *SelectorQuery    `json:"type"`
	Function     *SelectorFunction `json:"function"`
	Param        string            `json:"param"`
	Regex        string            `json:"regex"`
	Replace      string            `json:"replace"`
	Script       *SelectorScript   `json:"script"`
	DefaultValue string            `json:"defaultValue"`
}

type SelectorQuery struct {
	RuntimeType string `json:"runtimeType"`
	Css         string `json:"css"`
	Xpath       string `json:"xpath"`
	JsonPath    string `json:"jsonPath"`
}

type SelectorScript struct {
	RuntimeType string            `json:"runtimeType"`
	Replace     map[string]string `json:"replace"`
	Js          string            `json:"js"`
}

type SelectorFunction struct {
	RuntimeType string `json:"runtimeType"`
	Text        string `json:"text"`
	Attr        string `json:"attr"`
	Raw         string `json:"raw"`
}

func (s *Selector) IsEmpty() bool {
	return utils.IsEmptyOrWhiteSpace(s.Selector) &&
		s.Type.RuntimeType == SelectorTypeSelf &&
		s.Function.RuntimeType == SelectorFunctionText &&
		utils.IsEmptyOrWhiteSpace(s.Param) &&
		utils.IsEmptyOrWhiteSpace(s.Regex) &&
		s.Script.RuntimeType == ScriptOutput &&
		utils.IsEmptyOrWhiteSpace(s.DefaultValue)
}
