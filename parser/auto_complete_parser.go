package parser

import (
	"catweb_parser/models"
	"catweb_parser/results"
	"catweb_parser/selector"
	"catweb_parser/utils"
)

func AutoCompleteParser(content string, parser *models.AutoCompleteParser) (*results.AutoCompleteResult, error) {
	c, root, err := selector.CreateContext(content)
	if err != nil {
		return nil, err
	}
	return &results.AutoCompleteResult{
		Items: utils.Map(c.Nodes(root, parser.ItemSelector), func(node *selector.Node) *results.AutoCompleteResultItem {
			return &results.AutoCompleteResultItem{
				Title:    c.String(node, parser.ItemTitle),
				Subtitle: c.String(node, parser.ItemSubtitle),
				Complete: c.String(node, parser.ItemComplete),
			}
		}),
		IsSuccess:   c.SuccessFlag(root, parser.SuccessSelector),
		FailMessage: c.String(root, parser.FailedSelector),
		Env:         c.Env(root, parser.Extra),
	}, nil
}
