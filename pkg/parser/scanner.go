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
	"CONTAINS":   CONTAINS,
	"STARTSWITH": STARTSWITH,
	"MUST":       MUST,
	"SHOULD":     SHOULD,
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

	for isSpace(ch) || isEndOfLine(ch) {

		if isSpace(ch) {
			s.skipWhitespace() // consume and ignore whitespace
			ch = s.Peek()
		}

		if isEndOfLine(ch) {
			s.skipEndOfLine()
			ch = s.Peek()
		}
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
		ch := s.read();
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
	s.readWhile(isEndOfLine)
}

func (s *Scanner) scanQuotedString() (tok Token) {

	var buf bytes.Buffer
	buf.WriteRune(s.read())
	var latestChar rune

	for {
		ch := s.read()
		if ch == eof {
			break
		} else if ch == '\\' {
			latestChar = ch
			if s.Peek() != '"' {
				buf.WriteRune(ch)
			}
		} else if isDoubleQuote(ch) {
			if latestChar == '\\' {
				buf.WriteRune(ch)
			} else {
				buf.WriteRune(ch)
				break
			}
		} else {
			latestChar = ch
			buf.WriteRune(ch)
		}
	}

	tok = Token{
		Type: STRING,
		Text: buf.String(),
	}

	tok.Text = strings.Trim(tok.Text, "\"")

	return tok
}

func (s *Scanner) scanDigit() (tok Token) {
	return Token{
		Type: INTEGER,
		Text: s.readUntil(isSpace),
	}
}

func (s *Scanner) scanKeyword() (tok Token) {

	//TODO: scan identifier here, identifiers support dot to reach any json value like HttpResult.data.success etc.
	if isVarStart(rune(s.Latest.Type)) {
		return Token{
			Type: IDENTIFIER,
			Text: s.readWhile(isLetter),
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
	tok = Token{
		Type: int(ch),
		Text: string(ch),
	}
	return
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

func isDoubleQuote(ch rune) bool  { return ch == '"' }
func isSpace(ch rune) bool        { return ch == ' ' || ch == '\t' }
func isEndOfLine(ch rune) bool    { return ch == '\r' || ch == '\n' }
func isDigit(ch rune) bool        { return unicode.IsDigit(ch) }
func isLetter(ch rune) bool       { return ch == '_' || unicode.IsLetter(ch) }
func isAlphaNumeric(ch rune) bool { return ch == '_' || unicode.IsLetter(ch) || unicode.IsDigit(ch) }
func isOperator(ch rune) bool     { return ch == '<' || ch == '>' || ch == '=' || ch == '!' }
func isVarStart(ch rune) bool     { return ch == '{' }
func isVarEnd(ch rune) bool       { return ch == '}' }
