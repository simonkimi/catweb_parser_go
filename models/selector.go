package models

const (
	SelectorTypeCss      = "css"
	SelectorTypeXpath    = "xpath"
	SelectorTypeJsonPath = "jsonpath"

	SelectorFunctionAuto = "auto"
	SelectorFunctionText = "text"
	SelectorFunctionAttr = "attr"
	SelectorFunctionRaw  = "raw"

	ScriptOutput   = "output"
	ScriptJs       = "js"
	ScriptLua      = "lua"
	ScriptComputed = "computed"
	ScriptReplace  = "replace"
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

type ImageSelector struct {
	ImgUrl    string `json:"imgUrl"`
	ImgWidth  string `json:"imgWidth"`
	ImgHeight string `json:"imgHeight"`
	ImgX      string `json:"imgX"`
	ImgY      string `json:"imgY"`
}

type CommentSelector struct {
	Username *Selector `json:"username"`
	Time     *Selector `json:"time"`
	Score    *Selector `json:"score"`
	Content  *Selector `json:"content"`
	Avatar   *Selector `json:"avatar"`
}

type ExtraSelector struct {
	Id       string    `json:"id"`
	Selector *Selector `json:"selector"`
	Global   bool      `json:"global"`
}
