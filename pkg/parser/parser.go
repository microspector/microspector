// Code generated by goyacc - DO NOT EDIT.

package parser

import __yyfmt__ "fmt"

import (
	"strings"
	"sync"
)

type yySymType struct {
	yys                 int
	expression          Expression
	expressions         ExprArray
	val                 interface{}
	vals                []interface{}
	str                 ExprString
	integer             ExprInteger
	boolean             ExprBool
	bytes               []byte
	cmd                 Command
	variable            ExprVariable
	http_command_params []HttpCommandParam
	http_command_param  HttpCommandParam
}

type yyXError struct {
	state, xsym int
}

const (
	yyDefault  = 57406
	yyEofCode  = 57344
	AND        = 57395
	ASSERT     = 57361
	ASYNC      = 57365
	BODY       = 57377
	CMD        = 57364
	CONNECT    = 57372
	CONTAIN    = 57391
	CONTAINS   = 57390
	DEBUG      = 57359
	DELETE     = 57371
	ECHO       = 57366
	END        = 57360
	EOF        = 57347
	EOL        = 57346
	EQUAL      = 57383
	EQUALS     = 57382
	FALSE      = 57352
	FLOAT      = 57350
	FOLLOW     = 57378
	GE         = 57387
	GET        = 57367
	GT         = 57386
	HEAD       = 57368
	HEADER     = 57376
	HTTP       = 57356
	IDENTIFIER = 57404
	IN         = 57402
	INCLUDE    = 57362
	INSECURE   = 57381
	INTEGER    = 57349
	INTO       = 57403
	IS         = 57399
	ISNOT      = 57400
	KEYWORD    = 57354
	LE         = 57389
	LT         = 57388
	MATCH      = 57398
	MATCHES    = 57397
	MUST       = 57357
	NOFOLLOW   = 57379
	NOT        = 57401
	NOTEQUAL   = 57385
	NOTEQUALS  = 57384
	NULL       = 57353
	OPTIONS    = 57373
	OR         = 57396
	PATCH      = 57375
	POST       = 57369
	PUT        = 57370
	SECURE     = 57380
	SET        = 57355
	SHOULD     = 57358
	SLEEP      = 57363
	STARTSWITH = 57392
	STARTWITH  = 57393
	STRING     = 57348
	TRACE      = 57374
	TRUE       = 57351
	TYPE       = 57405
	WHEN       = 57394
	yyErrCode  = 57345

	yyMaxDepth = 200
	yyTabOfs   = -90
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
		37:    0,   // '%' (84x)
		57404: 1,   // IDENTIFIER (59x)
		123:   2,   // '{' (58x)
		36:    3,   // '$' (57x)
		57361: 4,   // ASSERT (57x)
		57364: 5,   // CMD (57x)
		57359: 6,   // DEBUG (57x)
		57366: 7,   // ECHO (57x)
		57360: 8,   // END (57x)
		57356: 9,   // HTTP (57x)
		57362: 10,  // INCLUDE (57x)
		57357: 11,  // MUST (57x)
		57355: 12,  // SET (57x)
		57358: 13,  // SHOULD (57x)
		57363: 14,  // SLEEP (57x)
		57344: 15,  // $end (56x)
		57365: 16,  // ASYNC (56x)
		57349: 17,  // INTEGER (55x)
		40:    18,  // '(' (54x)
		91:    19,  // '[' (54x)
		57352: 20,  // FALSE (54x)
		57350: 21,  // FLOAT (54x)
		57348: 22,  // STRING (54x)
		57351: 23,  // TRUE (54x)
		57394: 24,  // WHEN (53x)
		57403: 25,  // INTO (49x)
		42:    26,  // '*' (33x)
		43:    27,  // '+' (33x)
		45:    28,  // '-' (33x)
		47:    29,  // '/' (33x)
		57395: 30,  // AND (33x)
		57391: 31,  // CONTAIN (33x)
		57390: 32,  // CONTAINS (33x)
		57383: 33,  // EQUAL (33x)
		57382: 34,  // EQUALS (33x)
		57387: 35,  // GE (33x)
		57386: 36,  // GT (33x)
		57402: 37,  // IN (33x)
		57399: 38,  // IS (33x)
		57400: 39,  // ISNOT (33x)
		57389: 40,  // LE (33x)
		57388: 41,  // LT (33x)
		57398: 42,  // MATCH (33x)
		57397: 43,  // MATCHES (33x)
		57401: 44,  // NOT (33x)
		57385: 45,  // NOTEQUAL (33x)
		57384: 46,  // NOTEQUALS (33x)
		57396: 47,  // OR (33x)
		57392: 48,  // STARTSWITH (33x)
		57393: 49,  // STARTWITH (33x)
		57377: 50,  // BODY (25x)
		57378: 51,  // FOLLOW (25x)
		57376: 52,  // HEADER (25x)
		57381: 53,  // INSECURE (25x)
		57379: 54,  // NOFOLLOW (25x)
		57380: 55,  // SECURE (25x)
		57429: 56,  // variable (22x)
		93:    57,  // ']' (19x)
		57407: 58,  // array (19x)
		57416: 59,  // expr (19x)
		57424: 60,  // operator (19x)
		44:    61,  // ',' (18x)
		41:    62,  // ')' (16x)
		125:   63,  // '}' (2x)
		57408: 64,  // assert_command (2x)
		57409: 65,  // cmd_command (2x)
		57411: 66,  // command (2x)
		57412: 67,  // command_cond (2x)
		57413: 68,  // debug_command (2x)
		57414: 69,  // echo_command (2x)
		57415: 70,  // end_command (2x)
		57417: 71,  // http_command (2x)
		57418: 72,  // http_command_param (2x)
		57421: 73,  // include_command (2x)
		57423: 74,  // must_command (2x)
		57426: 75,  // set_command (2x)
		57427: 76,  // should_command (2x)
		57428: 77,  // sleep_command (2x)
		57410: 78,  // comma_separated_expressions (1x)
		57367: 79,  // GET (1x)
		57368: 80,  // HEAD (1x)
		57419: 81,  // http_command_params (1x)
		57420: 82,  // http_method (1x)
		57422: 83,  // microspector (1x)
		57373: 84,  // OPTIONS (1x)
		57375: 85,  // PATCH (1x)
		57369: 86,  // POST (1x)
		57370: 87,  // PUT (1x)
		57425: 88,  // run_comm (1x)
		57406: 89,  // $default (0x)
		38:    90,  // '&' (0x)
		124:   91,  // '|' (0x)
		57372: 92,  // CONNECT (0x)
		57371: 93,  // DELETE (0x)
		57347: 94,  // EOF (0x)
		57346: 95,  // EOL (0x)
		57345: 96,  // error (0x)
		57354: 97,  // KEYWORD (0x)
		57353: 98,  // NULL (0x)
		57374: 99,  // TRACE (0x)
		57405: 100, // TYPE (0x)
	}

	yySymNames = []string{
		"'%'",
		"IDENTIFIER",
		"'{'",
		"'$'",
		"ASSERT",
		"CMD",
		"DEBUG",
		"ECHO",
		"END",
		"HTTP",
		"INCLUDE",
		"MUST",
		"SET",
		"SHOULD",
		"SLEEP",
		"$end",
		"ASYNC",
		"INTEGER",
		"'('",
		"'['",
		"FALSE",
		"FLOAT",
		"STRING",
		"TRUE",
		"WHEN",
		"INTO",
		"'*'",
		"'+'",
		"'-'",
		"'/'",
		"AND",
		"CONTAIN",
		"CONTAINS",
		"EQUAL",
		"EQUALS",
		"GE",
		"GT",
		"IN",
		"IS",
		"ISNOT",
		"LE",
		"LT",
		"MATCH",
		"MATCHES",
		"NOT",
		"NOTEQUAL",
		"NOTEQUALS",
		"OR",
		"STARTSWITH",
		"STARTWITH",
		"BODY",
		"FOLLOW",
		"HEADER",
		"INSECURE",
		"NOFOLLOW",
		"SECURE",
		"variable",
		"']'",
		"array",
		"expr",
		"operator",
		"','",
		"')'",
		"'}'",
		"assert_command",
		"cmd_command",
		"command",
		"command_cond",
		"debug_command",
		"echo_command",
		"end_command",
		"http_command",
		"http_command_param",
		"include_command",
		"must_command",
		"set_command",
		"should_command",
		"sleep_command",
		"comma_separated_expressions",
		"GET",
		"HEAD",
		"http_command_params",
		"http_method",
		"microspector",
		"OPTIONS",
		"PATCH",
		"POST",
		"PUT",
		"run_comm",
		"$default",
		"'&'",
		"'|'",
		"CONNECT",
		"DELETE",
		"EOF",
		"EOL",
		"error",
		"KEYWORD",
		"NULL",
		"TRACE",
		"TYPE",
	}

	yyTokenLiteralStrings = map[int]string{}

	yyReductions = map[int]struct{ xsym, components int }{
		0:  {0, 1},
		1:  {83, 0},
		2:  {83, 2},
		3:  {88, 1},
		4:  {88, 2},
		5:  {67, 3},
		6:  {67, 5},
		7:  {67, 3},
		8:  {67, 1},
		9:  {66, 1},
		10: {66, 1},
		11: {66, 1},
		12: {66, 1},
		13: {66, 1},
		14: {66, 1},
		15: {66, 1},
		16: {66, 1},
		17: {66, 1},
		18: {66, 1},
		19: {66, 1},
		20: {75, 3},
		21: {71, 3},
		22: {71, 4},
		23: {81, 1},
		24: {81, 2},
		25: {72, 2},
		26: {72, 2},
		27: {72, 1},
		28: {72, 1},
		29: {72, 1},
		30: {72, 1},
		31: {68, 2},
		32: {70, 2},
		33: {70, 3},
		34: {64, 2},
		35: {74, 2},
		36: {76, 2},
		37: {73, 2},
		38: {77, 2},
		39: {65, 2},
		40: {69, 2},
		41: {82, 1},
		42: {82, 1},
		43: {82, 1},
		44: {82, 1},
		45: {82, 1},
		46: {82, 1},
		47: {58, 2},
		48: {58, 3},
		49: {78, 1},
		50: {78, 3},
		51: {59, 3},
		52: {59, 1},
		53: {59, 1},
		54: {59, 1},
		55: {59, 2},
		56: {59, 1},
		57: {59, 1},
		58: {59, 1},
		59: {59, 3},
		60: {59, 1},
		61: {56, 5},
		62: {56, 2},
		63: {56, 1},
		64: {60, 1},
		65: {60, 1},
		66: {60, 1},
		67: {60, 1},
		68: {60, 1},
		69: {60, 1},
		70: {60, 1},
		71: {60, 1},
		72: {60, 1},
		73: {60, 1},
		74: {60, 1},
		75: {60, 1},
		76: {60, 1},
		77: {60, 1},
		78: {60, 1},
		79: {60, 1},
		80: {60, 1},
		81: {60, 1},
		82: {60, 1},
		83: {60, 1},
		84: {60, 1},
		85: {60, 1},
		86: {60, 1},
		87: {60, 1},
		88: {60, 1},
		89: {60, 1},
	}

	yyXErrors = map[yyXError]string{}

	yyParseTab = [122][]uint16{
		// 0
		{4: 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 83: 91},
		{4: 111, 116, 109, 117, 110, 108, 114, 112, 107, 113, 115, 90, 94, 64: 100, 105, 95, 93, 98, 106, 99, 97, 73: 103, 101, 96, 102, 104, 88: 92},
		{4: 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88},
		{4: 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87},
		{4: 111, 116, 109, 117, 110, 108, 114, 112, 107, 113, 115, 64: 100, 105, 95, 211, 98, 106, 99, 97, 73: 103, 101, 96, 102, 104},
		// 5
		{4: 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 24: 205, 206},
		{4: 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 24: 81, 81},
		{4: 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 24: 80, 80},
		{4: 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 24: 79, 79},
		{4: 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 24: 78, 78},
		// 10
		{4: 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 24: 77, 77},
		{4: 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 24: 76, 76},
		{4: 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 24: 75, 75},
		{4: 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 24: 74, 74},
		{4: 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 24: 73, 73},
		// 15
		{4: 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 24: 72, 72},
		{4: 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 24: 71, 71},
		{1: 131, 129, 130, 56: 203},
		{79: 185, 187, 82: 184, 84: 188, 190, 186, 189},
		{124, 131, 129, 130, 17: 123, 120, 119, 127, 122, 121, 126, 56: 125, 58: 128, 183},
		// 20
		{124, 131, 129, 130, 17: 123, 120, 119, 127, 122, 121, 126, 181, 56: 125, 58: 128, 180},
		{124, 131, 129, 130, 17: 123, 120, 119, 127, 122, 121, 126, 56: 125, 58: 128, 179},
		{124, 131, 129, 130, 17: 123, 120, 119, 127, 122, 121, 126, 56: 125, 58: 128, 178},
		{124, 131, 129, 130, 17: 123, 120, 119, 127, 122, 121, 126, 56: 125, 58: 128, 177},
		{124, 131, 129, 130, 17: 123, 120, 119, 127, 122, 121, 126, 56: 125, 58: 128, 176},
		// 25
		{124, 131, 129, 130, 17: 123, 120, 119, 127, 122, 121, 126, 56: 125, 58: 128, 175},
		{124, 131, 129, 130, 17: 123, 120, 119, 127, 122, 121, 126, 56: 125, 58: 128, 174},
		{124, 131, 129, 130, 17: 123, 120, 119, 127, 122, 121, 126, 56: 125, 58: 128, 118},
		{166, 4: 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 24: 153, 50, 162, 164, 165, 163, 154, 150, 149, 142, 141, 146, 145, 161, 158, 159, 148, 147, 157, 156, 160, 144, 143, 155, 151, 152, 60: 140},
		{124, 131, 129, 130, 17: 123, 120, 119, 127, 122, 121, 126, 56: 125, 168, 128, 170, 78: 169},
		// 30
		{124, 131, 129, 130, 17: 123, 120, 119, 127, 122, 121, 126, 56: 125, 58: 128, 138},
		{38, 4: 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 24: 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 57: 38, 61: 38, 38},
		{37, 4: 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 24: 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 57: 37, 61: 37, 37},
		{36, 4: 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 24: 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 57: 36, 61: 36, 36},
		{17: 137},
		// 35
		{34, 4: 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 24: 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 57: 34, 61: 34, 34},
		{33, 4: 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 24: 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 57: 33, 61: 33, 33},
		{32, 4: 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 24: 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 57: 32, 61: 32, 32},
		{30, 4: 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 24: 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 57: 30, 61: 30, 30},
		{2: 133},
		// 40
		{1: 132},
		{27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 57: 27, 61: 27, 27},
		{28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 57: 28, 61: 28, 28},
		{1: 134},
		{63: 135},
		// 45
		{63: 136},
		{29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 57: 29, 61: 29, 29},
		{35, 4: 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 24: 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 57: 35, 61: 35, 35},
		{166, 24: 153, 26: 162, 164, 165, 163, 154, 150, 149, 142, 141, 146, 145, 161, 158, 159, 148, 147, 157, 156, 160, 144, 143, 155, 151, 152, 60: 140, 62: 139},
		{39, 4: 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 24: 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 57: 39, 61: 39, 39},
		// 50
		{124, 131, 129, 130, 17: 123, 120, 119, 127, 122, 121, 126, 56: 125, 58: 128, 167},
		{26, 26, 26, 26, 17: 26, 26, 26, 26, 26, 26, 26},
		{25, 25, 25, 25, 17: 25, 25, 25, 25, 25, 25, 25},
		{24, 24, 24, 24, 17: 24, 24, 24, 24, 24, 24, 24},
		{23, 23, 23, 23, 17: 23, 23, 23, 23, 23, 23, 23},
		// 55
		{22, 22, 22, 22, 17: 22, 22, 22, 22, 22, 22, 22},
		{21, 21, 21, 21, 17: 21, 21, 21, 21, 21, 21, 21},
		{20, 20, 20, 20, 17: 20, 20, 20, 20, 20, 20, 20},
		{19, 19, 19, 19, 17: 19, 19, 19, 19, 19, 19, 19},
		{18, 18, 18, 18, 17: 18, 18, 18, 18, 18, 18, 18},
		// 60
		{17, 17, 17, 17, 17: 17, 17, 17, 17, 17, 17, 17},
		{16, 16, 16, 16, 17: 16, 16, 16, 16, 16, 16, 16},
		{15, 15, 15, 15, 17: 15, 15, 15, 15, 15, 15, 15},
		{14, 14, 14, 14, 17: 14, 14, 14, 14, 14, 14, 14},
		{13, 13, 13, 13, 17: 13, 13, 13, 13, 13, 13, 13},
		// 65
		{12, 12, 12, 12, 17: 12, 12, 12, 12, 12, 12, 12},
		{11, 11, 11, 11, 17: 11, 11, 11, 11, 11, 11, 11},
		{10, 10, 10, 10, 17: 10, 10, 10, 10, 10, 10, 10},
		{9, 9, 9, 9, 17: 9, 9, 9, 9, 9, 9, 9},
		{8, 8, 8, 8, 17: 8, 8, 8, 8, 8, 8, 8},
		// 70
		{7, 7, 7, 7, 17: 7, 7, 7, 7, 7, 7, 7},
		{6, 6, 6, 6, 17: 6, 6, 6, 6, 6, 6, 6},
		{5, 5, 5, 5, 17: 5, 5, 5, 5, 5, 5, 5},
		{4, 4, 4, 4, 17: 4, 4, 4, 4, 4, 4, 4},
		{3, 3, 3, 3, 17: 3, 3, 3, 3, 3, 3, 3},
		// 75
		{2, 2, 2, 2, 17: 2, 2, 2, 2, 2, 2, 2},
		{1, 1, 1, 1, 17: 1, 1, 1, 1, 1, 1, 1},
		{166, 4: 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 24: 153, 31, 162, 164, 165, 163, 154, 150, 149, 142, 141, 146, 145, 161, 158, 159, 148, 147, 157, 156, 160, 144, 143, 155, 151, 152, 31, 31, 31, 31, 31, 31, 57: 31, 60: 140, 31, 31},
		{43, 4: 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 24: 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 57: 43, 61: 43, 43},
		{57: 171, 61: 172},
		// 80
		{166, 24: 153, 26: 162, 164, 165, 163, 154, 150, 149, 142, 141, 146, 145, 161, 158, 159, 148, 147, 157, 156, 160, 144, 143, 155, 151, 152, 57: 41, 60: 140, 41},
		{42, 4: 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 24: 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 57: 42, 61: 42, 42},
		{124, 131, 129, 130, 17: 123, 120, 119, 127, 122, 121, 126, 56: 125, 58: 128, 173},
		{166, 24: 153, 26: 162, 164, 165, 163, 154, 150, 149, 142, 141, 146, 145, 161, 158, 159, 148, 147, 157, 156, 160, 144, 143, 155, 151, 152, 57: 40, 60: 140, 40},
		{166, 4: 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 24: 153, 51, 162, 164, 165, 163, 154, 150, 149, 142, 141, 146, 145, 161, 158, 159, 148, 147, 157, 156, 160, 144, 143, 155, 151, 152, 60: 140},
		// 85
		{166, 4: 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 24: 153, 52, 162, 164, 165, 163, 154, 150, 149, 142, 141, 146, 145, 161, 158, 159, 148, 147, 157, 156, 160, 144, 143, 155, 151, 152, 60: 140},
		{166, 4: 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 24: 153, 53, 162, 164, 165, 163, 154, 150, 149, 142, 141, 146, 145, 161, 158, 159, 148, 147, 157, 156, 160, 144, 143, 155, 151, 152, 60: 140},
		{166, 4: 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 24: 153, 54, 162, 164, 165, 163, 154, 150, 149, 142, 141, 146, 145, 161, 158, 159, 148, 147, 157, 156, 160, 144, 143, 155, 151, 152, 60: 140},
		{166, 4: 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 24: 153, 55, 162, 164, 165, 163, 154, 150, 149, 142, 141, 146, 145, 161, 158, 159, 148, 147, 157, 156, 160, 144, 143, 155, 151, 152, 60: 140},
		{166, 4: 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 24: 153, 56, 162, 164, 165, 163, 154, 150, 149, 142, 141, 146, 145, 161, 158, 159, 148, 147, 157, 156, 160, 144, 143, 155, 151, 152, 60: 140},
		// 90
		{166, 4: 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 24: 153, 58, 162, 164, 165, 163, 154, 150, 149, 142, 141, 146, 145, 161, 158, 159, 148, 147, 157, 156, 160, 144, 143, 155, 151, 152, 60: 140},
		{124, 131, 129, 130, 17: 123, 120, 119, 127, 122, 121, 126, 56: 125, 58: 128, 182},
		{166, 4: 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 24: 153, 57, 162, 164, 165, 163, 154, 150, 149, 142, 141, 146, 145, 161, 158, 159, 148, 147, 157, 156, 160, 144, 143, 155, 151, 152, 60: 140},
		{166, 4: 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 24: 153, 59, 162, 164, 165, 163, 154, 150, 149, 142, 141, 146, 145, 161, 158, 159, 148, 147, 157, 156, 160, 144, 143, 155, 151, 152, 60: 140},
		{124, 131, 129, 130, 17: 123, 120, 119, 127, 122, 121, 126, 56: 125, 58: 128, 191},
		// 95
		{49, 49, 49, 49, 17: 49, 49, 49, 49, 49, 49, 49},
		{48, 48, 48, 48, 17: 48, 48, 48, 48, 48, 48, 48},
		{47, 47, 47, 47, 17: 47, 47, 47, 47, 47, 47, 47},
		{46, 46, 46, 46, 17: 46, 46, 46, 46, 46, 46, 46},
		{45, 45, 45, 45, 17: 45, 45, 45, 45, 45, 45, 45},
		// 100
		{44, 44, 44, 44, 17: 44, 44, 44, 44, 44, 44, 44},
		{166, 4: 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 24: 153, 69, 162, 164, 165, 163, 154, 150, 149, 142, 141, 146, 145, 161, 158, 159, 148, 147, 157, 156, 160, 144, 143, 155, 151, 152, 195, 196, 194, 198, 197, 199, 60: 140, 72: 193, 81: 192},
		{4: 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 24: 68, 68, 50: 195, 196, 194, 198, 197, 199, 72: 202},
		{4: 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 24: 67, 67, 50: 67, 67, 67, 67, 67, 67},
		{124, 131, 129, 130, 17: 123, 120, 119, 127, 122, 121, 126, 56: 125, 58: 128, 201},
		// 105
		{124, 131, 129, 130, 17: 123, 120, 119, 127, 122, 121, 126, 56: 125, 58: 128, 200},
		{4: 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 24: 63, 63, 50: 63, 63, 63, 63, 63, 63},
		{4: 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 24: 62, 62, 50: 62, 62, 62, 62, 62, 62},
		{4: 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 24: 61, 61, 50: 61, 61, 61, 61, 61, 61},
		{4: 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 24: 60, 60, 50: 60, 60, 60, 60, 60, 60},
		// 110
		{166, 4: 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 24: 153, 64, 162, 164, 165, 163, 154, 150, 149, 142, 141, 146, 145, 161, 158, 159, 148, 147, 157, 156, 160, 144, 143, 155, 151, 152, 64, 64, 64, 64, 64, 64, 60: 140},
		{166, 4: 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 24: 153, 65, 162, 164, 165, 163, 154, 150, 149, 142, 141, 146, 145, 161, 158, 159, 148, 147, 157, 156, 160, 144, 143, 155, 151, 152, 65, 65, 65, 65, 65, 65, 60: 140},
		{4: 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 24: 66, 66, 50: 66, 66, 66, 66, 66, 66},
		{124, 131, 129, 130, 17: 123, 120, 119, 127, 122, 121, 126, 56: 125, 58: 128, 204},
		{166, 4: 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 24: 153, 70, 162, 164, 165, 163, 154, 150, 149, 142, 141, 146, 145, 161, 158, 159, 148, 147, 157, 156, 160, 144, 143, 155, 151, 152, 60: 140},
		// 115
		{124, 131, 129, 130, 17: 123, 120, 119, 127, 122, 121, 126, 56: 125, 58: 128, 208},
		{1: 131, 129, 130, 56: 207},
		{4: 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83},
		{166, 4: 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 24: 153, 209, 162, 164, 165, 163, 154, 150, 149, 142, 141, 146, 145, 161, 158, 159, 148, 147, 157, 156, 160, 144, 143, 155, 151, 152, 60: 140},
		{1: 131, 129, 130, 56: 210},
		// 120
		{4: 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84},
		{4: 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86},
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
	const yyError = 96

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
			yylex.(*Lexer).wg.Add(1)
			yyS[yypt-0].cmd.Run(yylex.(*Lexer))
		}
	case 4:
		{
			yylex.(*Lexer).wg.Add(1)
			go yyS[yypt-0].cmd.Run(yylex.(*Lexer))
		}
	case 5:
		{
			yyVAL.cmd = yyS[yypt-2].cmd
			yyVAL.cmd.SetWhen(yyS[yypt-0].expression)
		}
	case 6:
		{

			yyS[yypt-4].cmd.SetWhen(yyS[yypt-2].expression)
			//TODO: check if it compatible with SetInto
			yyS[yypt-4].cmd.(*HttpCommand).SetInto(yyS[yypt-0].variable.Name)
			yyVAL.cmd = yyS[yypt-4].cmd
		}
	case 7:
		{
			//TODO: check if it compatible with SetInto
			yyS[yypt-2].cmd.(*HttpCommand).SetInto(yyS[yypt-0].variable.Name)
			yyVAL.cmd = yyS[yypt-2].cmd
		}
	case 8:
		{
			yyVAL.cmd = yyS[yypt-0].cmd
		}
	case 20:
		{
			yyVAL.cmd = &SetCommand{
				Name:  yyS[yypt-1].variable.Name,
				Value: yyS[yypt-0].expression,
			}
		}
	case 21:
		{
			yyVAL.cmd = &HttpCommand{
				Method: yyS[yypt-1].val.(string),
				Url:    yyS[yypt-0].expression,
			}
		}
	case 22:
		{
			yyVAL.cmd = &HttpCommand{
				Method:        yyS[yypt-2].val.(string),
				Url:           yyS[yypt-1].expression,
				CommandParams: yyS[yypt-0].http_command_params,
			}
		}
	case 23:
		{
			if yyVAL.http_command_params == nil {
				yyVAL.http_command_params = make([]HttpCommandParam, 0)
			}
			yyVAL.http_command_params = append(yyVAL.http_command_params, yyS[yypt-0].http_command_param)
		}
	case 24:
		{
			if yyVAL.http_command_params == nil {
				yyVAL.http_command_params = make([]HttpCommandParam, 0)
			}

			yyVAL.http_command_params = append(yyVAL.http_command_params, yyS[yypt-0].http_command_param)
		}
	case 25:
		{
			//addin header
			yyVAL.http_command_param = HttpCommandParam{
				ParamName:  yyS[yypt-1].val.(string),
				ParamValue: yyS[yypt-0].expression,
			}
		}
	case 26:
		{
			yyVAL.http_command_param = HttpCommandParam{
				ParamName:  yyS[yypt-1].val.(string),
				ParamValue: yyS[yypt-0].expression,
			}
		}
	case 27:
		{
			yyVAL.http_command_param = HttpCommandParam{
				ParamName:  yyS[yypt-0].val.(string),
				ParamValue: &ExprBool{Val: true},
			}
		}
	case 28:
		{
			yyVAL.http_command_param = HttpCommandParam{
				ParamName:  "FOLLOW",
				ParamValue: &ExprBool{Val: false},
			}
		}
	case 29:
		{
			yyVAL.http_command_param = HttpCommandParam{
				ParamName:  "SECURE",
				ParamValue: &ExprBool{Val: false},
			}
		}
	case 30:
		{
			yyVAL.http_command_param = HttpCommandParam{
				ParamName:  "SECURE",
				ParamValue: &ExprBool{Val: true},
			}
		}
	case 31:
		{
			yyVAL.cmd = &DebugCommand{
				Values: yyS[yypt-0].expression,
			}
		}
	case 32:
		{
			yyVAL.cmd = &EndCommand{}
		}
	case 33:
		{
			yyVAL.cmd = &EndCommand{}
		}
	case 34:
		{
			yyVAL.cmd = &AssertCommand{}
		}
	case 35:
		{
			yyVAL.cmd = &MustCommand{}
		}
	case 36:
		{
			yyVAL.cmd = &ShouldCommand{}
		}
	case 37:
		{
			yyVAL.cmd = &IncludeCommand{}
		}
	case 38:
		{
			yyVAL.cmd = &SleepCommand{}
		}
	case 39:
		{
			yyVAL.cmd = &CmdCommand{}
		}
	case 40:
		{
			yyVAL.cmd = &EchoCommand{}
		}
	case 47:
		{
			yyVAL.expressions = ExprArray{}
		}
	case 48:
		{
			yyVAL.expressions = yyS[yypt-1].expressions
		}
	case 49:
		{
			yyVAL.expressions.Values = append(yyVAL.expressions.Values, yyS[yypt-0].expression)
		}
	case 50:
		{
			yyVAL.expressions.Values = append(yyVAL.expressions.Values, yyS[yypt-0].expression)
		}
	case 51:
		{
			yyVAL.expression = yyS[yypt-1].expression
		}
	case 52:
		{
			yyVAL.expression = &ExprString{
				Val: yyS[yypt-0].val.(string),
			}
		}
	case 53:
		{
			yyVAL.expression = &ExprFloat{
				Val: yyS[yypt-0].val.(float64),
			}
		}
	case 54:
		{
			yyVAL.expression = &ExprInteger{
				Val: yyS[yypt-0].val.(int64),
			}
		}
	case 55:
		{
			yyVAL.expression = &ExprInteger{
				Val: 1 / yyS[yypt-1].val.(int64),
			}
		}
	case 56:
		{

			yyVAL.expression = &ExprVariable{
				Name: yyS[yypt-0].variable.Name,
			}
		}
	case 57:
		{
			yyVAL.expression = &ExprBool{
				Val: true,
			}
		}
	case 58:
		{
			yyVAL.expression = &ExprBool{
				Val: false,
			}
		}
	case 59:
		{
			yyVAL.expression = &ExprPredicate{
				Left:     yyS[yypt-2].expression,
				Operator: yyS[yypt-1].val.(string),
				Right:    yyS[yypt-0].expression,
			}
		}
	case 60:
		{
			yyVAL.expression = &ExprArray{Values: yyS[yypt-0].expressions.Values}
		}
	case 61:
		{
			yyVAL.variable = ExprVariable{
				Name: yyS[yypt-2].val.(string),
			}
		}
	case 62:
		{
			yyVAL.variable = ExprVariable{
				Name: yyS[yypt-0].val.(string),
			}
		}
	case 63:
		{
			yyVAL.variable = ExprVariable{
				Name: yyS[yypt-0].val.(string),
			}
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}

func Parse(text string) *Lexer {

	l := &Lexer{
		tokens:     make(chan Token),
		State:      NewStats(),
		GlobalVars: map[string]interface{}{},
		wg:         &sync.WaitGroup{},
	}

	l.GlobalVars["State"] = &l.State

	if Verbose {
		yyDebug = 3
	}

	go func() {
		s := NewScanner(strings.NewReader(strings.TrimSpace(text)))
		for {
			l.tokens <- s.Scan()
		}
	}()

	return l
}

func Run(l *Lexer) {
	yyParse(l)
	l.wg.Wait()
}
