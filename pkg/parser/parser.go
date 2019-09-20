// Code generated by goyacc - DO NOT EDIT.

package parser

import __yyfmt__ "fmt"

import (
	"log"
	"strings"
)

var GlobalVars = map[string]interface{}{}

type yySymType struct {
	yys      int
	val      interface{}
	vals     []interface{}
	str      string
	integer  int64
	boolean  bool
	bytes    []byte
	cmd      Command
	variable struct {
		name  string
		value interface{}
	}
	http_command_params []HttpCommandParam
	http_command_param  HttpCommandParam
}

type yyXError struct {
	state, xsym int
}

const (
	yyDefault  = 57395
	yyEofCode  = 57344
	AND        = 57386
	ASSERT     = 57360
	BODY       = 57373
	CONNECT    = 57368
	CONTAIN    = 57383
	CONTAINS   = 57382
	DEBUG      = 57358
	DELETE     = 57367
	END        = 57359
	EOF        = 57347
	EOL        = 57346
	EQUAL      = 57375
	EQUALS     = 57374
	FALSE      = 57352
	FLOAT      = 57350
	GE         = 57379
	GET        = 57363
	GT         = 57378
	HEAD       = 57364
	HEADER     = 57372
	HTTP       = 57355
	IDENTIFIER = 57393
	INCLUDE    = 57361
	INTEGER    = 57349
	INTO       = 57392
	IS         = 57390
	ISNOT      = 57391
	KEYWORD    = 57353
	LE         = 57381
	LT         = 57380
	MATCH      = 57389
	MATCHES    = 57388
	MUST       = 57356
	NOTEQUAL   = 57377
	NOTEQUALS  = 57376
	OPTIONS    = 57369
	OR         = 57387
	PATCH      = 57371
	POST       = 57365
	PUT        = 57366
	SET        = 57354
	SHOULD     = 57357
	SLEEP      = 57362
	STARTSWITH = 57384
	STRING     = 57348
	TRACE      = 57370
	TRUE       = 57351
	TYPE       = 57394
	WHEN       = 57385
	yyErrCode  = 57345

	yyMaxDepth = 200
	yyTabOfs   = -87
)

var (
	yyPrec = map[int]int{
		'|': 0,
		'&': 1,
		'+': 2,
		'-': 2,
		'*': 3,
		'/': 3,
		'%': 3,
	}

	yyXLAT = map[int]int{
		57393: 0,  // IDENTIFIER (65x)
		57344: 1,  // $end (64x)
		123:   2,  // '{' (64x)
		57360: 3,  // ASSERT (64x)
		57358: 4,  // DEBUG (64x)
		57359: 5,  // END (64x)
		57355: 6,  // HTTP (64x)
		57361: 7,  // INCLUDE (64x)
		57356: 8,  // MUST (64x)
		57354: 9,  // SET (64x)
		57357: 10, // SHOULD (64x)
		57362: 11, // SLEEP (64x)
		36:    12, // '$' (63x)
		57350: 13, // FLOAT (61x)
		57349: 14, // INTEGER (61x)
		57348: 15, // STRING (61x)
		57394: 16, // TYPE (61x)
		57385: 17, // WHEN (59x)
		57392: 18, // INTO (58x)
		40:    19, // '(' (37x)
		57386: 20, // AND (37x)
		57387: 21, // OR (37x)
		41:    22, // ')' (32x)
		57352: 23, // FALSE (30x)
		57351: 24, // TRUE (30x)
		57419: 25, // variable (28x)
		57397: 26, // any_value (26x)
		57413: 27, // number (26x)
		37:    28, // '%' (25x)
		42:    29, // '*' (25x)
		43:    30, // '+' (25x)
		45:    31, // '-' (25x)
		47:    32, // '/' (25x)
		57383: 33, // CONTAIN (20x)
		57382: 34, // CONTAINS (20x)
		57375: 35, // EQUAL (20x)
		57374: 36, // EQUALS (20x)
		57379: 37, // GE (20x)
		57378: 38, // GT (20x)
		57390: 39, // IS (20x)
		57391: 40, // ISNOT (20x)
		57381: 41, // LE (20x)
		57380: 42, // LT (20x)
		57389: 43, // MATCH (20x)
		57388: 44, // MATCHES (20x)
		57377: 45, // NOTEQUAL (20x)
		57376: 46, // NOTEQUALS (20x)
		57384: 47, // STARTSWITH (20x)
		57404: 48, // expr (19x)
		57373: 49, // BODY (15x)
		57372: 50, // HEADER (15x)
		57418: 51, // true_false (12x)
		57399: 52, // boolean_exp (11x)
		57405: 53, // expr_opr (11x)
		57414: 54, // operator (4x)
		125:   55, // '}' (2x)
		57407: 56, // http_command_param (2x)
		57396: 57, // any_command (1x)
		57398: 58, // assert_command (1x)
		57400: 59, // command (1x)
		57401: 60, // command_with_condition_opt (1x)
		57368: 61, // CONNECT (1x)
		57402: 62, // debug_command (1x)
		57367: 63, // DELETE (1x)
		57403: 64, // end_command (1x)
		57363: 65, // GET (1x)
		57364: 66, // HEAD (1x)
		57406: 67, // http_command (1x)
		57408: 68, // http_command_params (1x)
		57409: 69, // http_method (1x)
		57410: 70, // include_command (1x)
		57411: 71, // multi_any_value (1x)
		57412: 72, // must_command (1x)
		57369: 73, // OPTIONS (1x)
		57371: 74, // PATCH (1x)
		57365: 75, // POST (1x)
		57366: 76, // PUT (1x)
		57415: 77, // set_command (1x)
		57416: 78, // should_command (1x)
		57417: 79, // sleep_command (1x)
		57370: 80, // TRACE (1x)
		57395: 81, // $default (0x)
		38:    82, // '&' (0x)
		124:   83, // '|' (0x)
		57347: 84, // EOF (0x)
		57346: 85, // EOL (0x)
		57345: 86, // error (0x)
		57353: 87, // KEYWORD (0x)
	}

	yySymNames = []string{
		"IDENTIFIER",
		"$end",
		"'{'",
		"ASSERT",
		"DEBUG",
		"END",
		"HTTP",
		"INCLUDE",
		"MUST",
		"SET",
		"SHOULD",
		"SLEEP",
		"'$'",
		"FLOAT",
		"INTEGER",
		"STRING",
		"TYPE",
		"WHEN",
		"INTO",
		"'('",
		"AND",
		"OR",
		"')'",
		"FALSE",
		"TRUE",
		"variable",
		"any_value",
		"number",
		"'%'",
		"'*'",
		"'+'",
		"'-'",
		"'/'",
		"CONTAIN",
		"CONTAINS",
		"EQUAL",
		"EQUALS",
		"GE",
		"GT",
		"IS",
		"ISNOT",
		"LE",
		"LT",
		"MATCH",
		"MATCHES",
		"NOTEQUAL",
		"NOTEQUALS",
		"STARTSWITH",
		"expr",
		"BODY",
		"HEADER",
		"true_false",
		"boolean_exp",
		"expr_opr",
		"operator",
		"'}'",
		"http_command_param",
		"any_command",
		"assert_command",
		"command",
		"command_with_condition_opt",
		"CONNECT",
		"debug_command",
		"DELETE",
		"end_command",
		"GET",
		"HEAD",
		"http_command",
		"http_command_params",
		"http_method",
		"include_command",
		"multi_any_value",
		"must_command",
		"OPTIONS",
		"PATCH",
		"POST",
		"PUT",
		"set_command",
		"should_command",
		"sleep_command",
		"TRACE",
		"$default",
		"'&'",
		"'|'",
		"EOF",
		"EOL",
		"error",
		"KEYWORD",
	}

	yyTokenLiteralStrings = map[int]string{}

	yyReductions = map[int]struct{ xsym, components int }{
		0:  {0, 1},
		1:  {57, 0},
		2:  {57, 2},
		3:  {60, 5},
		4:  {60, 3},
		5:  {60, 3},
		6:  {60, 1},
		7:  {59, 1},
		8:  {59, 1},
		9:  {59, 1},
		10: {59, 1},
		11: {59, 1},
		12: {59, 1},
		13: {59, 1},
		14: {59, 1},
		15: {59, 1},
		16: {79, 2},
		17: {70, 2},
		18: {62, 2},
		19: {64, 3},
		20: {64, 2},
		21: {64, 1},
		22: {58, 2},
		23: {72, 2},
		24: {78, 2},
		25: {77, 3},
		26: {77, 3},
		27: {67, 4},
		28: {67, 3},
		29: {68, 1},
		30: {68, 2},
		31: {56, 2},
		32: {56, 2},
		33: {69, 1},
		34: {69, 1},
		35: {69, 1},
		36: {69, 1},
		37: {69, 1},
		38: {69, 1},
		39: {69, 1},
		40: {69, 1},
		41: {69, 1},
		42: {71, 1},
		43: {71, 2},
		44: {26, 1},
		45: {26, 1},
		46: {26, 1},
		47: {26, 1},
		48: {25, 5},
		49: {25, 2},
		50: {25, 1},
		51: {54, 1},
		52: {54, 1},
		53: {54, 1},
		54: {54, 1},
		55: {54, 1},
		56: {54, 1},
		57: {54, 1},
		58: {54, 1},
		59: {54, 1},
		60: {54, 1},
		61: {54, 1},
		62: {54, 1},
		63: {54, 1},
		64: {54, 1},
		65: {54, 1},
		66: {52, 1},
		67: {52, 3},
		68: {52, 3},
		69: {52, 3},
		70: {52, 1},
		71: {51, 1},
		72: {51, 1},
		73: {53, 3},
		74: {53, 3},
		75: {53, 3},
		76: {53, 3},
		77: {53, 3},
		78: {48, 3},
		79: {48, 3},
		80: {48, 3},
		81: {48, 3},
		82: {48, 3},
		83: {48, 3},
		84: {48, 1},
		85: {27, 1},
		86: {27, 1},
	}

	yyXErrors = map[yyXError]string{}

	yyParseTab = [124][]uint16{
		// 0
		{1: 86, 3: 86, 86, 86, 86, 86, 86, 86, 86, 86, 57: 88},
		{1: 87, 3: 104, 102, 103, 108, 101, 105, 107, 106, 100, 58: 95, 90, 89, 62: 93, 64: 94, 67: 92, 70: 98, 72: 96, 77: 91, 97, 99},
		{1: 85, 3: 85, 85, 85, 85, 85, 85, 85, 85, 85},
		{1: 81, 3: 81, 81, 81, 81, 81, 81, 81, 81, 81, 17: 206, 205},
		{1: 80, 3: 80, 80, 80, 80, 80, 80, 80, 80, 80, 17: 80, 80},
		// 5
		{1: 79, 3: 79, 79, 79, 79, 79, 79, 79, 79, 79, 17: 79, 79},
		{1: 78, 3: 78, 78, 78, 78, 78, 78, 78, 78, 78, 17: 78, 78},
		{1: 77, 3: 77, 77, 77, 77, 77, 77, 77, 77, 77, 17: 77, 77},
		{1: 76, 3: 76, 76, 76, 76, 76, 76, 76, 76, 76, 17: 76, 76},
		{1: 75, 3: 75, 75, 75, 75, 75, 75, 75, 75, 75, 17: 75, 75},
		// 10
		{1: 74, 3: 74, 74, 74, 74, 74, 74, 74, 74, 74, 17: 74, 74},
		{1: 73, 3: 73, 73, 73, 73, 73, 73, 73, 73, 73, 17: 73, 73},
		{1: 72, 3: 72, 72, 72, 72, 72, 72, 72, 72, 72, 17: 72, 72},
		{126, 2: 124, 12: 125, 128, 127, 120, 123, 25: 121, 204, 122},
		{126, 2: 124, 12: 125, 128, 127, 120, 123, 25: 121, 203, 122},
		// 15
		{126, 2: 124, 12: 125, 128, 127, 120, 123, 25: 121, 201, 122, 71: 200},
		{126, 66, 124, 66, 66, 66, 66, 66, 66, 66, 66, 66, 125, 128, 127, 120, 123, 197, 66, 145, 23: 149, 148, 121, 146, 122, 48: 191, 51: 144, 198, 147},
		{126, 2: 124, 12: 125, 128, 127, 120, 123, 19: 145, 23: 149, 148, 121, 146, 122, 48: 191, 51: 144, 196, 147},
		{126, 2: 124, 12: 125, 128, 127, 120, 123, 19: 145, 23: 149, 148, 121, 146, 122, 48: 191, 51: 144, 195, 147},
		{126, 2: 124, 12: 125, 128, 127, 120, 123, 19: 145, 23: 149, 148, 121, 146, 122, 48: 191, 51: 144, 194, 147},
		// 20
		{126, 2: 124, 12: 125, 25: 141},
		{61: 115, 63: 114, 65: 110, 111, 69: 109, 73: 116, 118, 112, 113, 80: 117},
		{126, 2: 124, 12: 125, 128, 127, 120, 123, 25: 121, 119, 122},
		{54, 2: 54, 12: 54, 54, 54, 54, 54},
		{53, 2: 53, 12: 53, 53, 53, 53, 53},
		// 25
		{52, 2: 52, 12: 52, 52, 52, 52, 52},
		{51, 2: 51, 12: 51, 51, 51, 51, 51},
		{50, 2: 50, 12: 50, 50, 50, 50, 50},
		{49, 2: 49, 12: 49, 49, 49, 49, 49},
		{48, 2: 48, 12: 48, 48, 48, 48, 48},
		// 30
		{47, 2: 47, 12: 47, 47, 47, 47, 47},
		{46, 2: 46, 12: 46, 46, 46, 46, 46},
		{1: 59, 3: 59, 59, 59, 59, 59, 59, 59, 59, 59, 17: 59, 59, 49: 137, 136, 56: 135, 68: 134},
		{43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 20: 43, 43, 43, 28: 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 49: 43, 43},
		{42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 20: 42, 42, 42, 28: 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 49: 42, 42},
		// 35
		{41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 20: 41, 41, 41, 28: 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 49: 41, 41},
		{40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 20: 40, 40, 40, 28: 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 49: 40, 40},
		{2: 130},
		{129},
		{37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 28: 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 49: 37, 37},
		// 40
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 20: 2, 2, 2, 28: 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 49: 2, 2},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 20: 1, 1, 1, 28: 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 49: 1, 1},
		{38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 28: 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 49: 38, 38},
		{131},
		{55: 132},
		// 45
		{55: 133},
		{39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 28: 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 49: 39, 39},
		{1: 60, 3: 60, 60, 60, 60, 60, 60, 60, 60, 60, 17: 60, 60, 49: 137, 136, 56: 140},
		{1: 58, 3: 58, 58, 58, 58, 58, 58, 58, 58, 58, 17: 58, 58, 49: 58, 58},
		{126, 2: 124, 12: 125, 128, 127, 120, 123, 25: 121, 139, 122},
		// 50
		{126, 2: 124, 12: 125, 128, 127, 120, 123, 25: 121, 138, 122},
		{1: 55, 3: 55, 55, 55, 55, 55, 55, 55, 55, 55, 17: 55, 55, 49: 55, 55},
		{1: 56, 3: 56, 56, 56, 56, 56, 56, 56, 56, 56, 17: 56, 56, 49: 56, 56},
		{1: 57, 3: 57, 57, 57, 57, 57, 57, 57, 57, 57, 17: 57, 57, 49: 57, 57},
		{126, 2: 124, 12: 125, 128, 127, 120, 123, 19: 145, 23: 149, 148, 121, 146, 122, 48: 142, 51: 144, 143, 147},
		// 55
		{1: 62, 3: 62, 62, 62, 62, 62, 62, 62, 62, 62, 17: 62, 62, 28: 177, 175, 173, 174, 176, 159, 158, 151, 150, 155, 154, 163, 164, 157, 156, 161, 162, 153, 152, 160, 54: 185},
		{1: 61, 3: 61, 61, 61, 61, 61, 61, 61, 61, 61, 17: 61, 61, 20: 189, 190},
		{1: 21, 3: 21, 21, 21, 21, 21, 21, 21, 21, 21, 17: 21, 21, 20: 21, 21, 21},
		{126, 2: 124, 12: 125, 128, 127, 120, 123, 19: 145, 23: 149, 148, 121, 146, 122, 48: 184, 51: 144, 183, 147},
		{1: 3, 3: 3, 3, 3, 3, 3, 3, 3, 3, 3, 17: 3, 3, 22: 3, 28: 3, 3, 3, 3, 3, 159, 158, 151, 150, 155, 154, 163, 164, 157, 156, 161, 162, 153, 152, 160, 54: 165},
		// 60
		{1: 17, 3: 17, 17, 17, 17, 17, 17, 17, 17, 17, 17: 17, 17, 20: 17, 17, 17},
		{1: 16, 3: 16, 16, 16, 16, 16, 16, 16, 16, 16, 17: 16, 16, 20: 16, 16, 16},
		{1: 15, 3: 15, 15, 15, 15, 15, 15, 15, 15, 15, 17: 15, 15, 20: 15, 15, 15},
		{36, 2: 36, 12: 36, 36, 36, 36, 36, 19: 36, 23: 36, 36},
		{35, 2: 35, 12: 35, 35, 35, 35, 35, 19: 35, 23: 35, 35},
		// 65
		{34, 2: 34, 12: 34, 34, 34, 34, 34, 19: 34, 23: 34, 34},
		{33, 2: 33, 12: 33, 33, 33, 33, 33, 19: 33, 23: 33, 33},
		{32, 2: 32, 12: 32, 32, 32, 32, 32, 19: 32, 23: 32, 32},
		{31, 2: 31, 12: 31, 31, 31, 31, 31, 19: 31, 23: 31, 31},
		{30, 2: 30, 12: 30, 30, 30, 30, 30, 19: 30, 23: 30, 30},
		// 70
		{29, 2: 29, 12: 29, 29, 29, 29, 29, 19: 29, 23: 29, 29},
		{28, 2: 28, 12: 28, 28, 28, 28, 28, 19: 28, 23: 28, 28},
		{27, 2: 27, 12: 27, 27, 27, 27, 27, 19: 27, 23: 27, 27},
		{26, 2: 26, 12: 26, 26, 26, 26, 26, 19: 26, 23: 26, 26},
		{25, 2: 25, 12: 25, 25, 25, 25, 25, 19: 25, 23: 25, 25},
		// 75
		{24, 2: 24, 12: 24, 24, 24, 24, 24, 19: 24, 23: 24, 24},
		{23, 2: 23, 12: 23, 23, 23, 23, 23, 19: 23, 23: 23, 23},
		{22, 2: 22, 12: 22, 22, 22, 22, 22, 19: 22, 23: 22, 22},
		{126, 2: 124, 12: 125, 128, 127, 120, 123, 19: 169, 23: 149, 148, 121, 166, 122, 48: 168, 51: 167},
		{1: 19, 3: 19, 19, 19, 19, 19, 19, 19, 19, 19, 17: 19, 19, 20: 19, 19, 19, 28: 3, 3, 3, 3, 3},
		// 80
		{1: 18, 3: 18, 18, 18, 18, 18, 18, 18, 18, 18, 17: 18, 18, 20: 18, 18, 18},
		{1: 13, 3: 13, 13, 13, 13, 13, 13, 13, 13, 13, 17: 13, 13, 20: 13, 13, 13, 28: 177, 175, 173, 174, 176},
		{126, 2: 124, 12: 125, 128, 127, 120, 123, 19: 169, 25: 121, 171, 122, 48: 170},
		{22: 172, 28: 177, 175, 173, 174, 176},
		{1: 3, 3: 3, 3, 3, 3, 3, 3, 3, 3, 3, 17: 3, 3, 20: 3, 3, 3, 28: 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
		// 85
		{1: 9, 3: 9, 9, 9, 9, 9, 9, 9, 9, 9, 17: 9, 9, 20: 9, 9, 9, 28: 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		{126, 2: 124, 12: 125, 128, 127, 120, 123, 19: 169, 25: 121, 171, 122, 48: 182},
		{126, 2: 124, 12: 125, 128, 127, 120, 123, 19: 169, 25: 121, 171, 122, 48: 181},
		{126, 2: 124, 12: 125, 128, 127, 120, 123, 19: 169, 25: 121, 171, 122, 48: 180},
		{126, 2: 124, 12: 125, 128, 127, 120, 123, 19: 169, 25: 121, 171, 122, 48: 179},
		// 90
		{126, 2: 124, 12: 125, 128, 127, 120, 123, 19: 169, 25: 121, 171, 122, 48: 178},
		{1: 4, 3: 4, 4, 4, 4, 4, 4, 4, 4, 4, 17: 4, 4, 20: 4, 4, 4, 28: 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
		{1: 5, 3: 5, 5, 5, 5, 5, 5, 5, 5, 5, 17: 5, 5, 20: 5, 5, 5, 28: 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5},
		{1: 6, 3: 6, 6, 6, 6, 6, 6, 6, 6, 6, 17: 6, 6, 20: 6, 6, 6, 28: 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{1: 7, 3: 7, 7, 7, 7, 7, 7, 7, 7, 7, 17: 7, 7, 20: 7, 7, 7, 28: 177, 175, 7, 7, 176, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7},
		// 95
		{1: 8, 3: 8, 8, 8, 8, 8, 8, 8, 8, 8, 17: 8, 8, 20: 8, 8, 8, 28: 177, 175, 8, 8, 176, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8},
		{20: 189, 190, 188},
		{22: 172, 28: 177, 175, 173, 174, 176, 159, 158, 151, 150, 155, 154, 163, 164, 157, 156, 161, 162, 153, 152, 160, 54: 185},
		{126, 2: 124, 12: 125, 128, 127, 120, 123, 19: 169, 25: 121, 187, 122, 48: 186},
		{1: 14, 3: 14, 14, 14, 14, 14, 14, 14, 14, 14, 17: 14, 14, 20: 14, 14, 14, 28: 177, 175, 173, 174, 176},
		// 100
		{1: 12, 3: 12, 12, 12, 12, 12, 12, 12, 12, 12, 17: 12, 12, 20: 12, 12, 12, 28: 3, 3, 3, 3, 3},
		{1: 20, 3: 20, 20, 20, 20, 20, 20, 20, 20, 20, 17: 20, 20, 20: 20, 20, 20},
		{126, 2: 124, 12: 125, 128, 127, 120, 123, 19: 145, 23: 149, 148, 121, 146, 122, 48: 191, 51: 144, 193, 147},
		{126, 2: 124, 12: 125, 128, 127, 120, 123, 19: 145, 23: 149, 148, 121, 146, 122, 48: 191, 51: 144, 192, 147},
		{28: 177, 175, 173, 174, 176, 159, 158, 151, 150, 155, 154, 163, 164, 157, 156, 161, 162, 153, 152, 160, 54: 185},
		// 105
		{1: 10, 3: 10, 10, 10, 10, 10, 10, 10, 10, 10, 17: 10, 10, 20: 189, 190, 10},
		{1: 11, 3: 11, 11, 11, 11, 11, 11, 11, 11, 11, 17: 11, 11, 20: 189, 190, 11},
		{1: 63, 3: 63, 63, 63, 63, 63, 63, 63, 63, 63, 17: 63, 63, 20: 189, 190},
		{1: 64, 3: 64, 64, 64, 64, 64, 64, 64, 64, 64, 17: 64, 64, 20: 189, 190},
		{1: 65, 3: 65, 65, 65, 65, 65, 65, 65, 65, 65, 17: 65, 65, 20: 189, 190},
		// 110
		{126, 2: 124, 12: 125, 128, 127, 120, 123, 19: 145, 23: 149, 148, 121, 146, 122, 48: 191, 51: 144, 199, 147},
		{1: 67, 3: 67, 67, 67, 67, 67, 67, 67, 67, 67, 17: 67, 67, 20: 189, 190},
		{1: 68, 3: 68, 68, 68, 68, 68, 68, 68, 68, 68, 17: 68, 68, 20: 189, 190},
		{126, 69, 124, 69, 69, 69, 69, 69, 69, 69, 69, 69, 125, 128, 127, 120, 123, 69, 69, 25: 121, 202, 122},
		{45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45},
		// 115
		{44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44},
		{1: 70, 3: 70, 70, 70, 70, 70, 70, 70, 70, 70, 17: 70, 70},
		{1: 71, 3: 71, 71, 71, 71, 71, 71, 71, 71, 71, 17: 71, 71},
		{126, 2: 124, 12: 125, 25: 208},
		{126, 2: 124, 12: 125, 128, 127, 120, 123, 19: 145, 23: 149, 148, 121, 146, 122, 48: 191, 51: 144, 207, 147},
		// 120
		{1: 82, 3: 82, 82, 82, 82, 82, 82, 82, 82, 82, 20: 189, 190},
		{1: 83, 3: 83, 83, 83, 83, 83, 83, 83, 83, 83, 17: 209},
		{126, 2: 124, 12: 125, 128, 127, 120, 123, 19: 145, 23: 149, 148, 121, 146, 122, 48: 191, 51: 144, 210, 147},
		{1: 84, 3: 84, 84, 84, 84, 84, 84, 84, 84, 84, 20: 189, 190},
	}
)

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyLexerEx interface {
	yyLexer
	Reduced(rule, state int, lval *yySymType) bool
}

func yySymName(c int) (s string) {
	x, ok := yyXLAT[c]
	if ok {
		return yySymNames[x]
	}

	if c < 0x7f {
		return __yyfmt__.Sprintf("%q", c)
	}

	return __yyfmt__.Sprintf("%d", c)
}

func yylex1(yylex yyLexer, lval *yySymType) (n int) {
	n = yylex.Lex(lval)
	if n <= 0 {
		n = yyEofCode
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("\nlex %s(%#x %d), lval: %+v\n", yySymName(n), n, n, lval)
	}
	return n
}

func yyParse(yylex yyLexer) int {
	const yyError = 86

	yyEx, _ := yylex.(yyLexerEx)
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, 200)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yyerrok := func() {
		if yyDebug >= 2 {
			__yyfmt__.Printf("yyerrok()\n")
		}
		Errflag = 0
	}
	_ = yyerrok
	yystate := 0
	yychar := -1
	var yyxchar int
	var yyshift int
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	if yychar < 0 {
		yylval.yys = yystate
		yychar = yylex1(yylex, &yylval)
		var ok bool
		if yyxchar, ok = yyXLAT[yychar]; !ok {
			yyxchar = len(yySymNames) // > tab width
		}
	}
	if yyDebug >= 4 {
		var a []int
		for _, v := range yyS[:yyp+1] {
			a = append(a, v.yys)
		}
		__yyfmt__.Printf("state stack %v\n", a)
	}
	row := yyParseTab[yystate]
	yyn = 0
	if yyxchar < len(row) {
		if yyn = int(row[yyxchar]); yyn != 0 {
			yyn += yyTabOfs
		}
	}
	switch {
	case yyn > 0: // shift
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		yyshift = yyn
		if yyDebug >= 2 {
			__yyfmt__.Printf("shift, and goto state %d\n", yystate)
		}
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	case yyn < 0: // reduce
	case yystate == 1: // accept
		if yyDebug >= 2 {
			__yyfmt__.Println("accept")
		}
		goto ret0
	}

	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			if yyDebug >= 1 {
				__yyfmt__.Printf("no action for %s in state %d\n", yySymName(yychar), yystate)
			}
			msg, ok := yyXErrors[yyXError{yystate, yyxchar}]
			if !ok {
				msg, ok = yyXErrors[yyXError{yystate, -1}]
			}
			if !ok && yyshift != 0 {
				msg, ok = yyXErrors[yyXError{yyshift, yyxchar}]
			}
			if !ok {
				msg, ok = yyXErrors[yyXError{yyshift, -1}]
			}
			if yychar > 0 {
				ls := yyTokenLiteralStrings[yychar]
				if ls == "" {
					ls = yySymName(yychar)
				}
				if ls != "" {
					switch {
					case msg == "":
						msg = __yyfmt__.Sprintf("unexpected %s", ls)
					default:
						msg = __yyfmt__.Sprintf("unexpected %s, %s", ls, msg)
					}
				}
			}
			if msg == "" {
				msg = "syntax error"
			}
			yylex.Error(msg)
			Nerrs++
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				row := yyParseTab[yyS[yyp].yys]
				if yyError < len(row) {
					yyn = int(row[yyError]) + yyTabOfs
					if yyn > 0 { // hit
						if yyDebug >= 2 {
							__yyfmt__.Printf("error recovery found error shift in state %d\n", yyS[yyp].yys)
						}
						yystate = yyn /* simulate a shift of "error" */
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery failed\n")
			}
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yySymName(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}

			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	r := -yyn
	x0 := yyReductions[r]
	x, n := x0.xsym, x0.components
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= n
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	exState := yystate
	yystate = int(yyParseTab[yyS[yyp].yys][x]) + yyTabOfs
	/* reduction by production r */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce using rule %v (%s), and goto state %d\n", r, yySymNames[x], yystate)
	}

	switch r {
	case 3:
		{
			//run command put result into variable WHEN boolean_exp is true
			if strings.Contains(yyS[yypt-2].variable.name, ".") {
				yylex.Error("nested variables are not supported yet")
			}

			if yyS[yypt-0].boolean {
				GlobalVars[yyS[yypt-2].variable.name] = yyS[yypt-4].cmd.Run()
			}
		}
	case 4:
		{
			//command INTO variable
			if strings.Contains(yyS[yypt-0].variable.name, ".") {
				yylex.Error("nested variables are not supported yet")
			}
			GlobalVars[yyS[yypt-0].variable.name] = yyS[yypt-2].cmd.Run()
		}
	case 5:
		{
			//run the command only if boolean_exp is true
			if yyS[yypt-0].boolean {
				yyS[yypt-2].cmd.Run()
			}
		}
	case 6:
		{
			//just run the command
			yyS[yypt-0].cmd.Run()
			//run command without condition

		}
	case 16:
		{
			yyVAL.cmd = &SleepCommand{
				Millisecond: intVal(yyS[yypt-0].val),
			}
		}
	case 17:
		{
			yyVAL.cmd = &IncludeCommand{
				File: yyS[yypt-0].val.(string),
			}
		}
	case 18:
		{
			yyVAL.cmd = &DebugCommand{
				Values: yyS[yypt-0].vals,
			}
		}
	case 19:
		{
			if yyS[yypt-0].boolean {
				return -1
			}

			yyVAL.cmd = &EndCommand{}
		}
	case 20:
		{
			if yyS[yypt-0].boolean {
				return -1
			}

			yyVAL.cmd = &EndCommand{}
		}
	case 21:
		{
			return -1
		}
	case 22:
		{
			if !yyS[yypt-0].boolean {
				State.Assertion.Failed++
			} else {
				State.Assertion.Succeeded++
			}
			yyVAL.cmd = &AssertCommand{}
		}
	case 23:
		{
			if !yyS[yypt-0].boolean {
				State.Must.Failed++
			} else {
				State.Must.Succeeded++
			}

			yyVAL.cmd = &MustCommand{}
		}
	case 24:
		{
			if !yyS[yypt-0].boolean {
				State.Should.Failed++
			} else {
				State.Should.Succeeded++
			}
			yyVAL.cmd = &ShouldCommand{}
		}
	case 25:
		{
			//GlobalVars[$2.name] = $3
			yyVAL.cmd = &SetCommand{
				Name:  yyS[yypt-1].variable.name,
				Value: yyS[yypt-0].val,
			}
		}
	case 26:
		{
			//GlobalVars[$2.name] = $3
			yyVAL.cmd = &SetCommand{
				Name:  yyS[yypt-1].variable.name,
				Value: yyS[yypt-0].boolean,
			}
		}
	case 27:
		{
			//call http with header here.
			yyVAL.cmd = &HttpCommand{
				Method:        yyS[yypt-2].val.(string),
				CommandParams: yyS[yypt-0].http_command_params,
				Url:           yyS[yypt-1].val.(string),
			}
		}
	case 28:
		{
			//simple http command
			yyVAL.cmd = &HttpCommand{
				Method: yyS[yypt-1].val.(string),
				Url:    yyS[yypt-0].val.(string),
			}
		}
	case 29:
		{
			if yyVAL.http_command_params == nil {
				yyVAL.http_command_params = make([]HttpCommandParam, 0)
			}
			yyVAL.http_command_params = append(yyVAL.http_command_params, yyS[yypt-0].http_command_param)
		}
	case 30:
		{
			if yyVAL.http_command_params == nil {
				yyVAL.http_command_params = make([]HttpCommandParam, 0)
			}

			yyVAL.http_command_params = append(yyVAL.http_command_params, yyS[yypt-0].http_command_param)
		}
	case 31:
		{
			//addin header
			yyVAL.http_command_param = HttpCommandParam{
				ParamName:  yyS[yypt-1].val.(string),
				ParamValue: yyS[yypt-0].val.(string),
			}
		}
	case 32:
		{
			//adding query param
			yyVAL.http_command_param = HttpCommandParam{
				ParamName:  yyS[yypt-1].val.(string),
				ParamValue: yyS[yypt-0].val.(string),
			}
		}
	case 42:
		{
			//getting a single value from multi_value exp
			yyVAL.vals = append(yyVAL.vals, yyS[yypt-0].val)
		}
	case 43:
		{
			//multi value
			yyVAL.vals = append(yyVAL.vals, yyS[yypt-0].val)
		}
	case 44:
		{
			//string_or_var : STRING
			if isTemplate(yyS[yypt-0].val.(string)) {
				yyVAL.val, _ = executeTemplate(yyS[yypt-0].val.(string), GlobalVars)
			} else {

			}
		}
	case 45:
		{
			//any_value : variable
			switch yyS[yypt-0].variable.value.(type) {
			case string:
				if isTemplate(yyS[yypt-0].variable.value.(string)) {
					yyVAL.val, _ = executeTemplate(yyS[yypt-0].variable.value.(string), GlobalVars)
				} else {
					yyVAL.val = yyS[yypt-0].variable.value
				}
			default:
				yyVAL.val = yyS[yypt-0].variable.value
			}

		}
	case 48:
		{
			//getting variable
			yyVAL.variable.name = yyS[yypt-2].val.(string)
			yyVAL.variable.value = query(yyS[yypt-2].val.(string), GlobalVars)
		}
	case 49:
		{
			yyVAL.variable.name = yyS[yypt-0].val.(string)
			yyVAL.variable.value = query(yyS[yypt-0].val.(string), GlobalVars)
		}
	case 50:
		{
			yyVAL.variable.name = yyS[yypt-0].val.(string)
			yyVAL.variable.value = query(yyS[yypt-0].val.(string), GlobalVars)
		}
	case 67:
		{
			//boolean_ex: '(' boolean_exp ')'
			yyVAL.boolean = yyS[yypt-1].boolean
		}
	case 68:
		{
			//boolean_ex: any_value operator any_value, oh conflicts start here :(
			operator_result := runop(yyS[yypt-2].val, yyS[yypt-1].val, yyS[yypt-0].val)
			yyVAL.boolean = operator_result
		}
	case 69:
		{
			//boolean_ex: any_value operator any_value, oh conflicts start here :(
			operator_result := runop(yyS[yypt-2].val, yyS[yypt-1].val, yyS[yypt-0].boolean)
			yyVAL.boolean = operator_result
		}
	case 71:
		{
			yyVAL.boolean = true
		}
	case 72:
		{
			yyVAL.boolean = false
		}
	case 73:
		{
			operator_result := runop(yyS[yypt-2].val, yyS[yypt-1].val, yyS[yypt-0].val)
			yyVAL.boolean = operator_result
		}
	case 74:
		{
			operator_result := runop(yyS[yypt-2].val, yyS[yypt-1].val, yyS[yypt-0].val)
			yyVAL.boolean = operator_result
		}
	case 75:
		{
			operator_result := runop(yyS[yypt-2].val, yyS[yypt-1].val, yyS[yypt-0].val)
			yyVAL.boolean = operator_result
		}
	case 76:
		{
			//boolean_ex: boolean_exp AND boolean_exp
			yyVAL.boolean = yyS[yypt-2].boolean && yyS[yypt-0].boolean
		}
	case 77:
		{
			//boolean_ex: boolean_exp OR boolean_exp
			yyVAL.boolean = yyS[yypt-2].boolean || yyS[yypt-0].boolean
		}
	case 78:
		{
			yyVAL.val = yyS[yypt-1].val
		}
	case 79:
		{
			yyVAL.val, _ = add(yyS[yypt-2].val, yyS[yypt-0].val)
		}
	case 80:
		{
			yyVAL.val, _ = subtract(yyS[yypt-0].val, yyS[yypt-2].val)
		}
	case 81:
		{
			yyVAL.val, _ = multiply(yyS[yypt-0].val, yyS[yypt-2].val)
		}
	case 82:
		{
			yyVAL.val, _ = divide(yyS[yypt-0].val, yyS[yypt-2].val)
		}
	case 83:
		{
			yyVAL.val, _ = mod(yyS[yypt-0].val, yyS[yypt-2].val)
		}
	case 85:
		{
			//number: INTEGER
			yyVAL.val = yyS[yypt-0].val
		}
	case 86:
		{
			//number: FLOAT
			yyVAL.val = yyS[yypt-0].val
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}

type lex struct {
	tokens chan Token
}

func (l *lex) All() []Token {
	tokens := make([]Token, 0)
	for {
		v := <-l.tokens
		if v.Type == EOF || v.Type == -1 {
			break
		}

		tokens = append(tokens, v)
	}

	return tokens
}

func (l *lex) Lex(lval *yySymType) int {
	v := <-l.tokens
	if v.Type == EOF || v.Type == -1 {
		return 0
	}
	lval.val = v.Val
	return v.Type
}

func (l *lex) Error(e string) {
	log.Fatal(e)
}

//TODO: use channels here.
//Parse parses a given string and returns a lex
func Parse(text string) *lex {

	l := &lex{
		tokens: make(chan Token),
	}

	if Verbose {
		yyDebug = 3
	}

	SetStateErrors()

	go func() {
		s := NewScanner(strings.NewReader(strings.TrimSpace(text)))
		for {
			l.tokens <- s.Scan()
		}
	}()

	return l
}

//Resets the state to start over
func Reset() {
	GlobalVars = map[string]interface{}{}
	State = NewStats()
}

func Run(l *lex) {
	yyParse(l)
}
