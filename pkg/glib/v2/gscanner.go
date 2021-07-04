// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib.h>
import "C"

// ErrorType: the possible errors, used in the @v_error field of Value, when the
// token is a G_TOKEN_ERROR.
type ErrorType int

const (
	// unknown: unknown error
	ErrorTypeUnknown ErrorType = 0
	// UnexpEOF: unexpected end of file
	ErrorTypeUnexpEOF ErrorType = 1
	// UnexpEOFInString: unterminated string constant
	ErrorTypeUnexpEOFInString ErrorType = 2
	// UnexpEOFInComment: unterminated comment
	ErrorTypeUnexpEOFInComment ErrorType = 3
	// NonDigitInConst: non-digit character in a number
	ErrorTypeNonDigitInConst ErrorType = 4
	// DigitRadix: digit beyond radix in a number
	ErrorTypeDigitRadix ErrorType = 5
	// FloatRadix: non-decimal floating point number
	ErrorTypeFloatRadix ErrorType = 6
	// FloatMalformed: malformed floating point number
	ErrorTypeFloatMalformed ErrorType = 7
)

// TokenType: the possible types of token returned from each
// g_scanner_get_next_token() call.
type TokenType int

const (
	// eof: the end of the file
	TokenTypeEOF TokenType = 0
	// LeftParen: a '(' character
	TokenTypeLeftParen TokenType = 40
	// RightParen: a ')' character
	TokenTypeRightParen TokenType = 41
	// LeftCurly: a '{' character
	TokenTypeLeftCurly TokenType = 123
	// RightCurly: a '}' character
	TokenTypeRightCurly TokenType = 125
	// LeftBrace: a '[' character
	TokenTypeLeftBrace TokenType = 91
	// RightBrace: a ']' character
	TokenTypeRightBrace TokenType = 93
	// EqualSign: a '=' character
	TokenTypeEqualSign TokenType = 61
	// comma: a ',' character
	TokenTypeComma TokenType = 44
	// none: not a token
	TokenTypeNone TokenType = 256
	// error: an error occurred
	TokenTypeError TokenType = 257
	// char: a character
	TokenTypeChar TokenType = 258
	// binary: a binary integer
	TokenTypeBinary TokenType = 259
	// octal: an octal integer
	TokenTypeOctal TokenType = 260
	// int: an integer
	TokenTypeInt TokenType = 261
	// hex: a hex integer
	TokenTypeHex TokenType = 262
	// float: a floating point number
	TokenTypeFloat TokenType = 263
	// string: a string
	TokenTypeString TokenType = 264
	// symbol: a symbol
	TokenTypeSymbol TokenType = 265
	// identifier: an identifier
	TokenTypeIdentifier TokenType = 266
	// IdentifierNull: a null identifier
	TokenTypeIdentifierNull TokenType = 267
	// CommentSingle: one line comment
	TokenTypeCommentSingle TokenType = 268
	// CommentMulti: multi line comment
	TokenTypeCommentMulti TokenType = 269
)

// ScannerConfig specifies the #GScanner parser configuration. Most settings can
// be changed during the parsing phase and will affect the lexical parsing of
// the next unpeeked token.
type ScannerConfig C.GScannerConfig

// WrapScannerConfig wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapScannerConfig(ptr unsafe.Pointer) *ScannerConfig {
	return (*ScannerConfig)(ptr)
}

// Native returns the underlying C source pointer.
func (s *ScannerConfig) Native() unsafe.Pointer {
	return unsafe.Pointer(s)
}