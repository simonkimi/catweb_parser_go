package results

type CommentResult struct {
	Username *string      `json:"username"`
	Content  *string      `json:"content"`
	Time     *string      `json:"time"`
	Score    *string      `json:"score"`
	Avatar   *ImageResult `json:"avatar"`
}

type TagResult struct {
	Text     *string `json:"text"`
	Color    *string `json:"color"`
	Category *string `json:"category"`
}

type ImageResult struct {
	Url      *string  `json:"url"`
	CacheKey *string  `json:"cacheKey"`
	Width    *float64 `json:"width"`
	Height   *float64 `json:"height"`
	X        *float64 `json:"x"`
	Y        *float64 `json:"y"`
}

type EnvResult struct {
	Id     string  `json:"id"`
	Global bool    `json:"global"`
	Value  *string `json:"value"`
}
