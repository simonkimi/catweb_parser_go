package models

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
	BadgeText       string           `json:"badgeText"`
	BadgeCategory   string           `json:"badgeCategory"`
}

type GalleryParser struct {
	ParserType        string             `json:"parserType"`
	Extra             []*ExtraSelector   `json:"extra"`
	Title             *Selector          `json:"title"`
	Subtitle          *Selector          `json:"subtitle"`
	UploadTime        *Selector          `json:"uploadTime"`
	Star              *Selector          `json:"star"`
	ImgCount          *Selector          `json:"imgCount"`
	PageCount         *Selector          `json:"pageCount"`
	Language          *Selector          `json:"language"`
	CoverImg          *ImageSelector     `json:"coverImg"`
	Description       *Selector          `json:"description"`
	SuccessSelector   *Selector          `json:"successSelector"`
	FailedSelector    *Selector          `json:"failedSelector"`
	ThumbnailSelector *Selector          `json:"thumbnailSelector"`
	Thumbnail         *ImageSelector     `json:"thumbnail"`
	Target            *Selector          `json:"target"`
	CommentSelector   *Selector          `json:"commentSelector"`
	Comments          []*CommentSelector `json:"comments"`
	Tag               *Selector          `json:"tag"`
	TagColor          *Selector          `json:"tagColor"`
	BadgeSelector     *Selector          `json:"badgeSelector"`
	BadgeText         *Selector          `json:"badgeText"`
	BadgeCategory     *Selector          `json:"badgeCategory"`
	ChapterSelector   *Selector          `json:"chapterSelector"`
	ChapterTitle      *Selector          `json:"chapterTitle"`
	ChapterSubtitle   *Selector          `json:"chapterSubtitle"`
	ChapterCover      *ImageSelector     `json:"chapterCover"`
	NextPage          *Selector          `json:"nextPage"`
	CountPrePage      *Selector          `json:"countPrePage"`
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
	ImgCount        *Selector        `json:"imgCount"`
	Language        *Selector        `json:"language"`
	PreviewImg      *ImageSelector   `json:"previewImg"`
	Target          *Selector        `json:"target"`
	Tag             *Selector        `json:"tag"`
	TagColor        *Selector        `json:"tagColor"`
	BadgeSelector   *Selector        `json:"badgeSelector"`
	BadgeText       *Selector        `json:"badgeText"`
	BadgeColor      *Selector        `json:"badgeColor"`
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
