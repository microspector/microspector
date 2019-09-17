package parser

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"unicode"
)

type Scanner struct {
	r      *bufio.Reader
	Latest Token
}

var keywords = map[string]int{
	"EOL":        EOL,
	"EOF":        EOF,
	"GET":        GET,
	"HEAD":       HEAD,
	"POST":       POST,
	"PUT":        PUT,
	"DELETE":     DELETE,
	"CONNECT":    CONNECT,
	"OPTIONS":    OPTIONS,
	"TRACE":      TRACE,
	"HEADER":     HEADER,
	"INTO":       INTO,
	"HTTP":       HTTP,
	"IDENTIFIER": IDENTIFIER,
	"STRING":     STRING,
	"KEYWORD":    KEYWORD,
	"QUERY":      QUERY,
	"WHEN":       WHEN,
	"TRUE":       TRUE,
	"FALSE":      FALSE,
	"AND":        AND,
	"OR":         OR,
	"SET":        SET,
	"EQUALS":     EQUALS,
	"NOTEQUALS":  NOTEQUALS,
	"CONTAINS":   CONTAINS,
	"STARTSWITH": STARTSWITH,
	"MUST":       MUST,
	"SHOULD":     SHOULD,
	"LT":         LT,
	"GT":         GT,
	"DEBUG":      DEBUG,
	"ASSERT":     ASSERT,
	"END":        END,
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{
		r: bufio.NewReader(r),
	}
}

var keywordStrings = map[int]string{}

func init() {
	for str, id := range keywords {
		keywordStrings[id] = str
	}
}

func KeywordString(id int) string {
	str, ok := keywordStrings[id]
	if !ok {
		return ""
	}
	return str
}

func (s *Scanner) Scan() Token {
	s.Latest = s.getNextToken()
	return s.Latest
}

func (s *Scanner) getNextToken() Token {
	ch := s.Peek()

	if isSpace(ch) {
		s.skipWhitespace()
		ch = s.Peek()
	}

	// skip comments.
	if ch == '#' {
		s.readUntilWith(isEndOfLine)
		ch = s.Peek()
	}

	if isOperator(ch) {
		return s.scanOperator()
	} else if isDigit(ch) {
		return s.scanDigit()
	} else if isLetter(ch) {
		return s.scanKeyword()
	} else if isDoubleQuote(ch) {
		return s.scanQuotedString()
	} else if ch == eof {
		return Token{
			Type: EOF,
			Text: string(s.read()),
		}
	}
	return Token{
		Type: int(ch),
		Text: string(s.read()),
	}
}

func (s *Scanner) Peek() rune {
	r := s.read()
	s.unread()
	return r
}

func (s *Scanner) PeekPrev() rune {
	s.unread()
	r := s.read()
	return r
}

type runeCheck func(rune) bool

func (s *Scanner) readUntil(until runeCheck) string {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if until(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return buf.String()
}

func (s *Scanner) readUntilWith(until runeCheck) string {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if until(ch) {
			buf.WriteRune(ch)
			break
		} else {
			buf.WriteRune(ch)
		}
	}
	return buf.String()
}

func (s *Scanner) readWhile(while runeCheck) string {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		ch := s.read()
		if while(ch) {
			buf.WriteRune(ch)
		} else {
			s.unread()
			break
		}
	}
	// unread the latest char we consume.
	return buf.String()
}

func (s *Scanner) skipWhitespace() {
	s.readWhile(isSpace)
}

func (s *Scanner) skipEndOfLine() {
	s.readUntilWith(isEndOfLine)
}

/**
\” – To escape “ within double quoted string.
\\ – To escape the backslash.
\n – To add line breaks between string.
\t – To add tab space.
\r – For carriage return.
*/
func (s *Scanner) scanQuotedString() (tok Token) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())
	for {
		ch := s.read()

		if ch == eof {
			panic("unexpected end of file while scanning a string, maybe an unclosed quote?")
		}

		if ch == '\\' {
			if needsEscape(s.Peek()) {
				switch s.read() {
				case 'n':
					buf.WriteRune('\n')
				case 'r':
					buf.WriteRune('\r')
				case 't':
					buf.WriteRune('\t')
				case '\\':
					buf.WriteRune('\\')
				case '"':
					buf.WriteRune('"')
				}

				continue
			} else {
				panic("unescaped backslash")
			}
		}

		if isDoubleQuote(ch) {
			buf.WriteRune(ch)
			break
		}

		buf.WriteRune(ch)
	}

	tok = Token{
		Type: STRING,
		Text: strings.Trim(buf.String(), "\""),
	}

	return tok
}

func (s *Scanner) scanDigit() (tok Token) {
	return Token{
		Type: INTEGER,
		Text: s.readWhile(isDigit),
	}
}

func (s *Scanner) scanKeyword() (tok Token) {
	if isVarStart(rune(s.Latest.Type)) {
		return Token{
			Type: IDENTIFIER,
			Text: s.readWhile(isIdentifierChar),
		}
	}

	keyword := s.readWhile(isLetter)
	token, ok := keywords[keyword]
	if ok {
		return Token{
			Type: token,
			Text: keyword,
		}
	}
	return Token{
		Type: KEYWORD,
		Text: keyword,
	}
}

func (s *Scanner) scanOperator() (tok Token) {
	ch := s.read()
	switch ch {
	case '<':
		if s.Peek() == '=' {
			s.read()
			return Token{
				Type: LE,
				Text: "<=",
			}
		}

		return Token{
			Type: LT,
			Text: "<",
		}
	case '>':

		if s.Peek() == '=' {
			s.read()
			return Token{
				Type: GE,
				Text: ">=",
			}
		}

		return Token{
			Type: GT,
			Text: ">",
		}

	case '!':
		if s.Peek() == '=' {
			s.read()
			return Token{
				Type: NOTEQUALS,
				Text: "!=",
			}
		}
		fallthrough
	case '=':

		if s.Peek() == '=' {
			s.read()
			return Token{
				Type: EQUALS,
				Text: "==",
			}
		}

		fallthrough

	default:
		return Token{
			Type: int(ch),
			Text: string(ch),
		}
	}

}

func (s *Scanner) unread() {
	_ = s.r.UnreadRune()
}

func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof

	}

	return ch
}

var eof = rune(0)

func isDoubleQuote(ch rune) bool    { return ch == '"' }
func needsEscape(ch rune) bool      { return ch == 'n' || ch == 't' || ch == '"' || ch == '\\' || ch == 'r' }
func isSpace(ch rune) bool          { return ch == ' ' || ch == '\t' || isEndOfLine(ch) }
func isEndOfLine(ch rune) bool      { return ch == '\r' || ch == '\n' }
func isDigit(ch rune) bool          { return unicode.IsDigit(ch) }
func isLetter(ch rune) bool         { return ch == '_' || unicode.IsLetter(ch) }
func isAlphaNumeric(ch rune) bool   { return ch == '_' || unicode.IsLetter(ch) || unicode.IsDigit(ch) }
func isIdentifierChar(ch rune) bool { return ch == '_' || ch == '.' || unicode.IsLetter(ch) || unicode.IsDigit(ch) }
func isOperator(ch rune) bool       { return ch == '<' || ch == '>' || ch == '=' || ch == '!' }
func isVarStart(ch rune) bool       { return ch == '{' }
func isVarEnd(ch rune) bool         { return ch == '}' }
