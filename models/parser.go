package models

import "encoding/json"

const (
	ListParserType         = "ListParser"
	DetailParserType       = "DetailParser"
	ImageReaderParserType  = "ImageReaderParser"
	AutoCompleteParserType = "AutoCompleteParser"
)

type ImageReaderParser struct {
	ParserType      string           `json:"parserType"`
	Extra           []*ExtraSelector `json:"extra"`
	Image           *ImageSelector   `json:"image"`
	LargerImage     *ImageSelector   `json:"largerImage"`
	RawImage        *ImageSelector   `json:"rawImage"`
	Rating          *Selector        `json:"rating"`
	Score           *Selector        `json:"score"`
	Source          *Selector        `json:"source"`
	UploadTime      *Selector        `json:"uploadTime"`
	SuccessSelector *Selector        `json:"successSelector"`
	FailedSelector  *Selector        `json:"failedSelector"`
	BadgeSelector   *Selector        `json:"badgeSelector"`
	BadgeItem       *TagSelector     `json:"badgeItem"`
	TagSelector     *Selector        `json:"tag"`
	TagItem         *TagSelector     `json:"tagItem"`
	CommentSelector *Selector        `json:"commentSelector"`
	CommentItem     *CommentSelector `json:"commentItem"`
}

type DetailParser struct {
	ParserType      string           `json:"parserType"`
	Extra           []*ExtraSelector `json:"extra"`
	Title           *Selector        `json:"title"`
	Subtitle        *Selector        `json:"subtitle"`
	UploadTime      *Selector        `json:"uploadTime"`
	Star            *Selector        `json:"star"`
	ImageCount      *Selector        `json:"imageCount"`
	PageCount       *Selector        `json:"pageCount"`
	Language        *Selector        `json:"language"`
	CoverImage      *ImageSelector   `json:"coverImage"`
	Description     *Selector        `json:"description"`
	SuccessSelector *Selector        `json:"successSelector"`
	FailedSelector  *Selector        `json:"failedSelector"`
	PreviewSelector *Selector        `json:"thumbnailSelector"`
	PreviewImage    *ImageSelector   `json:"thumbnail"`
	Target          *Selector        `json:"target"`
	CommentSelector *Selector        `json:"commentSelector"`
	CommentItem     *CommentSelector `json:"commentItem"`
	BadgeSelector   *Selector        `json:"badgeSelector"`
	BadgeItem       *TagSelector     `json:"badgeItem"`
	TagSelector     *Selector        `json:"tagSelector"`
	TagItem         *TagSelector     `json:"tagItem"`
	ChapterSelector *Selector        `json:"chapterSelector"`
	ChapterTitle    *Selector        `json:"chapterTitle"`
	ChapterSubtitle *Selector        `json:"chapterSubtitle"`
	ChapterCover    *ImageSelector   `json:"chapterCover"`
	NextPage        *Selector        `json:"nextPage"`
	CountPrePage    *Selector        `json:"countPrePage"`
}

type ListViewParser struct {
	ParserType      string           `json:"parserType"`
	Extra           []*ExtraSelector `json:"extra"`
	ItemSelector    *Selector        `json:"itemSelector"`
	SuccessSelector *Selector        `json:"successSelector"`
	FailedSelector  *Selector        `json:"failedSelector"`
	Title           *Selector        `json:"title"`
	Subtitle        *Selector        `json:"subtitle"`
	UploadTime      *Selector        `json:"uploadTime"`
	Star            *Selector        `json:"star"`
	ImageCount      *Selector        `json:"imageCount"`
	Language        *Selector        `json:"language"`
	PreviewImage    *ImageSelector   `json:"previewImage"`
	Target          *Selector        `json:"target"`
	BadgeSelector   *Selector        `json:"badgeSelector"`
	BadgeItem       *TagSelector     `json:"badgeItem"`
	Tag             *Selector        `json:"tag"`
	TagItem         *TagSelector     `json:"tagItem"`
	Paper           *Selector        `json:"paper"`
	IdCode          *Selector        `json:"idCode"`
	NextPage        *Selector        `json:"nextPage"`
}

type AutoCompleteParser struct {
	ParserType      string           `json:"parserType"`
	Extra           []*ExtraSelector `json:"extra"`
	ItemSelector    *Selector        `json:"itemSelector"`
	ItemComplete    *Selector        `json:"itemComplete"`
	ItemTitle       *Selector        `json:"itemTitle"`
	ItemSubtitle    *Selector        `json:"itemSubtitle"`
	SuccessSelector *Selector        `json:"successSelector"`
	FailedSelector  *Selector        `json:"failedSelector"`
}

func ParserBuilder(parserType string, data []byte) (any, error) {
	var parser any
	switch parserType {
	case ListParserType:
		parser = &ListViewParser{}
	case DetailParserType:
		parser = &DetailParser{}
	case ImageReaderParserType:
		parser = &ImageReaderParser{}
	case AutoCompleteParserType:
		parser = &AutoCompleteParser{}
	}
	err := json.Unmarshal(data, parser)
	if err != nil {
		return nil, err
	}
	return parser, nil
}
