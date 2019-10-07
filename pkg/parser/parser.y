%{
package parser

import (
    "strings"
    "fmt"
    "sync"
    "strconv"
    "github.com/microspector/microspector/pkg/templating"
)
%}


%token
EOL
EOF

//scalar type tokens
%token <val>
STRING
INTEGER
FLOAT
TRUE
FALSE
NULL

//command tokens
%token
<val>
KEYWORD
SET
HTTP
MUST
SHOULD
DEBUG
END
ASSERT
INCLUDE
SLEEP
CMD
ASYNC
ECHO

//http command tokens
%token <val>
GET
HEAD
POST
PUT
DELETE
CONNECT
OPTIONS
TRACE
PATCH
HEADER
BODY
FOLLOW
NOFOLLOW
SECURE
INSECURE

//condition tokens
%token <val>
EQUALS
EQUAL
NOTEQUALS
NOTEQUAL
GT
GE
LT
LE
CONTAINS
CONTAIN
STARTSWITH
STARTWITH
WHEN
AND
OR
MATCHES
MATCH
IS
ISNOT
NOT
IN

%token <val>
INTO

%token <val>
IDENTIFIER
TYPE


%type <Variable> variable
%type <val> http_method operator
%type <val> any_value func_call
%type <vals> multi_variable array comma_separated_values multi_any_value
%type <http_command_params> http_command_params
%type <http_command_param> http_command_param
%type <boolean> boolean_exp expr_opr true_false
%type <str> string_var


//arithmetic things
%type <val> expr
%left '|'
%left '&'
%left '+'  '-'
%left '*'  '/'  '%'


%type <cmd>
	command
	set_command
        http_command
	debug_command
	end_command
	assert_command
	must_command
	should_command
	include_command
	sleep_command
	cmd_command
	echo_command

%union{
	val interface{}
	vals []interface{}
	str string
	integer int64
	boolean bool
	bytes []byte
	cmd Command
	Variable struct{
		name string
		value interface{}
	}
	http_command_params []HttpCommandParam
	http_command_param  HttpCommandParam
}

%start statement

%%

expr		:
		'(' expr ')'
		|
		STRING
		|
		FLOAT
		|
		INTEGER
		|
		variable
		|
		TRUE
		|
		FALSE

variable	: '{''{'  IDENTIFIER '}''}'
		|
		'$' IDENTIFIER
		|
		IDENTIFIER




%%

func Parse(text string) *Lexer {

	l := &Lexer{
	 	tokens : make(chan Token),
	 	State:      NewStats(),
		GlobalVars: map[string]interface{}{},
		wg: &sync.WaitGroup{},
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


func Run(l *Lexer){
	yyParse(l)
	l.wg.Wait()
}