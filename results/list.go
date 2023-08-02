package results

type ListParserResultItem struct {
	Title      *string      `json:"title"`
	Subtitle   *string      `json:"subtitle"`
	UploadTime *string      `json:"uploadTime"`
	Star       *float64     `json:"star"`
	ImgCount   *int64       `json:"imgCount"`
	PreviewImg *ImageResult `json:"previewImg"`
}

type ListParserResult struct {
	Items       []*ListParserResultItem `json:"items"`
	NextPage    *string                 `json:"nextPage"`
	IsSuccess   *bool                   `json:"isSuccess"`
	FailMessage *string                 `json:"failMessage"`
	Env         []*EnvResult            `json:"env"`
}
