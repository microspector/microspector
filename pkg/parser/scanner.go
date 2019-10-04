package parser

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
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
	"INCLUDE":    INCLUDE,
	"SLEEP":      SLEEP,
	"IDENTIFIER": IDENTIFIER,
	"KEYWORD":    KEYWORD,
	"BODY":       BODY,
	"FOLLOW":     FOLLOW,
	"NOFOLLOW":   NOFOLLOW,
	"WHEN":       WHEN,
	"TRUE":       TRUE,
	"FALSE":      FALSE,
	"AND":        AND,
	"OR":         OR,
	"SET":        SET,
	"EQUALS":     EQUALS,
	"EQUAL":      EQUAL,
	"NOT":        NOT,
	"NOTEQUALS":  NOTEQUALS,
	"NOTEQUAL":   NOTEQUAL,
	"CONTAINS":   CONTAINS,
	"CONTAIN":    CONTAIN,
	"STARTSWITH": STARTSWITH,
	"STARTWITH":  STARTWITH,
	"IS":         IS,
	"ISNOT":      ISNOT,
	"MATCH":      MATCH,
	"MATCHES":    MATCHES,
	"MUST":       MUST,
	"SHOULD":     SHOULD,
	"LT":         LT,
	"GT":         GT,
	"DEBUG":      DEBUG,
	"ASSERT":     ASSERT,
	"END":        END,
	"NULL":       NULL,
	"ASYNC":      ASYNC,
	"CMD":        CMD,
	"IN":         IN,
	"INSECURE":   INSECURE,
	"SECURE":     SECURE,
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
reToken:
	ch := s.Peek()

	switch {
	case isSpace(ch):
		s.skipWhitespace()
		goto reToken
	case ch == '#':
		s.readUntilWith(isEndOfLine)
		goto reToken
	case isOperator(ch):
		return s.scanOperator()
	case isDigit(ch):
		return s.scanDigit()
	case isLetter(ch):
		return s.scanKeyword()
	case isQuote(ch):
		return s.scanQuotedString(ch)
	case ch == eof:
		return Token{
			Type: EOF,
			Val:  string(s.read()),
		}
	}

	return Token{
		Type: int(ch),
		Val:  string(s.read()),
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
func (s *Scanner) scanQuotedString(delimiter rune) (tok Token) {
	var buf bytes.Buffer
	s.read() //consume delimiter
	for {
		ch := s.read()

		if ch == eof {
			panic("unexpected end of file while scanning a string, maybe an unclosed quote?")
		}

		if ch == '\\' {
			if needsEscape(s.Peek(), delimiter) {
				switch s.read() {
				case 'n':
					buf.WriteRune('\n')
				case 'r':
					buf.WriteRune('\r')
				case 't':
					buf.WriteRune('\t')
				case '\\':
					buf.WriteRune('\\')
				case delimiter:
					buf.WriteRune(delimiter)
				}

				continue
			} else {
				panic("unescaped backslash")
			}
		}

		if ch == delimiter {
			//s.read() //consume delimiter
			break
		}

		buf.WriteRune(ch)
	}

	tok = Token{
		Type: STRING,
		Val:  buf.String(),
	}

	return tok
}

/**
* TODO:
* int8        the set of all signed  8-bit integers (-128 to 127)
* int16       the set of all signed 16-bit integers (-32768 to 32767)
* int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
* int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
 */
func (s *Scanner) scanDigit() (tok Token) {
	var buf bytes.Buffer
	foundDecimal := false
	for {
		ch := s.Peek()
		if isDigit(ch) {
			buf.WriteRune(s.read())
		} else if ch == '.' {
			if foundDecimal {
				//oh second dot? fuck.
				panic("i don't know what multiple dot means in a float val")
			}
			foundDecimal = true
			buf.WriteRune(s.read())
		} else {
			break
		}
	}

	if foundDecimal {
		f, _ := strconv.ParseFloat(buf.String(), 64)
		return Token{
			Type: FLOAT,
			Val:  f,
		}
	} else {
		i, _ := strconv.ParseInt(buf.String(), 10, 64)
		return Token{
			Type: INTEGER,
			Val:  i,
		}
	}
}

func (s *Scanner) scanKeyword() (tok Token) {
	if isVarStart(rune(s.Latest.Type)) {
		return Token{
			Type: IDENTIFIER,
			Val:  s.readWhile(isIdentifierChar),
		}
	}

	keyword := s.readWhile(isIdentifierChar)
	token, ok := keywords[strings.ToUpper(keyword)]
	if ok {
		return Token{
			Type: token,
			Val:  strings.ToUpper(keyword),
		}
	}

	if s.Latest.Type == IS || s.Latest.Type == ISNOT {
		return Token{
			Type: TYPE,
			Val:  keyword,
		}
	}

	return Token{
		Type: IDENTIFIER,
		Val:  keyword,
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
				Val:  "<=",
			}
		}

		return Token{
			Type: LT,
			Val:  "<",
		}
	case '>':

		if s.Peek() == '=' {
			s.read()
			return Token{
				Type: GE,
				Val:  ">=",
			}
		}

		return Token{
			Type: GT,
			Val:  ">",
		}

	case '!':
		if s.Peek() == '=' {
			s.read()
			return Token{
				Type: NOTEQUALS,
				Val:  "!=",
			}
		}
		fallthrough
	case '=':

		if s.Peek() == '=' {
			s.read()
			return Token{
				Type: EQUALS,
				Val:  "==",
			}
		}

		fallthrough
	default:
		return Token{
			Type: int(ch),
			Val:  string(ch),
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

func isQuote(ch rune) bool {
	return ch == '"' || ch == '\'' || ch == '`'
}

func needsEscape(ch, delim rune) bool {
	return ch == delim || ch == 'n' || ch == 't' || ch == '\\' || ch == 'r'
}

func isSpace(ch rune) bool {
	return ch == ' ' || ch == '\t' || isEndOfLine(ch)
}

func isEndOfLine(ch rune) bool {
	return ch == '\r' || ch == '\n'
}

func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}

func isLetter(ch rune) bool {
	return ch == '_' || unicode.IsLetter(ch)
}

func isAlphaNumeric(ch rune) bool {
	return ch == '_' || unicode.IsLetter(ch) || unicode.IsDigit(ch)
}

func isIdentifierChar(ch rune) bool {
	return ch == '_' || ch == '.' || unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '[' || ch == ']'
}

func isOperator(ch rune) bool {
	return ch == '<' || ch == '>' || ch == '=' || ch == '!'
}

func isVarStart(ch rune) bool {
	return ch == '{' || ch == '$'
}

func isVarEnd(ch rune) bool {
	return ch == '}'
}
