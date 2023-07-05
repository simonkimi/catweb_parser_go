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
	if selector.Selector == "" {
		return []*html.Node{node}, nil
	}
	if selector.Type == models.SelectorTypeCss {
		document := goquery.NewDocumentFromNode(node)
		return document.Find(selector.Selector).Nodes, nil
	}

	if selector.Type == models.SelectorTypeXpath {
		nodes, err := htmlquery.QueryAll(node, selector.Selector)
		if err != nil {
			return []*html.Node{}, models.NewParseError(models.ParserError, err.Error())
		}
		return nodes, nil
	}

	return nil, models.NewParseError(models.InternalError, "Unknown selector type"+selector.Type)
}

func queryHtmlFunction(selector *models.Selector, node *html.Node) (string, bool, *models.ParseError) {
	if selector.Selector == "" {
		if selector.Function == models.SelectorFunctionAuto && selector.Param == "" && selector.Regex == "" && selector.DefaultValue == "" {
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
			for _, key := range strings.Split(selector.Param, "") {
				value := htmlquery.SelectAttr(element, strings.TrimSpace(key))
				if value != "" {
					return value, true, nil
				}
			}
			if selector.DefaultValue != "" {
				return selector.DefaultValue, true, nil
			}
			return "", false, models.NewParseError(models.ElementNotFoundError, fmt.Sprintf("Seletor %s not found any attributes", selector.Selector))
		case models.SelectorFunctionText, models.SelectorFunctionAuto:
			return htmlquery.OutputHTML(element, true), true, nil
		case models.SelectorFunctionRaw:
			return innerText(element), true, nil
		}
	}
	if selector.DefaultValue != "" {
		return selector.DefaultValue, true, nil
	}
	return "", false, models.NewParseError(models.ElementNotFoundError,
		fmt.Sprintf("Selector %s not found any elements", selector.Selector))
}

func innerText(n *html.Node) string {
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
