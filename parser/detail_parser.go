package parser

import (
	"catweb_parser/models"
	"catweb_parser/results"
	"catweb_parser/selector"
	"catweb_parser/utils"
)

func DetailParser(content string, parser *models.DetailParser) (*results.DetailParserResult, error) {
	c, root, err := selector.CreateContext(content)
	if err != nil {
		return nil, err
	}
	result := &results.DetailParserResult{
		Title:        c.String(root, parser.Title),
		Subtitle:     c.String(root, parser.Subtitle),
		Language:     c.String(root, parser.Language),
		ImageCount:   c.Int(root, parser.ImageCount),
		UploadTime:   c.String(root, parser.UploadTime),
		CountPrePage: c.Int(root, parser.CountPrePage),
		Description:  c.String(root, parser.Description),
		Star:         c.Double(root, parser.Star),
		CoverImage:   c.Image(root, parser.CoverImage),
		NextPage:     c.String(root, parser.NextPage),
		IsSuccess:    c.SuccessFlag(root, parser.SuccessSelector),
		FailMessage:  c.String(root, parser.FailedSelector),
		Envs:         c.Env(root, parser.Extra),
		Previews: utils.Map(c.Nodes(root, parser.PreviewSelector), func(node *selector.Node) *results.PreviewItem {
			return &results.PreviewItem{
				PreviewImage: c.Image(node, parser.PreviewImage),
				Target:       c.String(node, parser.Target),
			}
		}),
		Badges: utils.Map(c.Nodes(root, parser.BadgeSelector), func(node *selector.Node) *results.TagResult {
			return c.Tag(node, parser.BadgeItem)
		}),
		Tags: utils.Map(c.Nodes(root, parser.TagSelector), func(node *selector.Node) *results.TagResult {
			return c.Tag(node, parser.TagItem)
		}),
		Comments: utils.Map(c.Nodes(root, parser.CommentSelector), func(node *selector.Node) *results.CommentResult {
			return c.Comment(node, parser.CommentItem)
		}),
	}
	result.Errors = *c.ErrorList
	return result, nil
}
