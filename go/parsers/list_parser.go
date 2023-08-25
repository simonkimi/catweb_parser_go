package parsers

import (
	"catweb_parser/models"
	"catweb_parser/results"
	"catweb_parser/selector"
	"catweb_parser/utils"
)

type ListViewParser struct {
	ParserType      string                  `json:"parserType"`
	Extra           []*models.ExtraSelector `json:"extra"`
	ItemSelector    *models.Selector        `json:"itemSelector"`
	SuccessSelector *models.Selector        `json:"successSelector"`
	FailedSelector  *models.Selector        `json:"failedSelector"`
	Title           *models.Selector        `json:"title"`
	Subtitle        *models.Selector        `json:"subtitle"`
	UploadTime      *models.Selector        `json:"uploadTime"`
	Star            *models.Selector        `json:"star"`
	ImageCount      *models.Selector        `json:"imageCount"`
	Language        *models.Selector        `json:"language"`
	PreviewImage    *models.ImageSelector   `json:"previewImage"`
	BadgeSelector   *models.Selector        `json:"badgeSelector"`
	BadgeItem       *models.TagSelector     `json:"badgeItem"`
	Tag             *models.Selector        `json:"tag"`
	TagItem         *models.TagSelector     `json:"tagItem"`
	Paper           *models.Selector        `json:"paper"`
	IdCode          *models.Selector        `json:"idCode"`
	NextPage        *models.Selector        `json:"nextPage"`
}

func (p *ListViewParser) Parse(content string) (*results.ListParserResult, error) {
	c, root, err := selector.CreateContext(content)
	if err != nil {
		return nil, err
	}
	return &results.ListParserResult{
		NextPage:    c.String(root, p.NextPage),
		IsSuccess:   c.SuccessFlag(root, p.SuccessSelector),
		FailMessage: c.String(root, p.FailedSelector),
		Envs:        c.Env(root, p.Extra),
		Items: utils.Map(c.Nodes(root, p.ItemSelector), func(node *selector.Node) *results.ListParserResultItem {
			return &results.ListParserResultItem{
				Title:        c.String(node, p.Title),
				Subtitle:     c.String(node, p.Subtitle),
				UploadTime:   c.String(node, p.UploadTime),
				Star:         c.Double(node, p.Star),
				ImageCount:   c.Int(node, p.ImageCount),
				PreviewImage: c.Image(node, p.PreviewImage),
				Language:     c.String(node, p.Language),
				IdCode:       c.String(node, p.IdCode),
				Paper:        c.String(node, p.Paper),
				Badges: utils.Map(c.Nodes(node, p.BadgeSelector), func(node *selector.Node) *results.TagResult {
					return c.Tag(node, p.BadgeItem)
				}),
				Tags: utils.Map(c.Nodes(node, p.Tag), func(node *selector.Node) *results.TagResult {
					return c.Tag(node, p.TagItem)
				}),
			}
		}),
		Errors: *c.ErrorList,
	}, nil
}
