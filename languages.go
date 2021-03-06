package languages

import "time"

// List returns all known languages list.
//
// Languages are not sorted by their name (order is unspecified).
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
//	- MarkupLang
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

// TypeSystem describes programming language type system properties.
type TypeSystem struct {
	Checking   TypeChecking
	Strictness TypeStrictness
}

// TypeChecking describes when type checking happens.
type TypeChecking int

// TypeChecking enum.
const (
	TypeCheckingStatic TypeChecking = iota
	TypeCheckingDynamic
	TypeCheckingMixed
)

// TypeStrictness describes how strict type system is.
// Usually, weak type systems allow implicit conversions while
// strong type systems require explicit conversions everywhere.
type TypeStrictness int

// TypeStrictness enum.
const (
	TypeStrictnessStrong TypeStrictness = iota
	TypeStrictnessWeak
)

// ProgLang is a programming language.
// See progLangList.go.
type ProgLang struct {
	langBase

	TypeSystem TypeSystem
}

// MarkupLang is a markup language.
// See markupLangList.go.
type MarkupLang struct {
	langBase
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
