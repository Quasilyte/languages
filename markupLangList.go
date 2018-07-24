package languages

import (
	"fmt"
)

func init() {
	type (
		langAttr interface{}

		langName string
		langURL  string
		langDate string
	)

	name := func(s string) langName { return langName(s) }
	url := func(s string) langURL { return langURL(s) }
	date := func(s string) langDate { return langDate(s) }

	markupLangList := [][]langAttr{
		{
			name("HTML"),
			url("https://en.wikipedia.org/wiki/HTML"),
			date("1993"),
		},
	}

	var utils initUtils
	// Push markup languages to the global list.
	for _, attrs := range markupLangList {
		var lang MarkupLang
		for _, attr := range attrs {
			switch v := attr.(type) {
			case langName:
				lang.name = string(v)
			case langURL:
				lang.url = string(v)
			case langDate:
				lang.date = utils.parseDate(string(v))
			default:
				panic(fmt.Sprintf("unexpected attr: %#v", v))
			}
		}
		langList = append(langList, &lang)
	}
}
