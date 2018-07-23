package languages

import (
	"fmt"
	"time"
)

// langList is an actual language list (well, slice) that includes all kinds
// of languages. Build during package init from all partial lists.
var langList []Lang

func init() {
	// Build combined list out of all partial lists.
	for _, l := range progLangList {
		langList = append(langList, l)
	}
	for _, l := range markupLangList {
		langList = append(langList, l)
	}
	// Now get rid of partial slices.
	progLangList = nil
	markupLangList = nil

	parseDate := func(name, date string) time.Time {
		if date == "" {
			return time.Time{}
		}
		t, err := time.Parse("02 January 2006", date)
		if err != nil {
			panic(fmt.Sprintf("parse %s date: %v", name, err))
		}
		return t
	}

	for _, l := range langList {
		switch l := l.(type) {
		case *ProgLang:
			l.langBase = langBase{
				name: l.name,
				url:  l.url,
				date: parseDate(l.name, l.date),
			}
		case *MarkupLang:
			l.langBase = langBase{
				name: l.name,
				url:  l.url,
				date: parseDate(l.name, l.date),
			}
		}
	}
}
