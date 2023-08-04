package parser

import (
	"catweb_parser/models"
	"catweb_parser/results"
	"catweb_parser/selector"
	"catweb_parser/utils"
)

func ListParser(content string, parser *models.ListViewParser) (*results.ListParserResult, error) {
	c, root, err := selector.CreateContext(content)
	if err != nil {
		return nil, err
	}
	return &results.ListParserResult{
		NextPage:    c.String(root, parser.NextPage),
		IsSuccess:   c.SuccessFlag(root, parser.SuccessSelector),
		FailMessage: c.String(root, parser.FailedSelector),
		Env:         c.Env(root, parser.Extra),
		Items: utils.Map(c.Nodes(root, parser.ItemSelector), func(node *selector.Node) *results.ListParserResultItem {
			return &results.ListParserResultItem{
				Title:        c.String(node, parser.Title),
				Subtitle:     c.String(node, parser.Subtitle),
				UploadTime:   c.String(node, parser.UploadTime),
				Star:         c.Double(node, parser.Star),
				ImageCount:   c.Int(node, parser.ImageCount),
				PreviewImage: c.Image(node, parser.PreviewImage),
			}
		}),
	}, nil
}
