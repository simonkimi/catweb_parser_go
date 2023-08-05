package parser

import (
	"catweb_parser/models"
	"catweb_parser/results"
	"catweb_parser/selector"
	"catweb_parser/utils"
)

func ImageReaderParser(content string, parser *models.ImageReaderParser) (*results.ImageReaderResult, error) {
	c, root, err := selector.CreateContext(content)
	if err != nil {
		return nil, err
	}
	result := &results.ImageReaderResult{
		Image:       c.Image(root, parser.Image),
		LargerImage: c.Image(root, parser.LargerImage),
		RawImage:    c.Image(root, parser.RawImage),
		UploadTime:  c.String(root, parser.UploadTime),
		Source:      c.String(root, parser.Source),
		Rating:      c.String(root, parser.Rating),
		Score:       c.Double(root, parser.Score),
		Badges: utils.Map(c.Nodes(root, parser.BadgeSelector), func(node *selector.Node) *results.TagResult {
			return c.Tag(node, parser.BadgeItem)
		}),
		Tags: utils.Map(c.Nodes(root, parser.TagSelector), func(node *selector.Node) *results.TagResult {
			return c.Tag(node, parser.TagItem)
		}),
		Comments: utils.Map(c.Nodes(root, parser.CommentSelector), func(node *selector.Node) *results.CommentResult {
			return c.Comment(node, parser.CommentItem)
		}),
		IsSuccess:   c.SuccessFlag(root, parser.SuccessSelector),
		FailMessage: c.String(root, parser.FailedSelector),
		Envs:        c.Env(root, parser.Extra),
	}
	result.Errors = *c.ErrorList
	return result, nil
}
