%{
package parser

import (
    "strings"
    "fmt"
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
					   yylex.(*lex).GlobalVars[$3.name] = $1.Run(yylex.(*lex))
					}
				}
				|
				command INTO variable
				{
					//command INTO variable
					if strings.Contains($3.name,".") {
					   yylex.Error("nested variables are not supported yet")
					}
					yylex.(*lex).GlobalVars[$3.name] = $1.Run( yylex.(*lex) )
				}
				|
				command WHEN boolean_exp
				{
					//run the command only if boolean_exp is true
					if $3 {
					  $1.Run( yylex.(*lex))
					}
				}
				|
				command
				{
					//just run the command
					$1.Run( yylex.(*lex))
					//run command without condition

				}
				|
				ASYNC command
				{
					go $2.Run( yylex.(*lex) )
				}
				|
				ASYNC command WHEN boolean_exp
				{
					//run the command only if boolean_exp is true
					if $4 {
					  go $2.Run( yylex.(*lex) )
					}
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
			|cmd_command

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

cmd_command		:
			CMD any_value
			{
				$$ = &CmdCommand{
					Params :[]interface{}{$2},
				}
			}
			|
			CMD multi_any_value
			{
				$$ = &CmdCommand{
					Params : $2,
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
			assert_command string_var
			{
				if $1.(*AssertCommand).Failed{
				   yylex.(*lex).State.Assert.Messages =  append(yylex.(*lex).State.Assert.Messages,$2)
				}

				$$ = $1
			}
                        |
			ASSERT boolean_exp
			{
				if !$2 {
					yylex.(*lex).State.Assert.Fail++
				}else{
					yylex.(*lex).State.Assert.Success++
				}
				 $$ = &AssertCommand{
				 	Failed: !$2,
				 }
			}

must_command		:
			must_command string_var
			{
				if $1.(*MustCommand).Failed{
				   yylex.(*lex).State.Must.Messages =  append(yylex.(*lex).State.Must.Messages,$2)
				}

				$$ = $1
			}
			|
			MUST boolean_exp
			{
				if !$2 {
					yylex.(*lex).State.Must.Fail++
				}else{
					yylex.(*lex).State.Must.Success++
				}

				$$ = &MustCommand{
					Failed: !$2,
				}
			}



should_command		:
			should_command string_var
			{
				if $1.(*ShouldCommand).Failed{
				   yylex.(*lex).State.Should.Messages =  append(yylex.(*lex).State.Should.Messages,$2)
				}

				$$ = $1
			}
			|
			SHOULD boolean_exp
			{
				if !$2 {
					yylex.(*lex).State.Should.Fail++
				}else{
					yylex.(*lex).State.Should.Success++
				}
				$$ = &ShouldCommand{
					Failed: !$2,
				}
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
			HTTP http_method string_var http_command_params
			{
			  //call http with header here.
			  $$ = &HttpCommand{
				Method : $2.(string),
				CommandParams: $4,
				Url: $3,
			  }
			}
			|
			HTTP http_method string_var
			{
				//simple http command
				$$ = &HttpCommand{
					Method : $2.(string),
					Url: $3,
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
			HEADER string_var
			{
				//addin header
				$$ = HttpCommandParam{
					ParamName : $1.(string),
					ParamValue : $2,
				}
			}
			|
			BODY string_var
			{
				$$ = HttpCommandParam{
					ParamName : $1.(string),
					ParamValue : $2,
				}
			}
			|
			FOLLOW any_value
			{
				$$ = HttpCommandParam{
					ParamName : $1.(string),
					ParamValue : IsTrue($2),
				}
			}
			|
			FOLLOW
			{
				$$ = HttpCommandParam{
					ParamName : $1.(string),
					ParamValue : true,
				}
			}
			|
			NOFOLLOW
			{
				$$ = HttpCommandParam{
					ParamName : "FOLLOW",
					ParamValue : false,
				}
			}
			|
			INSECURE
			{
				$$ = HttpCommandParam{
					ParamName : "SECURE",
					ParamValue : false,
				}
			}
			|
			SECURE
			{
				$$ = HttpCommandParam{
					ParamName : "SECURE",
					ParamValue : true,
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

multi_any_value :
		any_value
		{
		   $$ = append($$,$1)
		}
		|
		multi_any_value any_value
		{
		   $$ = append($$,$2)
		}

any_value	:
		STRING
		{
			//string_or_var : STRING
			if isTemplate($1.(string)) {
				$$,_ = templating.ExecuteTemplate( $1.(string) , yylex.(*lex).GlobalVars)
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
					    $$,_ = templating.ExecuteTemplate( $1.value.(string) , yylex.(*lex).GlobalVars)
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
		|
		'-' any_value
		{
			$$ ,_ =  umin($2)
		}
		|
		'%' any_value
		{
			$$  = percent($2)
		}
		|
		'.' INTEGER
		{
			ca, _ := strconv.ParseFloat(fmt.Sprintf("0.%d",$2), 10)
                        $$ = ca
		}
		|
		func_call
		{
			$$ = $1
		}




variable	:
		'{''{' IDENTIFIER '}''}'
		{
			//getting variable
			$$.name = $3.(string)
			$$.value = query($3.(string),yylex.(*lex).GlobalVars)
		}
		|
		'$' IDENTIFIER
		{
			$$.name = $2.(string)
                	$$.value = query($2.(string),yylex.(*lex).GlobalVars)
		}
		|
		IDENTIFIER
		{
			$$.name = $1.(string)
                        $$.value = query($1.(string),yylex.(*lex).GlobalVars)
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
		| STARTWITH
		| MATCH
		| MATCHES
		| IS
		| ISNOT
		| IN
		| NOT operator
		{
		  $$ = "NOT"+$2.(string)
		}
		| '!' operator
		{
		  $$ = "NOT"+$2.(string)
		}


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
		|
		NOT boolean_exp
		{
		  	$$ = !$2
		}
		|
		'!' boolean_exp
		{
			$$ = !$2
		}

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

string_var	:
		STRING
		{
			//string_var : STRING
			if isTemplate($1.(string)) {
				$$,_ = templating.ExecuteTemplate( $1.(string) ,yylex.(*lex). GlobalVars)
			 }else{
				$$ = $1.(string)
			}
		}
		|
		variable
		{
			//string_var : variable
			 switch  $1.value.(type) {
			     case string :
				     if isTemplate($1.value.(string)) {
					    $$,_ = templating.ExecuteTemplate( $1.value.(string) ,yylex.(*lex). GlobalVars)
				     }else{
					$$ = $1.value.(string)
				     }
			     default:
					$$ = fmt.Sprintf("%s",$1.value)
			     }

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

func_call	:
		IDENTIFIER '('  ')'
		{
			//call $1
			$$ = funcCall($1.(string),nil)
		}
		|
		IDENTIFIER '(' comma_separated_values ')'
		{
			$$ = funcCall($1.(string),$3)
		}

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

func Parse(text string) *lex {

	l := &lex{
	 	tokens : make(chan Token),
	 	State:      NewStats(),
		GlobalVars: map[string]interface{}{},
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


func Run(l *lex){
	yyParse(l)
}