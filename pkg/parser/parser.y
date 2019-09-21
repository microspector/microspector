%{
package parser

import (
    "log"
    "strings"

)
var GlobalVars = map[string]interface{}{}
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
WHEN
AND
OR
MATCHES
MATCH
IS
ISNOT

%token <val>
INTO

%token <val>
IDENTIFIER
TYPE


%type <variable> variable
%type <val> http_method operator
%type <val> any_value
%type <vals> multi_variable array comma_separated_values
%type <http_command_params> http_command_params
%type <http_command_param> http_command_param
%type <boolean> boolean_exp expr_opr true_false


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

%union{
	val interface{}
	vals []interface{}
	str string
	integer int64
	boolean bool
	bytes []byte
	cmd Command
	variable struct{
		name string
		value interface{}
	}
	http_command_params []HttpCommandParam
	http_command_param  HttpCommandParam
}

%start any_command

%%

any_command		:
			/*empty*/
			| any_command command_with_condition_opt


command_with_condition_opt	:
				command INTO variable WHEN boolean_exp
				{
					//run command put result into variable WHEN boolean_exp is true
					if strings.Contains($3.name,".") {
						yylex.Error("nested variables are not supported yet")
					}

					if $5 {
					   GlobalVars[$3.name] = $1.Run()
					}
				}
				|
				command INTO variable
				{
					//command INTO variable
					if strings.Contains($3.name,".") {
					   yylex.Error("nested variables are not supported yet")
					}
					GlobalVars[$3.name] = $1.Run()
				}
				|
				command WHEN boolean_exp
				{
					//run the command only if boolean_exp is true
					if $3 {
					  $1.Run()
					}
				}
				|
				command
				{
					//just run the command
					$1.Run()
					//run command without condition

				}


command			:
			set_command
			|http_command
			|debug_command
			|end_command
			|assert_command
			|must_command
			|should_command
			|include_command
			|sleep_command

sleep_command		:
			SLEEP any_value
			{
				$$ = &SleepCommand{
					Millisecond:intVal($2),
				}
			}

include_command		:
			INCLUDE any_value
			{
				$$ = &IncludeCommand{
					File :  $2.(string),
				}
			}


debug_command		:
			DEBUG multi_variable
			{
				$$ = &DebugCommand{
					Values : $2,
				}
			}

end_command		:
			END WHEN boolean_exp
			{
				if $3 {
				   return -1
				 }

				$$ = &EndCommand{}
			}
			|
			END boolean_exp
			{
				 if $2 {
				    return -1
				 }

				 $$ = &EndCommand{}
			}
			| END
			{
				return -1
			}

assert_command		:
			ASSERT boolean_exp
			{
				if !$2 {
					State.Assertion.Failed++
				}else{
					State.Assertion.Succeeded++
				}
				 $$ = &AssertCommand{}
			}

must_command		:
			MUST boolean_exp
			{
				if !$2 {
						State.Must.Failed++
					}else{
						State.Must.Succeeded++
					}

				$$ = &MustCommand{}
			}



should_command		:
			SHOULD boolean_exp
			{
				if !$2 {
						State.Should.Failed++
					}else{
						State.Should.Succeeded++
					}
				$$ = &ShouldCommand{}
			}

set_command		:
			SET variable array
			{
				$$ = &SetCommand{
					Name:$2.name,
					Value:$3,
				}
			}
			|
			SET variable expr
			{
				//GlobalVars[$2.name] = $3
				$$ = &SetCommand{
					Name:$2.name,
					Value:$3,
				}
			}
			|
			SET variable boolean_exp
			{
				//GlobalVars[$2.name] = $3
				$$ = &SetCommand{
					Name:$2.name,
					Value:$3,
				}
			}

http_command		:
			HTTP http_method any_value http_command_params
			{
			  //call http with header here.
			  $$ = &HttpCommand{
				Method : $2.(string),
				CommandParams: $4,
				Url: $3.(string),
			  }
			}
			|
			HTTP http_method any_value
			{
				//simple http command
				$$ = &HttpCommand{
					Method : $2.(string),
					Url: $3.(string),
				  }
			}

http_command_params	:
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

http_command_param	:
			HEADER STRING
			{
				var passedValue string
				if isTemplate($2.(string)) {
					passedValue,_ = executeTemplate( $2.(string) , GlobalVars)
				 }else{
					passedValue = $2.(string)
				}

				//addin header
				$$ = HttpCommandParam{
					ParamName : $1.(string),
					ParamValue : passedValue,
				}
			}
			|
			BODY STRING
			{
				var passedValue string
				if isTemplate($2.(string)) {
					passedValue,_ = executeTemplate( $2.(string) , GlobalVars)
				 }else{
					passedValue = $2.(string)
				}

				//adding query param
				$$ = HttpCommandParam{
						ParamName : $1.(string),
						ParamValue : passedValue,
					}
			}


http_method	: GET | HEAD | POST | PUT | DELETE | CONNECT | OPTIONS | TRACE | PATCH



multi_variable	:
		variable
		{
			//getting a single value from multi_value exp
			$$ = append($$,$1)
		}
		|
		multi_variable variable
		{
			//multi value
			$$ = append($$,$2)
		}

any_value	:
		STRING
		{
			//string_or_var : STRING
			if isTemplate($1.(string)) {
				$$,_ = executeTemplate( $1.(string) , GlobalVars)
			 }else{
				$$ = $1.(string)
			}
		}
		|
		variable
		{
			//any_value : variable
			 switch  $1.value.(type) {
			     case string :
				     if isTemplate($1.value.(string)) {
					    $$,_ = executeTemplate( $1.value.(string) , GlobalVars)
				     }else{
					$$ = $1.value
				     }
			     default:
					$$ = $1.value
			     }

		}
		|
		INTEGER
                {
			//number: INTEGER
			$$ = $1;
                }
                |
                FLOAT
                {
			//number: FLOAT
			$$ = $1
                }
		|
		TYPE
		|
		NULL
		{
			$$ = nil
		}
		|
		array
		{
			$$ = $1
		}



variable	:
		'{''{' IDENTIFIER '}''}'
		{
			//getting variable
			$$.name = $3.(string)
			$$.value = query($3.(string),GlobalVars)
		}
		|
		'$' IDENTIFIER
		{
			$$.name = $2.(string)
                	$$.value = query($2.(string),GlobalVars)
		}
		|
		IDENTIFIER
		{
			$$.name = $1.(string)
                        $$.value = query($1.(string),GlobalVars)
		}

operator	:
		EQUALS
		| EQUAL
		| NOTEQUALS
		| NOTEQUAL
		| GT
		| GE
		| LT
		| LE
		| CONTAINS
		| CONTAIN
		| STARTSWITH
		| MATCH
		| MATCHES
		| IS
		| ISNOT


boolean_exp	:
		'(' boolean_exp ')'
		{
		 	//boolean_ex: '(' boolean_exp ')'
		  	$$ = $2
		}
		|
		true_false
		|
		any_value operator any_value
		{
			//boolean_ex: any_value operator any_value, oh conflicts start here :(
			operator_result := runop($1,$2,$3)
			$$ = operator_result
		}
		|
		any_value operator true_false
		{
			//boolean_ex: any_value operator any_value, oh conflicts start here :(
			operator_result := runop($1,$2,$3)
			$$ = operator_result
		}
		|
		expr_opr

true_false	:
		TRUE
		{
			$$ = true
		}
		|
		FALSE
		{
			$$ = false
		}

expr_opr	:
		'(' expr_opr ')'
		{
			$$ = $2
		}
		|
		expr operator expr
		{
			operator_result := runop($1,$2,$3)
			$$ = operator_result
		}
		|
		any_value operator expr
		{
			operator_result := runop($1,$2,$3)
                         $$ = operator_result
		}
		|
		expr operator any_value
		{
			operator_result := runop($1,$2,$3)
			 $$ = operator_result
		}
		|
		boolean_exp AND boolean_exp
		{
			//boolean_ex: boolean_exp AND boolean_exp
			$$ = $1 && $3
		}
		|
		boolean_exp OR boolean_exp
		{
			//boolean_ex: boolean_exp OR boolean_exp
			$$ = $1 || $3
		}

expr	:
	'(' expr ')'
		{ $$  =  $2 }
	|    expr '+' expr
		{ $$,_  = add ($1 , $3) }
	|    expr '-' expr
		{ $$,_  =  subtract($3 , $1) }
	|    expr '*' expr
		{ $$,_  =  multiply($3 , $1) }
	|    expr '/' expr
		{ $$,_  =  divide($3 , $1) }
	|    expr '%' expr
		{ $$,_  =  mod($3 , $1) }
	|    any_value
	;


array	:
	'['']'
	{
		$$ = make([]interface{},0)
	}
	|
	'[' comma_separated_values ']'
	{
		$$ = $2
	}

comma_separated_values	:
			any_value
			{
			    $$ = append($$,$1)
			}
			|
			comma_separated_values ',' any_value
			{
			   $$ = append($$,$3)
			}


%%


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

		tokens = append(tokens,v)
	}

	return tokens
}

func (l *lex) Lex(lval *yySymType) int {
    v := <- l.tokens
    if v.Type == EOF || v.Type == -1{
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
	 tokens : make(chan Token),
	}

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

//Resets the state to start over
func Reset(){
   	GlobalVars = map[string]interface{}{}
   	State = NewStats()
}

func Run(l *lex){
	yyParse(l)
}