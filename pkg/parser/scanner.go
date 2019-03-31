package parser

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"unicode"
)

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{
		r: bufio.NewReader(r),
	}
}

func (s *Scanner) Scan() (tok Token) {
	ch := s.Peek()

	if isSpace(ch) {
		s.scanWhitespace() // consume and ignore whitespace
		ch = s.Peek()
	} else if isComma(ch) {
		s.read() // comsume comma
		ch = s.Peek()
	}

	if isVarStart(ch) {
		return s.scanVariable()
	} else if isArrayStart(ch) {
		return s.scanArray()
	} else if isOperator(ch) {
		return s.scanOperator()
	} else if isDigit(ch) {
		return s.scanDigit()
	} else if isLetter(ch) {
		return s.scanKeyword()
	} else if isDoubleQuote(ch) {
		return s.scanQuotedString()
	} else if isEndOfLine(ch) {
		return Token{
			Type: EOL,
			Text: string(s.read()),
		}
	} else if ch == eof {
		return Token{
			Type: EOF,
			Text: string(s.read()),
		}
	}

	return Token{
		Type: ILLEGAL,
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

func (s *Scanner) scanWhitespace() (tok Token) {
	return Token{
		Type: WHITESPACE,
		Text: s.readWhile(isSpace),
	}
}

func (s *Scanner) scanQuotedString() (tok Token) {

	var buf bytes.Buffer
	buf.WriteRune(s.read())
	var latestChar rune

	for {
		ch := s.read()
		if ch == eof || isEndOfLine(ch) {
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

	return Token{
		Type: KEYWORD,
		Text: s.readUntil(isSpace),
	}
}

func (s *Scanner) scanCommand() (tok Token) {
	return Token{
		Type: COMMAND,
		Text: s.readUntilWith(isEndOfLine),
	}
}

func (s *Scanner) scanVariable() (tok Token) {
	tok = Token{
		Type: VARIABLE,
		Text: s.readUntilWith(isVarEnd) + string(s.read()),
	}

	tok.Text = strings.TrimSpace(  strings.Trim(tok.Text, "{}") )

	return tok
}

func (s *Scanner) scanArray() (tok Token) {

	tok = Token{
		Type: ARRAY,
		Text: s.readUntilWith(isArrayEnd),
	}
	tok.Text = strings.Trim(tok.Text, "[]")
	tok.Tokenize()
	return
}

func (s *Scanner) scanOperator() (tok Token) {
	tok = Token{
		Type: OPERATOR,
		Text: s.readUntil(isSpace),
	}
	return
}

func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

var eof = rune(0)

func (s *Scanner) unread()        { _ = s.r.UnreadRune() }
func isDoubleQuote(ch rune) bool  { return ch == '"' }
func isSpace(ch rune) bool        { return ch == ' ' || ch == '\t' }
func isEndOfLine(ch rune) bool    { return ch == '\r' || ch == '\n' }
func isDigit(ch rune) bool        { return unicode.IsDigit(ch) }
func isLetter(ch rune) bool       { return ch == '_' || unicode.IsLetter(ch) }
func isAlphaNumeric(ch rune) bool { return ch == '_' || unicode.IsLetter(ch) || unicode.IsDigit(ch) }
func isOperator(ch rune) bool     { return ch == '<' || ch == '>' || ch == '=' || ch == '!' }
func isArrayStart(ch rune) bool   { return ch == '[' }
func isArrayEnd(ch rune) bool     { return ch == ']' }
func isVarStart(ch rune) bool     { return ch == '{' }
func isVarEnd(ch rune) bool       { return ch == '}' }
func isComma(ch rune) bool        { return ch == ',' }
