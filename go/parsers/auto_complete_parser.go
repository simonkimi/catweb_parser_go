package parsers

import (
	"catweb_parser/models"
	"catweb_parser/results"
	"catweb_parser/selector"
	"catweb_parser/utils"
)

type AutoCompleteParser struct {
	ParserType      string                  `json:"parserType"`
	Extra           []*models.ExtraSelector `json:"extra"`
	ItemSelector    *models.Selector        `json:"itemSelector"`
	ItemComplete    *models.Selector        `json:"itemComplete"`
	ItemTitle       *models.Selector        `json:"itemTitle"`
	ItemSubtitle    *models.Selector        `json:"itemSubtitle"`
	SuccessSelector *models.Selector        `json:"successSelector"`
	FailedSelector  *models.Selector        `json:"failedSelector"`
}

func (p *AutoCompleteParser) Parse(content string) (*results.AutoCompleteResult, error) {
	c, root, err := selector.CreateContext(content)
	if err != nil {
		return nil, err
	}
	return &results.AutoCompleteResult{
		Items: utils.Map(c.Nodes(root, p.ItemSelector), func(node *selector.Node) *results.AutoCompleteResultItem {
			return &results.AutoCompleteResultItem{
				Title:    c.String(node, p.ItemTitle),
				Subtitle: c.String(node, p.ItemSubtitle),
				Complete: c.String(node, p.ItemComplete),
			}
		}),
		IsSuccess:   c.SuccessFlag(root, p.SuccessSelector),
		FailMessage: c.String(root, p.FailedSelector),
		Envs:        c.Env(root, p.Extra),
	}, nil
}
