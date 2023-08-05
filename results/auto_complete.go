package results

import "catweb_parser/models"

type AutoCompleteResultItem struct {
	Title    *string `json:"title"`
	Subtitle *string `json:"subtitle"`
	Complete *string `json:"complete"`
}

type AutoCompleteResult struct {
	Items       []*AutoCompleteResultItem `json:"items"`
	IsSuccess   *bool                     `json:"isSuccess"`
	FailMessage *string                   `json:"failMessage"`
	Envs        []*EnvResult              `json:"env"`
	Errors      []*models.ParseError      `json:"errors"`
}
