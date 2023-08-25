package main

import (
	"catweb_parser/models"
	"catweb_parser/selector"
	"os"
	"testing"
)

func TestParser(t *testing.T) {
	itemSelector := &models.Selector{
		Selector: "//div[@class='gl1t'] | //*[starts-with(@class, 'gl2')]/..",
		Type:     models.SelectorTypeXpath,
	}

	tagItemSelector := &models.Selector{
		Selector: ".cn, .cs",
		Type:     models.SelectorTypeCss,
		Function: "text",
	}

	tagSelector := &models.TagSelector{
		Text: &models.Selector{
			Selector: "",
			Type:     "css",
			Function: "text",
			Param:    "",
			Regex:    "",
			Replace:  "",
			Script: &models.ScriptField{
				Script: "",
				Type:   "output",
			},
			DefaultValue: "",
		},
		Category: &models.Selector{
			Selector: "",
			Type:     "css",
			Function: "text",
			Param:    "",
			Regex:    "",
			Replace:  "",
			Script: &models.ScriptField{
				Script: "",
				Type:   "output",
			},
			DefaultValue: "",
		},
		Color: &models.Selector{
			Function: "attr",
			Param:    "class",
			Regex:    "ct.",
			Script: &models.ScriptField{
				Script: "a={ct2:\"#f66158\",ct3:\"#f09e19\",ct4:\"#d2d303\",ct5:\"#0fa911\",cta:\"#2fd92c\",ct9:\"#0bbfd3\",ct6:\"#4f5ce6\",ct7:\"#9030df\",ct8:\"#f38af2\",ct1:\"#8a8a8a\"},a[$arg];",
				Type:   "js",
			},
		},
	}

	html, err := os.ReadFile("test/list.htm")
	if err != nil {
		panic(err)
	}

	c, root, err := selector.CreateContext(string(html))
	if err != nil {
		panic(err)
	}

	nodes := c.Nodes(root, itemSelector)
	node := nodes[0]

	tags := c.Nodes(node, tagItemSelector)
	tag := tags[0]

	item := c.Tag(tag, tagSelector)

	if item.Text != nil {
		t.Log("Text", *item.Text)
	}

	if item.Color != nil {
		t.Log("Color", *item.Color)
	}

	if item.Category != nil {
		t.Log("Category", *item.Category)
	}
}
