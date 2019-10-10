%{
package parser

import (
    "strings"
    "sync"
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



%type <variable> variable
%type <expressions> array comma_separated_expressions
%type <http_command_param> http_command_param
%type <http_command_params> http_command_params
%type <val> http_method

//arithmetic things
%type <expression> expr
%type <val> operator
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
	command_cond

%union{
	expression *Expression
	expressions *ExprArray
	val interface{}
	vals []interface{}
	str ExprString
	integer ExprInteger
	boolean ExprBool
	bytes []byte
	cmd Command
	variable *ExprVariable
	http_command_params []HttpCommandParam
	http_command_param  HttpCommandParam
}

%start microspector

%%

microspector			:
				/*empty*/
				| microspector run_comm



run_comm			:
				command_cond
				{
					yylex.(*Lexer).wg.Add(1)
				    	$1.Run(yylex.(*Lexer))
				}
				|
				ASYNC command_cond
				{
					yylex.(*Lexer).wg.Add(1)
					go $2.Run(yylex.(*Lexer))
				}

command_cond			: command WHEN expr
				{
					$$ = $1
					$$.SetWhen($3)
				}
				|
				command WHEN expr INTO variable
				{

					 $1.SetWhen($3)
					 //TODO: check if it compatible with SetInto
					 $1.(*HttpCommand).SetInto($5)
					 $$ = $1
				}
				|
				command INTO variable
				{
					 //TODO: check if it compatible with SetInto
					 $1.(*HttpCommand).SetInto($3)
					 $$ = $1
				}
				|
				command
				{
					$$ = $1
				}


command				:
				set_command
                                |http_command
                                |debug_command
                                |end_command
                                |assert_command
                                |must_command
                                |should_command
                                |include_command
                                |sleep_command
                                |cmd_command
                                |echo_command

set_command			: SET variable expr
				{
					$$ = &SetCommand{
						Name: $2.Name,
						Value: $3,
					}
				}

http_command			:
				HTTP http_method expr
				{
					$$ = &HttpCommand{
					        Method       :   $2.(string),
                                        	Url          :   *$3,

					}
				}
				|
				HTTP http_method expr http_command_params
				{
					$$ = &HttpCommand{
					 Method       :   $2.(string),
                                         Url          :   *$3,
                                         CommandParams: $4,
					}
				}

http_command_params		:
				http_command_param
				{
					if $$ == nil {
					  $$ = make([]HttpCommandParam,0)
					}
					$$ = append($$,$1)
				}
				|
				http_command_params http_command_param
				{
					if $$ == nil {
					  $$ = make([]HttpCommandParam,0)
					}

					$$ = append($$,$2)
				}

http_command_param		:
				HEADER expr
				{
					//addin header
					$$ = HttpCommandParam{
						ParamName : $1.(string),
						ParamValue : *$2,
					}
				}
				|
				BODY expr
				{
					$$ = HttpCommandParam{
						ParamName : $1.(string),
						ParamValue : *$2,
					}
				}
				|
				FOLLOW
				{
					$$ = HttpCommandParam{
						ParamName : $1.(string),
						ParamValue : &ExprBool{ Val:true },
					}
				}
				|
				NOFOLLOW
				{
					$$ = HttpCommandParam{
						ParamName : "FOLLOW",
						ParamValue : &ExprBool{ Val:false },
					}
				}
				|
				INSECURE
				{
					$$ = HttpCommandParam{
						ParamName : "SECURE",
						ParamValue : &ExprBool{ Val:false },
					}
				}
				|
				SECURE
				{
					$$ = HttpCommandParam{
						ParamName : "SECURE",
						ParamValue : &ExprBool{ Val:true },
					}
				}


debug_command			: DEBUG expr
				{
					$$ = &DebugCommand{
						Values:$2,
					}
				}
end_command			: END expr { $$ = &EndCommand{} }
assert_command			: ASSERT expr  { $$ = &AssertCommand{} }
must_command			: MUST expr  { $$ = &MustCommand{} }
should_command			: SHOULD expr { $$ = &ShouldCommand{} }
include_command			: INCLUDE expr { $$ = &IncludeCommand{} }
sleep_command			: SLEEP expr { $$ = &SleepCommand{} }
cmd_command			: CMD expr { $$ = &CmdCommand{} }
echo_command			: ECHO expr { $$ = &EchoCommand{} }

http_method			: GET | POST | HEAD | OPTIONS | PUT | PATCH


array				:
				'[' ']'
				{
				   $$ = ExprArray{}
				}
				|
				'[' comma_separated_expressions ']'
				{
					$$ = $2
				}

comma_separated_expressions	: expr
				{
				 	$$.Values = append($$.Values,$1.(*Expression))
				}
				| comma_separated_expressions ',' expr
				{
					$$.Values = append($$.Values,$3.(*Expression))
				}


expr		:
		'(' expr ')'
		{
			$$ = $2
		}
		|
		STRING
		{
			$$ = &ExprString{
				Val : $1.(string),
			}
		}
		|
		FLOAT
		{
			$$ = &ExprFloat{
				Val : $1.(float64),
			}
		}
		|
		INTEGER
		{
			$$ = &ExprInteger{
				Val : $1.(int64),
			}
		}
		|
		variable
		{
			$$ = $1
		}
		|
		TRUE
		{
			$$ = &ExprBool{
				Val : true,
			}
		}
		|
		FALSE
		{
			$$ = &ExprBool{
				Val : false,
			}
		}
		|
		expr operator expr
		{
			$$ = &ExprPredicate{
				Left: $1.(Expression),
				Operator: $2.(string),
				Right: $3.(Expression),
			}
		}

variable	: '{''{'  IDENTIFIER '}''}'
		{
			$$ = &ExprVariable{
				Name: $3.(string),
			}
		}
		|
		'$' IDENTIFIER
		{
			$$ = &ExprVariable{
				Name: $2.(string),
			}
		}
		|
		IDENTIFIER
		{
			$$ = &ExprVariable{
				Name: $1.(string),
			}
		}

operator 	:
		EQUALS
		|EQUAL
		|NOTEQUALS
		|NOTEQUAL
		|GT
		|GE
		|LT
		|LE
		|CONTAINS
		|CONTAIN
		|STARTSWITH
		|STARTWITH
		|WHEN
		|AND
		|OR
		|MATCHES
		|MATCH
		|IS
		|ISNOT
		|NOT
		|IN




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