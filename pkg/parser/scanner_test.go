package parser

import (
	"gotest.tools/assert"
	"testing"
)

func TestScanner_Set(t *testing.T) {

	tokens := Parse(`set {{ Domain }} "microspector.com" 
SET {{ Domain }} "microspector.com"`).All()

	assert.Equal(t, len(tokens), 14)
	assert.Equal(t, tokens[0].Type, SET)
	assert.Equal(t, tokens[1].Type, int('{'))
	assert.Equal(t, tokens[2].Type, int('{'))
	assert.Equal(t, tokens[3].Type, IDENTIFIER)
	assert.Equal(t, tokens[3].Val, "Domain")
	assert.Equal(t, tokens[4].Type, int('}'))
	assert.Equal(t, tokens[5].Type, int('}'))
	assert.Equal(t, tokens[6].Type, STRING)
	assert.Equal(t, tokens[6].Val, "microspector.com")

	tokens = Parse(`SET {{ Domain }} 100`).All()

	assert.Assert(t, len(tokens) == 7)
	assert.Equal(t, tokens[0].Type, SET)
	assert.Equal(t, tokens[1].Type, int('{'))
	assert.Equal(t, tokens[2].Type, int('{'))
	assert.Equal(t, tokens[3].Type, IDENTIFIER)
	assert.Equal(t, tokens[3].Val, "Domain")
	assert.Equal(t, tokens[4].Type, int('}'))
	assert.Equal(t, tokens[5].Type, int('}'))
	assert.Equal(t, tokens[6].Type, INTEGER)
	assert.Equal(t, tokens[6].Val, int64(100))

	tokens = Parse(`SET Domain 100`).All()

	assert.Assert(t, len(tokens) == 3)
	assert.Equal(t, tokens[0].Type, SET)
	assert.Equal(t, tokens[1].Type, IDENTIFIER)
	assert.Equal(t, tokens[1].Val, "Domain")
	assert.Equal(t, tokens[2].Type, INTEGER)
	assert.Equal(t, tokens[2].Val, int64(100))

}

func TestScanner_Must(t *testing.T) {

	tokens := Parse(`MUST {{ Domain }} EQUALS "microspector.com"`).All()

	assert.Assert(t, len(tokens) == 8)
	assert.Equal(t, tokens[0].Type, MUST)
	assert.Equal(t, tokens[1].Type, int('{'))
	assert.Equal(t, tokens[2].Type, int('{'))
	assert.Equal(t, tokens[3].Type, IDENTIFIER)
	assert.Equal(t, tokens[3].Val, "Domain")
	assert.Equal(t, tokens[4].Type, int('}'))
	assert.Equal(t, tokens[5].Type, int('}'))
	assert.Equal(t, tokens[6].Type, EQUALS)
	assert.Equal(t, tokens[7].Type, STRING)
	assert.Equal(t, tokens[7].Val, "microspector.com")

}

func TestScanner_Debug(t *testing.T) {
	tokens := Parse(`DEBUG {{ Domain }}`).All()

	assert.Assert(t, len(tokens) == 6)
	assert.Equal(t, tokens[0].Type, DEBUG)
	assert.Equal(t, tokens[1].Type, int('{'))
	assert.Equal(t, tokens[2].Type, int('{'))
	assert.Equal(t, tokens[3].Type, IDENTIFIER)
	assert.Equal(t, tokens[3].Val, "Domain")
	assert.Equal(t, tokens[4].Type, int('}'))
	assert.Equal(t, tokens[5].Type, int('}'))

	tokens = Parse(`DEBUG {{ Domain }} "microspector.com" {{ Domain }}`).All()

	assert.Assert(t, len(tokens) == 12)
	assert.Equal(t, tokens[0].Type, DEBUG)
	assert.Equal(t, tokens[1].Type, int('{'))
	assert.Equal(t, tokens[2].Type, int('{'))
	assert.Equal(t, tokens[3].Type, IDENTIFIER)
	assert.Equal(t, tokens[3].Val, "Domain")
	assert.Equal(t, tokens[4].Type, int('}'))
	assert.Equal(t, tokens[5].Type, int('}'))
	assert.Equal(t, tokens[6].Type, STRING)
	assert.Equal(t, tokens[6].Val, "microspector.com")
	assert.Equal(t, tokens[7].Type, int('{'))
	assert.Equal(t, tokens[8].Type, int('{'))
	assert.Equal(t, tokens[9].Type, IDENTIFIER)
	assert.Equal(t, tokens[9].Val, "Domain")
	assert.Equal(t, tokens[10].Type, int('}'))
	assert.Equal(t, tokens[11].Type, int('}'))
}
