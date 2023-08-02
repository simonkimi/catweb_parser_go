package parser

import (
	"catweb_parser/models"
	"catweb_parser/results"
	"catweb_parser/selector"
	"catweb_parser/utils"
)

func DetailParser(content string, parser *models.DetailParser) (*results.DetailParserResult, error) {
	context, root, err := selector.CreateContext(content)
	if err != nil {
		return nil, err
	}
	result := &results.DetailParserResult{
		Title:        context.String(root, parser.Title),
		Subtitle:     context.String(root, parser.Subtitle),
		Language:     context.String(root, parser.Language),
		ImageCount:   context.Int(root, parser.ImageCount),
		UploadTime:   context.String(root, parser.UploadTime),
		CountPrePage: context.Int(root, parser.CountPrePage),
		Description:  context.String(root, parser.Description),
		Star:         context.Double(root, parser.Star),
		CoverImage:   context.Image(root, parser.CoverImage),
		NextPage:     context.String(root, parser.NextPage),
		IsSuccess:    context.SuccessFlag(root, parser.SuccessSelector),
		FailMessage:  context.String(root, parser.FailedSelector),
		Env:          context.Env(root, parser.Extra),
		Previews: utils.Map(context.Nodes(root, parser.PreviewSelector), func(node *selector.Node) *results.PreviewItem {
			return &results.PreviewItem{
				PreviewImage: context.Image(node, parser.PreviewImage),
				Target:       context.String(node, parser.Target),
			}
		}),
		Badges: utils.Map(context.Nodes(root, parser.BadgeSelector), func(node *selector.Node) *results.TagResult {
			return context.Tag(node, parser.BadgeItem)
		}),
		Tags: utils.Map(context.Nodes(root, parser.TagSelector), func(node *selector.Node) *results.TagResult {
			return context.Tag(node, parser.TagItem)
		}),
		Comments: utils.Map(context.Nodes(root, parser.CommentSelector), func(node *selector.Node) *results.Comment {
			return context.Comment(node, parser.CommentItem)
		}),
	}
	return result, nil
}
