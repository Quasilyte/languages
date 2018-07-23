package languages

import "time"

// List returns all known languages list.
//
// Languages are not sorted by their name.
// Use sort.Slice in combination with Lang.Name() method.
func List() []Lang {
	list := make([]Lang, len(langList))
	copy(list, langList)
	return list
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
	//
	// Permitted to return an empty string in case there is no known site.
	URL() string

	// FirstAppearance returns a date that can be interpreted as "creation time"
	// or "first public announce date".
	//
	// If first appearance date is unknown, zero time object is returned,
	// so Lang.FirstAppearance().IsZero() will return true.
	FirstAppearance() time.Time
}

// TypeCheckKind describes when type checking happens.
type TypeCheckKind int

// TypeCheckKind enum.
const (
	TypeCheckStatic TypeCheckKind = iota
	TypeCheckDynamic
	TypeCheckMixed
)

// TypeStrictnessKind describes how strict type system is.
// Usually, weak type systems allow implicit conversions while
// strong type systems require explicit conversions everywhere.
type TypeStrictnessKind int

// TypeStrictnessKind enum.
const (
	TypeStrictnessStrong TypeStrictnessKind = iota
	TypeStrictnessWeak
)

// Note that fields from langBase are duplicated in every Lang implementation.
// This is intentional, to make initialization of language objects simpler.
// These duplicated fields are used to initialize langBase fields during init.
// langBase is needed to avoid Lang method implementations duplication
// (which would require to write dull doc-comments for every implementation).

// ProgLang is a programming language.
// See progLangList.go.
type ProgLang struct {
	langBase

	name string
	url  string
	date string

	TypeCheck      TypeCheckKind
	TypeStrictness TypeStrictnessKind
}

// MarkupLang is a markup language.
// See markupLangList.go.
type MarkupLang struct {
	langBase

	name string
	url  string
	date string
}

// langBase implements Lang interface methods.
// See interface definition itself for documentation on Name(), URL(), etc.
type langBase struct {
	name string
	url  string
	date time.Time
}

func (l *langBase) Name() string               { return l.name }
func (l *langBase) URL() string                { return l.url }
func (l *langBase) FirstAppearance() time.Time { return l.date }
