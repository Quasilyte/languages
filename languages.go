package languages

// List returns all known languages list.
func List() []Lang {
	return langList
}

// Lang describes a single formal language.
// This includes programming languages, markup languages and any other
// kinds of "computer" languages.
//
// Use type assertion to discover type-specific language details.
// Types that implement Lang:
//	- ProgLang
type Lang interface {
	// Name returns a full artificial language name.
	Name() string

	// URL is an http or https location that can be used to find
	// language information and/or documentation.
	// In the simplest case, it can be an English wikipedia page.
	// Permitted to return an empty string in case there is no known site.
	URL() string
}

// Note that fields from langBase are duplicated in every Lang implementation.
// This is intentional, to make initialization of language objects simpler.
// These duplicated fields are used to initialize langBase fields during init.
// langBase is needed to avoid Lang method implementations duplication
// (which would require to write dull doc-comments for every implementation).

// ProgLang is a programming language.
type ProgLang struct {
	langBase

	name string
	url  string
}

// langBase implements Lang interface methods.
// See interface definition itself for documentation on Name(), URL(), etc.
type langBase struct {
	name string
	url  string
}

func (l *langBase) Name() string { return l.name }

func (l *langBase) URL() string { return l.url }