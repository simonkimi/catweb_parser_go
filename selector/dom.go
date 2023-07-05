package selector

import (
	"catweb_parser/models"
	"catweb_parser/utils"
	"encoding/json"
	"fmt"
	"github.com/dop251/goja"
	"golang.org/x/net/html"
	"regexp"
	"strconv"
	"strings"
)

const (
	DomHtml = iota
	DomJson
)

type Node struct {
	nodeType int
	jsonNode any
	htmlNode *html.Node
}

func (n *Node) SelectNodes(selector *models.Selector) ([]*Node, *models.ParseError) {
	if n.nodeType == DomJson {
		nodes, err := queryJsonElements(selector, n.jsonNode)
		wraps := utils.Select(nodes, func(e any) *Node {
			return &Node{
				jsonNode: e,
				nodeType: DomJson,
			}
		})
		return wraps, err
	}
	if n.nodeType == DomHtml {
		nodes, err := queryHtmlElements(selector, n.htmlNode)
		wraps := utils.Select(nodes, func(e *html.Node) *Node {
			return &Node{
				htmlNode: e,
				nodeType: DomHtml,
			}
		})
		return wraps, err
	}
	return nil, models.NewParseError(models.InternalError, fmt.Sprintf("Unknown node type: %d", n.nodeType))
}

func (n *Node) queryValue(selector *models.Selector, errList *[]*models.ParseError) *string {
	var value string
	var found bool
	var err *models.ParseError
	if n.nodeType == DomJson {
		value, found, err = queryJsonFunction(selector, n.jsonNode)
	}
	if n.nodeType == DomHtml {
		value, found, err = queryHtmlFunction(selector, n.htmlNode)
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

func (n *Node) String(selector *models.Selector, errList *[]*models.ParseError) *string {
	value := n.queryValue(selector, errList)
	if value == nil {
		return nil
	}
	// 正则替换
	value = regexReplace(selector, value, errList)
	value = execScript(selector, value, errList)
	if value != nil {
		return value
	}
	if selector.DefaultValue != "" {
		return &selector.DefaultValue
	}
	return nil
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
	if selector.Script.Script == "" || selector.Script.Type == models.ScriptOutput || input == nil {
		return input
	}
	value := *input
	switch selector.Script.Type {
	case models.ScriptComputed:
		// 计算属性, 直接调用js进行计算
		vm := goja.New()
		result, err := vm.RunString(value)
		if err != nil {
			*errList = append(*errList, models.NewParseError(models.ParserError, fmt.Sprintf("Selector %s script error: %s", selector.Selector, err.Error())))
			return nil
		}
		r := result.String()
		return &r
	case models.ScriptJs:
		// 执行javascript
		vm := goja.New()
		err := vm.Set("$arg", value)
		if err != nil {
			*errList = append(*errList, models.NewParseError(models.ParserError, fmt.Sprintf("Selector %s script error: %s", selector.Selector, err.Error())))
			return nil
		}
		result, err := vm.RunString(selector.Script.Script)
		if err != nil {
			*errList = append(*errList, models.NewParseError(models.ParserError, fmt.Sprintf("Selector %s script error: %s", selector.Selector, err.Error())))
			return nil
		}
		r := result.String()
		return &r
	case models.ScriptReplace:
		// 替换, 传入json映射, 返回替换后的对象
		script := strings.TrimSpace(selector.Script.Script)
		if !strings.HasPrefix(script, "{") || !strings.HasSuffix(script, "}") {
			*errList = append(*errList, models.NewParseError(models.ParserError, fmt.Sprintf("Selector %s script error: %s", selector.Selector, "script must be a json object")))
			return nil
		}
		objs := make(map[string]string)
		err := json.Unmarshal([]byte(script), &objs)
		if err != nil {
			*errList = append(*errList, models.NewParseError(models.ParserError, fmt.Sprintf("Selector %s script error: %s", selector.Selector, err.Error())))
			return nil
		}
		val, exist := objs[*input]
		if exist {
			return &val
		}
		return nil
	}
	*errList = append(*errList, models.NewParseError(models.InternalError, fmt.Sprintf("Unknown script type: %s", selector.Script.Type)))
	return nil
}
