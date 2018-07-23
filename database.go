package languages

var langList = []Lang{
	&ProgLang{
		name: "Ada",
		url:  "https://en.wikipedia.org/wiki/Ada_(programming_language)",
	},
}

func init() {
	for _, l := range langList {
		switch l := l.(type) {
		case *ProgLang:
			l.langBase = langBase{name: l.name, url: l.url}
		}
	}
}
