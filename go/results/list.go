package results

import "catweb_parser/models"

type ListParserResultItem struct {
	Title        *string      `json:"title"`
	Subtitle     *string      `json:"subtitle"`
	UploadTime   *string      `json:"uploadTime"`
	Star         *float64     `json:"star"`
	ImageCount   *int64       `json:"imageCount"`
	PreviewImage *ImageResult `json:"previewImage"`
	Badges       []*TagResult `json:"badges"`
	Tags         []*TagResult `json:"tags"`
	Language     *string      `json:"language"`
	Target       *string      `json:"target"`
	Paper        *string      `json:"paper"`
}

type ListParserResult struct {
	Items       []*ListParserResultItem `json:"items"`
	NextPage    *string                 `json:"nextPage"`
	IsSuccess   *bool                   `json:"isSuccess"`
	FailMessage *string                 `json:"failMessage"`
	Envs        []*EnvResult            `json:"envs"`
	Errors      []*models.ParseError    `json:"errors"`
}
