package selector

import (
	"bytes"
	"catweb_parser/models"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"strings"
)

func queryHtmlElements(selector *models.Selector, node *html.Node) ([]*html.Node, *models.ParseError) {
	if strings.TrimSpace(selector.Selector) == "" || selector.Type == models.SelectorTypeSelf {
		return []*html.Node{node}, nil
	}

	if selector.Type == models.SelectorTypeXpath {
		nodes, err := htmlquery.QueryAll(node, selector.Selector)
		if err != nil {
			return []*html.Node{}, models.NewParseError(models.ParserError, err.Error())
		}
		return nodes, nil
	}

	if selector.Type == models.SelectorTypeCss || selector.Type == models.SelectorTypeDefault {
		document := goquery.NewDocumentFromNode(node)
		return document.Find(selector.Selector).Nodes, nil
	}

	return nil, models.NewParseError(models.InternalError, "Unknown selector type"+selector.Type)
}

func queryHtmlFunction(selector *models.Selector, node *html.Node) (string, bool, *models.ParseError) {
	if selector == nil {
		return "", false, nil
	}
	if strings.TrimSpace(selector.Selector) == "" && selector.Type != models.SelectorTypeSelf && selector.Function == models.SelectorFunctionDefault {
		if selector.Param == "" && selector.Regex == "" && selector.DefaultValue == "" {
			return "", false, nil
		}
	}

	elements, err := queryHtmlElements(selector, node)
	if err != nil {
		return "", false, nil
	}
	for _, element := range elements {
		switch selector.Function {
		case models.SelectorFunctionAttr:
			for _, key := range strings.Split(selector.Param, ",") {
				value := htmlquery.SelectAttr(element, strings.TrimSpace(key))
				if value != "" {
					return value, true, nil
				}
			}
			if selector.DefaultValue != "" {
				return selector.DefaultValue, true, nil
			}
			return "", false, models.NewParseError(models.ElementNotFoundError, fmt.Sprintf("Seletor %s not found any %s attributes. reg: %s, replace: %s", selector.Selector, selector.Param, selector.Regex, selector.Replace))
		case models.SelectorFunctionText, models.SelectorFunctionDefault:
			return InnerText(element), true, nil
		case models.SelectorFunctionRaw:
			return htmlquery.OutputHTML(element, true), true, nil
		}
	}
	if selector.DefaultValue != "" {
		return selector.DefaultValue, true, nil
	}
	return "", false, models.NewParseError(models.ElementNotFoundError,
		fmt.Sprintf("Selector %s not found any elements", selector.Selector))
}

func InnerText(n *html.Node) string {
	var output func(*bytes.Buffer, *html.Node)
	output = func(buf *bytes.Buffer, n *html.Node) {
		switch n.Type {
		case html.TextNode:
			buf.WriteString(n.Data)
			return
		case html.CommentNode:
			return
		case html.ElementNode:
			if n.Data == "br" {
				buf.WriteString("\n")
			}
		}
		for child := n.FirstChild; child != nil; child = child.NextSibling {
			output(buf, child)
		}
	}

	var buf bytes.Buffer
	output(&buf, n)
	return buf.String()
}
