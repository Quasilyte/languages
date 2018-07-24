package languages

import (
	"fmt"
	"time"
)

// langList is an actual language list (well, slice) that includes all kinds
// of languages. Build during package init from all partial lists.
var langList []Lang

// initUtils is used as a function container that are used during knowledge base
// initialization process to avoid code duplication while still keeping global
// symbols amount low.
type initUtils struct{}

func (initUtils) parseDate(date string) time.Time {
	if date == "" {
		return time.Time{}
	}
	// Try year-only format first.
	t, err := time.Parse("2006", date)
	if err == nil {
		return t // OK, it was year-only
	}
	t, err = time.Parse("02 January 2006", date)
	if err != nil {
		panic(fmt.Sprintf("parse date: %v", err))
	}
	return t
}
