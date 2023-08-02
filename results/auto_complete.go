package results

type AutoCompleteResultItem struct {
	Title    *string `json:"title"`
	Subtitle *string `json:"subtitle"`
	Complete *string `json:"complete"`
}

type AutoCompleteResult struct {
	Items       []*AutoCompleteResultItem `json:"items"`
	IsSuccess   *bool                     `json:"isSuccess"`
	FailMessage *string                   `json:"failMessage"`
	Env         []*EnvEntity              `json:"env"`
}
