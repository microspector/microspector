%{
package parser

import (
    "strings"
    "sync"
    "strconv"
    "fmt"
    "reflect"
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
LOOP
ENDLOOP

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
'+' '-' '/' '%' '*'

%token <val>
INTO

%token <val>
IDENTIFIER
TYPE



%type <variable> variable
%type <expressions> array comma_separated_expressions multi_expressions
%type <http_command_param> http_command_param
%type <http_command_params> http_command_params
%type <val> http_method

//arithmetic things
%type <expression> expr func_call predicate_expr math_expression
%type <val> operator operator_math
%type <cmd_list> command_list

%left ANR OR

%left EQUALS
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
      MATCHES
      MATCH
      IS
      ISNOT
      NOT
      IN

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
	comm_in_loop

%union{
	expression Expression
	expressions ExprArray
	val interface{}
	vals []interface{}
	str ExprString
	integer ExprInteger
	boolean ExprBool
	bytes []byte
	cmd Command
	cmd_list []Command
	variable ExprVariable
	http_command_params []HttpCommandParam
	http_command_param  HttpCommandParam
}

%start microspector

%%

microspector			:
				/*empty*/
				| microspector command_list
				{
					for _,cm := range $2{
						yylex.(*Lexer).wg.Add(1)
						if cm.IsAsync() {
						  go cm.Run(yylex.(*Lexer))
						}else{
							r:= cm.Run(yylex.(*Lexer))
							if r == ErrStopExecution{
								return -1
							}
						}
					}
				}




command_list			:
				command_cond
				{
					$$ = append($$,$1)
				}
				|
				command_list command_cond
				{
					$$ = append($$,$2)
				}



command_cond			:
				command WHEN predicate_expr
				{
					$1.SetWhen($3)
					$$ = $1
				}
				|
				command INTO variable WHEN predicate_expr
				{

					 $1.SetWhen($5)
					 $1.(IntoCommand).SetInto($3.Name)
					 $$ = $1
				}
				|
				command INTO variable
				{
					 //TODO: check if it compatible with SetInto
					 $1.(IntoCommand).SetInto($3.Name)
					 $$ = $1
				}
				|
				command
				{
					$$ = $1
				}
				|
				ASYNC command_cond
				{
					$2.SetAsync(true)
					$$ = $2
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
                                |comm_in_loop

set_command			:
				SET variable expr
				{
					$$ = &SetCommand{
						Name: $2.Name,
						Value: $3,
					}
				}
				|
				SET variable predicate_expr
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
                                        	Url          :   $3,

					}
				}
				|
				HTTP http_method expr http_command_params
				{
					$$ = &HttpCommand{
					 Method       :   $2.(string),
                                         Url          :   $3,
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
						ParamValue : $2,
					}
				}
				|
				BODY expr
				{
					$$ = HttpCommandParam{
						ParamName : $1.(string),
						ParamValue : $2,
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

comm_in_loop			:
				LOOP variable IN variable command_list ENDLOOP
				{
					$$ = &LoopCommand{
						Var : $2,
						In : $4,
						Commands : $5,
					}

				}

debug_command			: DEBUG multi_expressions
				{
					$$ = &DebugCommand{
						Values:$2,
					}
				}
end_command			: END predicate_expr { $$ = &EndCommand{ Expr:$2 } }
				| END { $$ = &EndCommand{}  }
assert_command			: ASSERT predicate_expr  { $$ = &AssertCommand{ Expr:$2 } }
				| ASSERT predicate_expr expr { $$ = &AssertCommand{ Expr:$2,Message:$3 } }
must_command			: MUST predicate_expr  { $$ = &MustCommand{ Expr:$2 } }
				| MUST predicate_expr expr { $$ = &MustCommand{ Expr:$2,Message:$3 } }
should_command			: SHOULD predicate_expr { $$ = &ShouldCommand{ Expr:$2 } }
				| SHOULD predicate_expr expr { $$ = &ShouldCommand{ Expr:$2,Message:$3 } }
include_command			: INCLUDE expr { $$ = &IncludeCommand{ Expr:$2 } }
sleep_command			: SLEEP expr { $$ = &SleepCommand{ Expr:$2 } }
cmd_command			: CMD multi_expressions { $$ = &CmdCommand{ Params:$2 } }
echo_command			: ECHO expr { $$ = &EchoCommand{Expr:$2} }
				| ECHO expr multi_expressions { $$ = &EchoCommand{Expr:$2,Params:$3} }

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
				 	$$.Values = append($$.Values,$1)
				}
				| comma_separated_expressions ',' expr
				{
					$$.Values = append($$.Values,$3)
				}

multi_expressions		: expr
				{
					$$.Values = append($$.Values,$1)
				}
				| multi_expressions  expr
				{
					$$.Values = append($$.Values,$2)
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
		math_expression
		{
			$$ = $1
		}
		|
		array
		{
			$$ =  &ExprArray{ Values: $1.Values }
		}
		|
		func_call
		{
		   $$ = $1
		}
		|
		TYPE
		{
		   $$ = &ExprType{
		   	Name: $1.(string),
		   }
		}

math_expression	:
		math_expression '+' math_expression
		{
			$$ = &ExprArithmetic{
				Left: $1,
				Operator: "+",
				Right: $3,
			}
		}
		| math_expression '-' math_expression
		{
			$$ = &ExprArithmetic{
				Left: $1,
				Operator: "-",
				Right: $3,
			}
		}
		| math_expression '*' math_expression
		{
			$$ = &ExprArithmetic{
				Left: $1,
				Operator: "*",
				Right: $3,
			}
		}
		| math_expression '/' math_expression
		{
			$$ = &ExprArithmetic{
				Left: $1,
				Operator: "/",
				Right: $3,
			}
		}
		| '(' math_expression ')'
		{
			$$ = $2
		}
		| INTEGER
		{
			$$ = &ExprInteger{
				Val: $1.(int64),
			}
		}
		| FLOAT
		{
			$$ = &ExprFloat{
				Val : $1.(float64),
			}
		}
		|
		'%' INTEGER
		{
			$$ = &ExprFloat{
				Val: float64( $2.(int64) ) / 100,
			}
		}
		|
		'-' INTEGER
		{
			$$ = &ExprInteger{
				Val: - $2.(int64),
			}
		}
		|
		'.' INTEGER
		{
			ca, _ := strconv.ParseFloat(fmt.Sprintf("0.%d",$2), 10)
			$$ = &ExprFloat{
				Val : ca,
			}

		}
		|
		variable
		{
			$$ =  &ExprVariable{
				Name: $1.Name,
			}
		}



predicate_expr	:
		variable
		{
			// convert variable to a predicate expression
			$$ = &ExprPredicate{
				Left: &ExprVariable{
					Name: $1.Name,
				},
				Operator: "equals",
				Right: &ExprBool{
					Val : true,
				},
			}
		}
		|
		'(' predicate_expr ')'
		{
			$$ = $2
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
				Left: $1,
				Operator: $2.(string),
				Right: $3,
			}
		}
		|
		predicate_expr AND predicate_expr
		{
			$$ = &ExprPredicate{
				Left: $1,
				Operator: "AND",
				Right: $3,
			}
		}
		|
		predicate_expr OR predicate_expr
		{
			$$ = &ExprPredicate{
				Left: $1,
				Operator: "OR",
				Right: $3,
			}
		}
		|
		expr operator predicate_expr{
			$$ = &ExprPredicate{
				Left: $1,
				Operator:  $2.(string),
				Right: $3,
			}
		}
		|
		predicate_expr operator expr{
			$$ = &ExprPredicate{
				Left: $1,
				Operator:  $2.(string),
				Right: $3,
			}
		}
		|
		NOT predicate_expr
		{
			$$ = $2

			t := reflect.TypeOf($2)

			if t ==  reflect.TypeOf(&ExprBool{}).Elem() {
				$2.(*ExprBool).Val  = !($2.(*ExprBool).Val)
			} else if t == reflect.TypeOf(&ExprPredicate{})  {
				$2.(*ExprPredicate).Not  = !($2.(*ExprPredicate).Not)
			}else{
				  $$  = &ExprBool{
					Val : !IsTrue($2),
				 }
			}
		}





func_call	:
		IDENTIFIER '('  ')'
		{
			//call $1
			$$ = &ExprFunc{
				Name:$1.(string),
				Params : nil,
			}
		}
		|
		IDENTIFIER '(' comma_separated_expressions ')'
		{
			$$ = &ExprFunc{
				Name:$1.(string),
				Params : $3.Values,
			}
		}


variable	: '{''{'  IDENTIFIER '}''}'
		{
			$$ = ExprVariable{
				Name: $3.(string),
			}
		}
		|
		'$' IDENTIFIER
		{
			$$ = ExprVariable{
				Name: $2.(string),
			}
		}
		|
		IDENTIFIER
		{
			$$ = ExprVariable{
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
		|MATCHES
		|MATCH
		|IS
		|ISNOT
		|NOT
		|IN
		| NOT operator
		{
			$$  = "NOT" + $2.(string)
		}
		| '!' operator
		{
			$$  = "NOT" + $2.(string)
		}

operator_math	:
		'*'
		|'/'
		|'+'
		|'-'
		|'%'




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