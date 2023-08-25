package selector

import (
	"catweb_parser/models"
	"fmt"
	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/oj"
	"strconv"
)

func queryJsonElements(selector *models.Selector, node any) ([]any, *models.ParseError) {
	expr, err := jp.ParseString(selector.Selector)
	if err != nil {
		return nil, models.NewParseError(models.ParserError, err.Error())
	}
	return expr.Get(node), nil
}

func queryJsonFunction(selector *models.Selector, node any) (string, bool, *models.ParseError) {
	if selector == nil || selector.IsEmpty() {
		return "", false, nil
	}
	elements, err := queryJsonElements(selector, node)
	if err != nil {
		return "", false, nil
	}

	for _, element := range elements {
		switch element.(type) {
		case string:
			return element.(string), true, nil
		case int:
			return strconv.Itoa(element.(int)), true, nil
		case float64:
			return strconv.FormatFloat(element.(float64), 'f', -1, 64), true, nil
		case bool:
			return strconv.FormatBool(element.(bool)), true, nil
		case map[string]any:
			return oj.JSON(element), true, nil
		}
	}

	return "", false, models.NewParseError(models.ElementNotFoundError, fmt.Sprintf("Element not found: %s", selector.Selector))
}
