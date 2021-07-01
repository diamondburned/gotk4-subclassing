package girgen

import (
	"log"
	"regexp"
	"strings"

	"github.com/diamondburned/gotk4/gir"
)

// TODO: refactor to add method accuracy

// Preprocessor describes something that can preprocess anything in the given
// list of repositories. This is useful for renaming functions, classes or
// anything else.
type Preprocessor interface {
	Preprocess(repos gir.Repositories)
}

// FilterMatcher describes a filter for a GIR type.
type FilterMatcher interface {
	// Filter matches for the girType within the given namespace from the
	// namespace generator. The GIR type will never have a namespace prefix.
	Filter(*NamespaceGenerator, gir, c string) (keep bool)
}

type absoluteFilter struct {
	namespace string
	matcher   string
}

// AbsoluteFilter matches the names absolutely.
func AbsoluteFilter(abs string) FilterMatcher {
	parts := strings.SplitN(abs, ".", 2)
	if len(parts) != 2 {
		log.Panicf("missing namespace for AbsoluteFilter %q", abs)
	}

	return absoluteFilter{parts[0], parts[1]}
}

func (abs absoluteFilter) Filter(ng *NamespaceGenerator, names *FilterTypeName) (keep bool) {
	if abs.namespace == "C" {
		return names.CType != abs.matcher
	}

	typ, eq := EqNamespace(abs.namespace, names.GIRType)
	return !eq || typ != abs.matcher
}

type typeRenamer struct {
	namespace string
	from, to  string
}

// TypeRenamer creates a new filter matcher that renames a type. The given GIR
// type must contain the namespace, but the given name must not. The GIR type is
// absolutely matched, similarly to AbsoluteFilter.
func TypeRenamer(girType, newName string) FilterMatcher {
	parts := strings.SplitN(girType, ".", 2)
	if len(parts) != 2 {
		log.Panicf("missing namespace for TypeRenamer %q", girType)
	}
	if parts[0] == "C" {
		log.Panicf("TypeRenamer must not rename C type")
	}
	return typeRenamer{parts[0], parts[1], newName}
}

func (ren typeRenamer) Filter(ng *NamespaceGenerator, names *FilterTypeName) (keep bool) {
	typ, eq := EqNamespace(ren.namespace, names.GIRType)
	if eq && typ == ren.from {
		names.GIRType = ren.namespace + "." + ren.to
	}

	return true
}

type regexFilter struct {
	namespace string
	matcher   *regexp.Regexp
}

// RegexFilter returns a regex filter for FilterMatcher. A regex filter's format
// is a string consisting of two parts joined by a period: a namespace and a
// matcher. The only regex part is the matcher.
func RegexFilter(matcher string) FilterMatcher {
	parts := strings.SplitN(matcher, ".", 2)
	if len(parts) != 2 {
		log.Panicf("invalid regex filter format %q", matcher)
	}

	regex := regexp.MustCompile(wholeMatchRegex(parts[1]))

	return &regexFilter{
		namespace: parts[0],
		matcher:   regex,
	}
}

func wholeMatchRegex(regex string) string {
	// special regex
	if strings.Contains(regex, "(?") {
		return regex
	}

	// must whole match
	return "^" + regex + "$"
}

// Filter implements FilterMatcher.
func (rf *regexFilter) Filter(ng *NamespaceGenerator, names *FilterTypeName) (keep bool) {
	switch rf.namespace {
	case "C":
		return !rf.matcher.MatchString(names.CType)
	case "*":
		return !rf.matcher.MatchString(names.GIRType)
	}

	typ, eq := EqNamespace(rf.namespace, names.GIRType)
	if !eq {
		return true
	}

	return !rf.matcher.MatchString(typ)
}

// EqNamespace is used for FilterMatchers to compare types and namespaces.
func EqNamespace(nsp, girType string) (typ string, ok bool) {
	namespace, typ := gir.SplitGIRType(girType)
	return typ, namespace == nsp
}

type fileFilter struct {
	match *regexp.Regexp
}

// FileFilter filters based on the source position.
func FileFilter(regex string) FilterMatcher {
	return fileFilter{regexp.MustCompile(regex)}
}

func (ff fileFilter) Filter(ng *NamespaceGenerator, names *FilterTypeName) (keep bool) {
	res := ng.FindType(names.GIRType)
	if res == nil {
		return true
	}

	var source *gir.SourcePosition

	switch {
	case res.Alias != nil:
		source = res.Alias.SourcePosition
	case res.Class != nil:
		source = res.Class.SourcePosition
	case res.Interface != nil:
		source = res.Interface.SourcePosition
	case res.Record != nil:
		source = res.Record.SourcePosition
	case res.Enum != nil:
		source = res.Enum.SourcePosition
	case res.Function != nil:
		source = res.Function.SourcePosition
	case res.Union != nil:
		source = res.Union.SourcePosition
	case res.Bitfield != nil:
		source = res.Bitfield.SourcePosition
	case res.Callback != nil:
		source = res.Callback.SourcePosition
	}

	if source == nil {
		return true
	}

	return !ff.match.MatchString(source.Filename)
}