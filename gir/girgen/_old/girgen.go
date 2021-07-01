package girgen

import (
	"log"
	"path/filepath"
	"runtime"
	"strconv"
	"text/template"

	"github.com/diamondburned/gotk4/gir"
	"github.com/fatih/color"
)

func newGoTemplate(block string) *template.Template {
	_, file, _, _ := runtime.Caller(1)
	base := filepath.Base(file)

	t := template.New(base)
	t.Funcs(template.FuncMap{
		"PascalToGo":     PascalToGo,
		"UnexportPascal": UnexportPascal,
		"SnakeToGo":      SnakeToGo,
		"FirstLetter":    FirstLetter,
		"GoDoc":          GoDoc,
	})
	t = template.Must(t.Parse(block))
	return t
}

// Generator is a big generator that manages multiple repositories.
type Generator struct {
	Repos   gir.Repositories
	ModPath ModulePathFunc
	Filters []FilterMatcher

	Logger   *log.Logger
	LogLevel LogLevel
	Color    bool
}

// ModulePathFunc returns the Go module import path from the given namespace.
type ModulePathFunc func(*gir.Namespace) string

// NewGenerator creates a new generator with sane defaults.
func NewGenerator(repos gir.Repositories, modPath ModulePathFunc) *Generator {
	return &Generator{
		Repos:   repos,
		ModPath: modPath,
		Filters: []FilterMatcher{
			// These are already manually covered in the girgen code; they are
			// provided by package gotk3/glib.
			AbsoluteFilter("GLib.Error"),
			// Ignore generating everything in GObject, but allow resolving its
			// types.
			RegexFilter("GObject..*"),

			// This is not supported by Go. We might be able to support it in
			// the future using a 16-byte data structure, but the actual size
			// isn't well defined as far as I know.
			AbsoluteFilter("*.long double"),

			// Special marking for internal types from GLib (apparently for
			// glib:get-type).
			AbsoluteFilter("C.intern"),
		},
		LogLevel: LogUnsupported,
	}
}

// AddFilters adds the given list of filters.
func (g *Generator) AddFilters(filters []FilterMatcher) {
	g.Filters = append(g.Filters, filters...)
}

// UseNamespace creates a new namespace generator using the given namespace.
func (g *Generator) UseNamespace(namespace, version string) *NamespaceGenerator {
	res := g.Repos.FindNamespace(gir.VersionedName(namespace, version))
	if res == nil {
		return nil
	}

	return NewNamespaceGenerator(g, res)
}

type LogLevel uint8

const (
	LogDebug LogLevel = iota
	// LogUnsupported is used for Logs that report conditions impossible to do
	// in Go properly.
	LogUnsupported
	// LogUnknown is reserved for Logging down unknown types or types that
	// cannot be resolved.
	LogUnknown
	// LogSkip is reserved for Logging down skipped types.
	LogSkip
	// LogUnusuality is reserved for Logging down unexpected GIR values or
	// formats. It may be used to Log things yet to be supported but can be.
	LogUnusuality
	// LogError is reserved for fatal and/or unexpected errors.
	LogError
)

func (lvl LogLevel) prefix() string {
	switch lvl {
	case LogDebug:
		return "debug:"
	case LogUnsupported:
		return "unsupported:"
	case LogUnknown:
		return "unknown type:"
	case LogSkip:
		return "skipped:"
	case LogUnusuality:
		return "unusuality:"
	case LogError:
		return "error:"
	default:
		return ""
	}
}

func (lvl LogLevel) colorf(f string, v ...interface{}) string {
	switch lvl {
	case LogUnsupported:
		return color.YellowString(f, v...)
	case LogUnknown:
		return color.BlueString(f, v...)
	case LogSkip:
		return color.GreenString(f, v...)
	case LogUnusuality:
		return color.RedString(f, v...)
	case LogError:
		return color.New(color.Bold, color.FgHiRed).Sprintf(f, v...)
	case LogDebug:
		fallthrough
	default:
		return color.New(color.Faint).Sprintf(f, v...)
	}
}

type LineLogger interface {
	Logln(LogLevel, ...interface{})
}

var (
	_ LineLogger = (*Generator)(nil)
	_ LineLogger = (*NamespaceGenerator)(nil)
)

// Logln Logs using the Logger.
func (g *Generator) Logln(level LogLevel, v ...interface{}) {
	if g.Logger == nil || g.LogLevel > level {
		return
	}

	prefix := level.prefix()
	if prefix != "" {
		if g.Color {
			prefix = level.colorf(prefix)
		}
		v = append(v, nil)
		copy(v[1:], v)
		v[0] = prefix
	}

	g.Logger.Println(v...)
}

func tryLogln(logger LineLogger, level LogLevel, v ...interface{}) {
	if logger == nil {
		// Intentionally nil.
		return
	}

	logger.Logln(level, v...)
}

func warnUnknownType(logger LineLogger, typ string) {
	tryLogln(logger, LogUnknown, strconv.Quote(typ))
}