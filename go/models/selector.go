package models

import "strings"

const (
	SelectorTypeDefault  = ""
	SelectorTypeSelf     = "self"
	SelectorTypeCss      = "css"
	SelectorTypeXpath    = "xpath"
	SelectorTypeJsonPath = "jsonpath"

	SelectorFunctionDefault = ""
	SelectorFunctionText    = "text"
	SelectorFunctionAttr    = "attr"
	SelectorFunctionRaw     = "raw"

	ScriptOutput = "output"
	ScriptJs     = "js"
)

type ScriptField struct {
	Script string `json:"script"`
	Type   string `json:"type"`
}

type Selector struct {
	Selector     string       `json:"selector"`
	Type         string       `json:"type"`
	Function     string       `json:"function"`
	Param        string       `json:"param"`
	Regex        string       `json:"regex"`
	Replace      string       `json:"replace"`
	Script       *ScriptField `json:"script"`
	DefaultValue string       `json:"defaultValue"`
}

func (s *Selector) IsEmpty() bool {
	return strings.TrimSpace(s.Selector) == "" && s.Type != SelectorTypeSelf &&
		(s.Function == SelectorFunctionDefault) &&
		s.Param == "" && s.Regex == "" && (s.Script == nil || s.Script.Script == "")
}

type ImageSelector struct {
	Url      *Selector `json:"url"`
	CacheKey *Selector `json:"cacheKey"`
	Width    *Selector `json:"width"`
	Height   *Selector `json:"height"`
	X        *Selector `json:"x"`
	Y        *Selector `json:"y"`
}

type CommentSelector struct {
	Username *Selector      `json:"username"`
	Time     *Selector      `json:"time"`
	Score    *Selector      `json:"score"`
	Content  *Selector      `json:"content"`
	Avatar   *ImageSelector `json:"avatar"`
}

type ExtraSelector struct {
	Id       string    `json:"id"`
	Selector *Selector `json:"selector"`
	Global   bool      `json:"global"`
}

type TagSelector struct {
	Text     *Selector `json:"text"`
	Color    *Selector `json:"color"`
	Category *Selector `json:"category"`
}
