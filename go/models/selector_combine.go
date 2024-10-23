package models

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
