package main

import (
	"catweb_parser/models"
	"catweb_parser/selector"
	"os"
	"testing"
)

func TestParser(t *testing.T) {
	s := &models.Selector{
		Selector: ".ir",
		Function: "attr",
		Param:    "style",
		Regex:    "background-position:-?(\\d+)px -?(\\d+)px",
		Replace:  "5-$1/16-($2-1)/40",
		Script: &models.ScriptField{
			Script: "eval($arg)",
			Type:   models.ScriptJs,
		},
	}
	html, err := os.ReadFile("./test_data/detail.html")
	if err != nil {
		t.Fatal(err)
	}

	c, root, err := selector.CreateContext(string(html))
	result := c.Double(root, s)
	if result == nil {
		t.Fatal(result)
	}
	println(*result)
}
