package parsers

import (
	"catweb_parser/models"
	"catweb_parser/results"
	"catweb_parser/selector"
	"catweb_parser/utils"
)

type DetailParser struct {
	ParserType      string                  `json:"parserType"`
	Extra           []*models.ExtraSelector `json:"extra"`
	Title           *models.Selector        `json:"title"`
	Subtitle        *models.Selector        `json:"subtitle"`
	UploadTime      *models.Selector        `json:"uploadTime"`
	Star            *models.Selector        `json:"star"`
	ImageCount      *models.Selector        `json:"imageCount"`
	PageCount       *models.Selector        `json:"pageCount"`
	Language        *models.Selector        `json:"language"`
	CoverImage      *models.ImageSelector   `json:"coverImage"`
	Description     *models.Selector        `json:"description"`
	SuccessSelector *models.Selector        `json:"successSelector"`
	FailedSelector  *models.Selector        `json:"failedSelector"`
	PreviewSelector *models.Selector        `json:"thumbnailSelector"`
	PreviewImage    *models.ImageSelector   `json:"thumbnail"`
	Target          *models.Selector        `json:"target"`
	CommentSelector *models.Selector        `json:"commentSelector"`
	CommentItem     *models.CommentSelector `json:"commentItem"`
	BadgeSelector   *models.Selector        `json:"badgeSelector"`
	BadgeItem       *models.TagSelector     `json:"badgeItem"`
	TagSelector     *models.Selector        `json:"tagSelector"`
	TagItem         *models.TagSelector     `json:"tagItem"`
	ChapterSelector *models.Selector        `json:"chapterSelector"`
	ChapterTitle    *models.Selector        `json:"chapterTitle"`
	ChapterSubtitle *models.Selector        `json:"chapterSubtitle"`
	ChapterCover    *models.ImageSelector   `json:"chapterCover"`
	NextPage        *models.Selector        `json:"nextPage"`
	CountPrePage    *models.Selector        `json:"countPrePage"`
}

func (p *DetailParser) Parse(content string) (*results.DetailParserResult, error) {
	c, root, err := selector.CreateContext(content)
	if err != nil {
		return nil, err
	}
	result := &results.DetailParserResult{
		Title:        c.String(root, p.Title),
		Subtitle:     c.String(root, p.Subtitle),
		Language:     c.String(root, p.Language),
		ImageCount:   c.Int(root, p.ImageCount),
		UploadTime:   c.String(root, p.UploadTime),
		CountPrePage: c.Int(root, p.CountPrePage),
		Description:  c.String(root, p.Description),
		Star:         c.Double(root, p.Star),
		CoverImage:   c.Image(root, p.CoverImage),
		NextPage:     c.String(root, p.NextPage),
		IsSuccess:    c.SuccessFlag(root, p.SuccessSelector),
		FailMessage:  c.String(root, p.FailedSelector),
		Envs:         c.Env(root, p.Extra),
		Previews: utils.Map(c.Nodes(root, p.PreviewSelector), func(node *selector.Node) *results.PreviewItem {
			return &results.PreviewItem{
				PreviewImage: c.Image(node, p.PreviewImage),
				Target:       c.String(node, p.Target),
			}
		}),
		Badges: utils.Map(c.Nodes(root, p.BadgeSelector), func(node *selector.Node) *results.TagResult {
			return c.Tag(node, p.BadgeItem)
		}),
		Tags: utils.Map(c.Nodes(root, p.TagSelector), func(node *selector.Node) *results.TagResult {
			return c.Tag(node, p.TagItem)
		}),
		Comments: utils.Map(c.Nodes(root, p.CommentSelector), func(node *selector.Node) *results.CommentResult {
			return c.Comment(node, p.CommentItem)
		}),
		Errors: *c.ErrorList,
	}
	return result, nil
}
