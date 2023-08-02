package results

type ImageReaderResult struct {
	Image       *ImageResult `json:"image"`
	LargerImage *ImageResult `json:"largerImage"`
	RawImage    *ImageResult `json:"rawImage"`
	UploadTime  string       `json:"uploadTime"`
	Source      string       `json:"source"`
	Rating      string       `json:"rating"`
	Score       *float64     `json:"score"`
	Badges      []*TagResult `json:"badges"`
	Tags        []*TagResult `json:"tags"`
	Comments    []*Comment   `json:"comments"`
	IsSuccess   bool         `json:"isSuccess"`
	FailMessage string       `json:"failMessage"`
	Env         []*EnvEntity `json:"env"`
}
