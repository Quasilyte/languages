package languages

var progLangList = []*ProgLang{
	{
		name: "Ada",
		url:  "https://en.wikipedia.org/wiki/Ada_(programming_language)",
	},
}

var langList []Lang

func init() {
	// Build combined list out of all partial lists.
	for _, l := range progLangList {
		langList = append(langList, l)
	}
	// Now get rid of partial slices.
	progLangList = nil

	for _, l := range langList {
		switch l := l.(type) {
		case *ProgLang:
			l.langBase = langBase{name: l.name, url: l.url}
		}
	}
}
