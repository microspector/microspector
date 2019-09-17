package parser

import (
	"gotest.tools/assert"
	"testing"
)

func TestScanner_Set(t *testing.T) {

	l := Parse(`SET {{ Domain }} "microspector.com"`)

	assert.Assert(t, len(l.tokens) == 7)
	assert.Equal(t, l.tokens[0].Type, SET)
	assert.Equal(t, l.tokens[1].Type, int('{'))
	assert.Equal(t, l.tokens[2].Type, int('{'))
	assert.Equal(t, l.tokens[3].Type, IDENTIFIER)
	assert.Equal(t, l.tokens[3].Text, "Domain")
	assert.Equal(t, l.tokens[4].Type, int('}'))
	assert.Equal(t, l.tokens[5].Type, int('}'))
	assert.Equal(t, l.tokens[6].Type, STRING)
	assert.Equal(t, l.tokens[6].Text, "microspector.com")


	l = Parse(`SET {{ Domain }} 100`)

	assert.Assert(t, len(l.tokens) == 7)
	assert.Equal(t, l.tokens[0].Type, SET)
	assert.Equal(t, l.tokens[1].Type, int('{'))
	assert.Equal(t, l.tokens[2].Type, int('{'))
	assert.Equal(t, l.tokens[3].Type, IDENTIFIER)
	assert.Equal(t, l.tokens[3].Text, "Domain")
	assert.Equal(t, l.tokens[4].Type, int('}'))
	assert.Equal(t, l.tokens[5].Type, int('}'))
	assert.Equal(t, l.tokens[6].Type, INTEGER)
	assert.Equal(t, l.tokens[6].Text, "100")

}

func TestScanner_Must(t *testing.T) {

	l := Parse(`MUST {{ Domain }} EQUALS "microspector.com"`)

	assert.Assert(t, len(l.tokens) == 8)
	assert.Equal(t, l.tokens[0].Type, MUST)
	assert.Equal(t, l.tokens[1].Type, int('{'))
	assert.Equal(t, l.tokens[2].Type, int('{'))
	assert.Equal(t, l.tokens[3].Type, IDENTIFIER)
	assert.Equal(t, l.tokens[3].Text, "Domain")
	assert.Equal(t, l.tokens[4].Type, int('}'))
	assert.Equal(t, l.tokens[5].Type, int('}'))
	assert.Equal(t, l.tokens[6].Type, EQUALS)
	assert.Equal(t, l.tokens[7].Type, STRING)
	assert.Equal(t, l.tokens[7].Text, "microspector.com")

}

func TestScanner_Debug(t *testing.T) {
	l := Parse(`DEBUG {{ Domain }}`)

	assert.Assert(t, len(l.tokens) == 6)
	assert.Equal(t, l.tokens[0].Type, DEBUG)
	assert.Equal(t, l.tokens[1].Type, int('{'))
	assert.Equal(t, l.tokens[2].Type, int('{'))
	assert.Equal(t, l.tokens[3].Type, IDENTIFIER)
	assert.Equal(t, l.tokens[3].Text, "Domain")
	assert.Equal(t, l.tokens[4].Type, int('}'))
	assert.Equal(t, l.tokens[5].Type, int('}'))

	l2 := Parse(`DEBUG {{ Domain }} "microspector.com" {{ Domain }}`)

	assert.Assert(t, len(l2.tokens) == 12)
	assert.Equal(t, l2.tokens[0].Type, DEBUG)
	assert.Equal(t, l2.tokens[1].Type, int('{'))
	assert.Equal(t, l2.tokens[2].Type, int('{'))
	assert.Equal(t, l2.tokens[3].Type, IDENTIFIER)
	assert.Equal(t, l2.tokens[3].Text, "Domain")
	assert.Equal(t, l2.tokens[4].Type, int('}'))
	assert.Equal(t, l2.tokens[5].Type, int('}'))
	assert.Equal(t, l2.tokens[6].Type, STRING)
	assert.Equal(t, l2.tokens[6].Text, "microspector.com")
	assert.Equal(t, l2.tokens[7].Type, int('{'))
	assert.Equal(t, l2.tokens[8].Type, int('{'))
	assert.Equal(t, l2.tokens[9].Type, IDENTIFIER)
	assert.Equal(t, l2.tokens[9].Text, "Domain")
	assert.Equal(t, l2.tokens[10].Type, int('}'))
	assert.Equal(t, l2.tokens[11].Type, int('}'))
}
