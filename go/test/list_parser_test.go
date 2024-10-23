package test

import (
	"catweb_parser/parsers"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestListParser(t *testing.T) {
	html, err := os.ReadFile("testdata/list_html.html")
	require.Nil(t, err)
	parserData, err := os.ReadFile("testdata/list_parser.json")
	require.Nil(t, err)

	parser := &parsers.ListViewParser{}
	err = json.Unmarshal(parserData, parser)
	require.Nil(t, err)

	require.NotEmpty(t, html)
	require.NotEmpty(t, parserData)

	result, err := parser.Parse(string(html))
	require.Nil(t, err)
	require.NotNil(t, result)
}
