package results

import "catweb_parser/models"

type ImageReaderResult struct {
	Image       *ImageResult         `json:"image"`
	LargerImage *ImageResult         `json:"largerImage"`
	RawImage    *ImageResult         `json:"rawImage"`
	UploadTime  *string              `json:"uploadTime"`
	Source      *string              `json:"source"`
	Rating      *string              `json:"rating"`
	Score       *float64             `json:"score"`
	Badges      []*TagResult         `json:"badges"`
	Tags        []*TagResult         `json:"tags"`
	Comments    []*CommentResult     `json:"comments"`
	IsSuccess   *bool                `json:"isSuccess"`
	FailMessage *string              `json:"failMessage"`
	Envs        []*EnvResult         `json:"envs"`
	Errors      []*models.ParseError `json:"errors"`
}
