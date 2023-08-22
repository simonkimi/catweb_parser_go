package parsers

import (
	"catweb_parser/models"
	"catweb_parser/results"
	"catweb_parser/selector"
	"catweb_parser/utils"
)

type ImageReaderParser struct {
	ParserType      string                  `json:"parserType"`
	Extra           []*models.ExtraSelector `json:"extra"`
	Image           *models.ImageSelector   `json:"image"`
	LargerImage     *models.ImageSelector   `json:"largerImage"`
	RawImage        *models.ImageSelector   `json:"rawImage"`
	Rating          *models.Selector        `json:"rating"`
	Score           *models.Selector        `json:"score"`
	Source          *models.Selector        `json:"source"`
	UploadTime      *models.Selector        `json:"uploadTime"`
	SuccessSelector *models.Selector        `json:"successSelector"`
	FailedSelector  *models.Selector        `json:"failedSelector"`
	BadgeSelector   *models.Selector        `json:"badgeSelector"`
	BadgeItem       *models.TagSelector     `json:"badgeItem"`
	TagSelector     *models.Selector        `json:"tag"`
	TagItem         *models.TagSelector     `json:"tagItem"`
	CommentSelector *models.Selector        `json:"commentSelector"`
	CommentItem     *models.CommentSelector `json:"commentItem"`
}

func (p *ImageReaderParser) Parse(content string) (*results.ImageReaderResult, error) {
	c, root, err := selector.CreateContext(content)
	if err != nil {
		return nil, err
	}
	return &results.ImageReaderResult{
		Image:       c.Image(root, p.Image),
		LargerImage: c.Image(root, p.LargerImage),
		RawImage:    c.Image(root, p.RawImage),
		UploadTime:  c.String(root, p.UploadTime),
		Source:      c.String(root, p.Source),
		Rating:      c.String(root, p.Rating),
		Score:       c.Double(root, p.Score),
		Badges: utils.Map(c.Nodes(root, p.BadgeSelector), func(node *selector.Node) *results.TagResult {
			return c.Tag(node, p.BadgeItem)
		}),
		Tags: utils.Map(c.Nodes(root, p.TagSelector), func(node *selector.Node) *results.TagResult {
			return c.Tag(node, p.TagItem)
		}),
		Comments: utils.Map(c.Nodes(root, p.CommentSelector), func(node *selector.Node) *results.CommentResult {
			return c.Comment(node, p.CommentItem)
		}),
		IsSuccess:   c.SuccessFlag(root, p.SuccessSelector),
		FailMessage: c.String(root, p.FailedSelector),
		Envs:        c.Env(root, p.Extra),
		Errors:      *c.ErrorList,
	}, nil
}
