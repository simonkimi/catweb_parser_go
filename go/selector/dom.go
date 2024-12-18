package selector

import (
	"catweb_parser/models"
	"catweb_parser/results"
	"catweb_parser/utils"
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/dop251/goja"
	"github.com/ohler55/ojg/oj"
	"golang.org/x/net/html"
	"regexp"
	"strconv"
	"strings"
)

const (
	DomHtml = iota
	DomJson
	DomXml
)

type Node struct {
	nodeType int
	jsonNode any
	HtmlNode *html.Node
}

type ParserContext struct {
	ErrorList *[]*models.ParseError
}

func (c *ParserContext) AddError(err *models.ParseError) {
	*c.ErrorList = append(*c.ErrorList, err)
}

func CreateContext(content string) (*ParserContext, *Node, error) {
	node, err := createNode(content)
	if err != nil {
		return nil, nil, err
	}
	errList := make([]*models.ParseError, 0)
	return &ParserContext{
		ErrorList: &errList,
	}, node, nil
}

func createNode(content string) (*Node, error) {
	if strings.HasPrefix(content, "{") || strings.HasPrefix(content, "[") {
		node, err := oj.ParseString(content)
		if err != nil {
			return nil, err
		}
		return &Node{
			nodeType: DomJson,
			jsonNode: node,
		}, nil
	}

	root, err := htmlquery.Parse(strings.NewReader(content))
	if err != nil {
		return nil, err
	}
	return &Node{
		nodeType: DomHtml,
		HtmlNode: root,
	}, nil
}

func (n *Node) queryValue(selector *models.Selector, errList *[]*models.ParseError) *string {
	var value string
	var found bool
	var err *models.ParseError
	if n.nodeType == DomJson {
		value, found, err = queryJsonFunction(selector, n.jsonNode)
	}
	if n.nodeType == DomHtml {
		value, found, err = queryHtmlFunction(selector, n.HtmlNode)
	}
	if err != nil {
		*errList = append(*errList, err)
		return nil
	}
	if !found {
		return nil
	}
	return &value
}

func regexReplace(selector *models.Selector, input *string, errList *[]*models.ParseError) *string {
	if selector.Regex == "" || input == nil {
		return input
	}
	value := *input
	reg, err := regexp.Compile(selector.Regex)
	if err != nil {
		*errList = append(*errList, models.NewParseError(models.ParserError, fmt.Sprintf("Selector %s regex error: %s", selector.Selector, err.Error())))
		return nil
	}
	match := reg.FindStringSubmatch(value)
	if match == nil {
		*errList = append(*errList, models.NewParseError(models.ElementNotFoundError, fmt.Sprintf("Selector %s regex not match", selector.Selector)))
		return nil
	}
	if selector.Replace == "" {
		return &match[len(match)-1]
	} else {
		rep := selector.Replace
		for i := len(match) - 1; i >= 1; i-- {
			rep = strings.Replace(rep, "$"+strconv.Itoa(i), match[i], -1)
		}
		return &rep
	}
}

func execScript(selector *models.Selector, input *string, errList *[]*models.ParseError) *string {
	if selector.Script == nil || selector.Script.RuntimeType == models.ScriptOutput || input == nil {
		return input
	}
	value := *input
	switch selector.Script.RuntimeType {
	case models.ScriptJs:
		vm := goja.New()
		err := vm.Set("$arg", value)
		if err != nil {
			*errList = append(*errList, models.NewParseError(models.ParserError, fmt.Sprintf("Selector %s script error: %s", selector.Selector, err.Error())))
			return nil
		}
		result, err := vm.RunString(selector.Script.Js)
		if err != nil {
			*errList = append(*errList, models.NewParseError(models.ParserError, fmt.Sprintf("Selector %s script error: %s", selector.Selector, err.Error())))
			return nil
		}
		r := result.String()
		return &r
	case models.ScriptReplace:
		if selector.Script.Replace == nil {
			*errList = append(*errList, models.NewParseError(models.ParserError, fmt.Sprintf("Selector %s script error: replace is nil", selector.Selector)))
			return nil
		}
		newValue, exist := selector.Script.Replace[value]
		if !exist {
			*errList = append(*errList, models.NewParseError(models.ElementNotFoundError, fmt.Sprintf("Selector %s script replace not found", selector.Selector)))
			return nil
		}
		return &newValue
	case models.ScriptOutput:
		return input
	}
	*errList = append(*errList, models.NewParseError(models.InternalError, fmt.Sprintf("Unknown script type: %s", selector.Script.RuntimeType)))
	return nil
}

func (c *ParserContext) SuccessFlag(node *Node, selector *models.Selector) *bool {
	value := c.String(node, selector)
	if value == nil {
		return nil
	}
	v := true
	return &v
}

func (c *ParserContext) Env(node *Node, selector []*models.ExtraSelector) []*results.EnvResult {
	var envs []*results.EnvResult
	for _, s := range selector {
		env := &results.EnvResult{
			Id:     s.Id,
			Value:  c.String(node, s.Selector),
			Global: s.Global,
		}
		envs = append(envs, env)
	}
	return envs
}

func (c *ParserContext) String(node *Node, selector *models.Selector) *string {
	value := node.queryValue(selector, c.ErrorList)
	if value == nil {
		return nil
	}
	// 正则替换
	value = regexReplace(selector, value, c.ErrorList)
	value = execScript(selector, value, c.ErrorList)
	if value != nil {
		return value
	}
	if selector.DefaultValue != "" {
		return &selector.DefaultValue
	}
	return nil
}

func (c *ParserContext) Int(node *Node, selector *models.Selector) *int64 {
	value := c.String(node, selector)
	if value == nil {
		return nil
	}
	v, err := strconv.ParseInt(*value, 10, 64)
	if err != nil {
		c.AddError(models.NewParseError(models.ConverterError, fmt.Sprintf("Selector %s parse %s to int error: %s", selector.Selector, value, err.Error())))
		return nil
	}
	return &v
}

func (c *ParserContext) Double(node *Node, selector *models.Selector) *float64 {
	value := c.String(node, selector)
	if value == nil {
		return nil
	}
	v, err := strconv.ParseFloat(*value, 64)
	if err != nil {
		c.AddError(models.NewParseError(models.ConverterError, fmt.Sprintf("Selector %s parse %s to double error: %s", selector.Selector, value, err.Error())))
		return nil
	}
	return &v
}

func (c *ParserContext) Image(node *Node, selector *models.ImageSelector) *results.ImageResult {
	return &results.ImageResult{
		Url:      c.String(node, selector.Url),
		CacheKey: c.String(node, selector.CacheKey),
		Width:    c.Double(node, selector.Width),
		Height:   c.Double(node, selector.Height),
		X:        c.Double(node, selector.X),
		Y:        c.Double(node, selector.Y),
	}
}

func (c *ParserContext) Tag(node *Node, selector *models.TagSelector) *results.TagResult {
	if selector.Text == nil || selector.Text.IsEmpty() {
		selector.Text = &models.Selector{
			Type: &models.SelectorQuery{
				RuntimeType: models.SelectorTypeSelf,
			},
			Function: &models.SelectorFunction{
				RuntimeType: models.SelectorFunctionText,
			},
		}
	}

	return &results.TagResult{
		Text:     c.String(node, selector.Text),
		Color:    c.String(node, selector.Color),
		Category: c.String(node, selector.Category),
	}
}

func (c *ParserContext) Comment(node *Node, selector *models.CommentSelector) *results.CommentResult {
	return &results.CommentResult{
		Username: c.String(node, selector.Username),
		Content:  c.String(node, selector.Content),
		Time:     c.String(node, selector.Time),
		Score:    c.String(node, selector.Score),
		Avatar:   c.Image(node, selector.Avatar),
	}
}

func (c *ParserContext) Nodes(node *Node, selector *models.Selector) []*Node {
	if node.nodeType == DomJson {
		nodes, err := queryJsonElements(selector, node.jsonNode)
		if err != nil {
			c.AddError(err)
			return nil
		}
		return utils.Map(nodes, func(e any) *Node {
			return &Node{
				jsonNode: e,
				nodeType: DomJson,
			}
		})
	}
	if node.nodeType == DomHtml {
		nodes, err := queryHtmlElements(selector, node.HtmlNode)
		if err != nil {
			c.AddError(err)
			return nil
		}
		return utils.Map(nodes, func(e *html.Node) *Node {
			return &Node{
				HtmlNode: e,
				nodeType: DomHtml,
			}
		})
	}

	c.AddError(models.NewParseError(models.InternalError, fmt.Sprintf("Unknown node type: %d", node.nodeType)))
	return nil
}
