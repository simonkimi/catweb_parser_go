package results

import (
	"catweb_parser/json_type/boolean"
)

type ListParserResultItem struct {
	Title      string       `json:"title"`
	Subtitle   string       `json:"subtitle"`
	UploadTime string       `json:"uploadTime"`
	Star       *float64     `json:"star"`
	ImgCount   *int64       `json:"imgCount"`
	PreviewImg *ImageResult `json:"previewImg"`
}

type ListParserResult struct {
	Items       []ListParserResultItem `json:"items"`
	NextPage    string                 `json:"nextPage"`
	IsSuccess   *boolean.Boolean       `json:"isSuccess"`
	FailMessage string                 `json:"failMessage"`
	Env         []*EnvEntity           `json:"env"`
}
