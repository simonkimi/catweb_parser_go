package results

type PreviewItem struct {
	PreviewImage *ImageResult `json:"previewImage"`
	Target       *string      `json:"target"`
}

type DetailParserResult struct {
	Title        *string        `json:"title"`
	Subtitle     *string        `json:"subtitle"`
	Language     *string        `json:"language"`
	ImageCount   *int64         `json:"imageCount"`
	UploadTime   *string        `json:"uploadTime"`
	CountPrePage *int64         `json:"countPrePage"`
	Description  *string        `json:"description"`
	Star         *float64       `json:"star"`
	Previews     []*PreviewItem `json:"previews"`
	CoverImage   *ImageResult   `json:"coverImage"`
	Badges       []*TagResult   `json:"badges"`
	Tags         []*TagResult   `json:"tags"`
	Comments     []*Comment     `json:"comments"`
	NextPage     *string        `json:"nextPage"`
	IsSuccess    *bool          `json:"isSuccess"`
	FailMessage  *string        `json:"failMessage"`
	Env          []*EnvEntity   `json:"env"`
}
