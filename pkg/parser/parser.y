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
STRING
TRUE
FALSE

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

%token <val>
INTO

%token <val>
IDENTIFIER


%type <variable> variable
%type <val> http_method operator
%type <val> any_value string_or_var
%type <vals> multi_any_value
%type <http_command_params> http_command_params
%type <http_command_param> http_command_param
%type <boolean> boolean_exp expr_opr


//arithmetic things
%type <val> expr number
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

%union{
	val interface{}
	vals []interface{}
	str string
	integer int
	boolean bool
	flt int64
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
				command INTO variable WHEN boolean_exp {
					//run command put result into variable WHEN boolean_exp is true
					if strings.Contains($3.name,".") {
						yylex.Error("nested variables are not supported yet")
					}

					if $5 {
					   GlobalVars[$3.name] = $1.Run()
					}
				}
				|
				command INTO variable {
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
				command {
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


debug_command		:
			DEBUG multi_any_value {
			  $$ = &DebugCommand{
				Values : $2,
			   }
			}

end_command		:
			END WHEN boolean_exp {
				if $3 {
				   return -1
				 }
				  $$ = &EndCommand{}
			}
			|
			END boolean_exp {
			 if $2 {
			    return -1
			 }

			 $$ = &EndCommand{}
			}
			| END {
				return -1
			}

assert_command		:
			ASSERT boolean_exp {
				if !$2 {
					State.Assertion.Failed++
				}else{
					State.Assertion.Succeeded++
				}
				 $$ = &AssertCommand{}
			}

must_command		:
			MUST boolean_exp {
				if !$2 {
						State.Must.Failed++
					}else{
						State.Must.Succeeded++
					}

				$$ = &MustCommand{}
			}

should_command		:
			SHOULD boolean_exp {
				if !$2 {
						State.Should.Failed++
					}else{
						State.Should.Succeeded++
					}
				$$ = &ShouldCommand{}
			}

set_command		:
			SET variable any_value {
				//GlobalVars[$2.name] = $3
				$$ = &SetCommand{
					Name:$2.name,
					Value:$3,
				}
			}
			|
			SET variable expr {
				//GlobalVars[$2.name] = $3
				$$ = &SetCommand{
					Name:$2.name,
					Value:$3,
				}
			}
			|
			SET variable boolean_exp {
				//GlobalVars[$2.name] = $3
				$$ = &SetCommand{
					Name:$2.name,
					Value:$3,
				}
			}

http_command		:
			HTTP http_method string_or_var http_command_params  {
			  //call http with header here.
			  $$ = &HttpCommand{
				Method : $2.(string),
				CommandParams: $4,
				Url: $3.(string),
			  }
			}
			| HTTP http_method string_or_var {
				//simple http command
				$$ = &HttpCommand{
					Method : $2.(string),
					Url: $3.(string),
				  }
			}

http_command_params	:
			http_command_param {
				if $$ == nil {
				  $$ = make([]HttpCommandParam,0)
				}
				$$ = append($$,$1)
			}
			| http_command_params http_command_param  {
				if $$ == nil {
				  $$ = make([]HttpCommandParam,0)
				}

				$$ = append($$,$2)
			}

http_command_param	:
			HEADER STRING {
				//addin header
				$$ = HttpCommandParam{
					ParamName : $1.(string),
					ParamValue : $2.(string),
				}
			}
			|
			BODY STRING {
				//adding query param
				$$ = HttpCommandParam{
						ParamName : $1.(string),
						ParamValue : $2.(string),
					}
			}


http_method	:
		GET
		{
			//http get
			$$ = $1
		}
		| HEAD
		{
		      $$ = $1
		}
		|POST
		{
		      $$ = $1
		}
		|PUT
		{
		      $$ = $1
		}
		|DELETE
		{
		      $$ = $1
		}
		|CONNECT
		{
		      $$ = $1
		}
		|OPTIONS
		{
		      $$ = $1
		}
		|TRACE
		{
		      $$ = $1
		}
		|PATCH
		{
		      $$ = $1
		}

multi_any_value	:
		any_value
		{
			//getting a single value from multi_value exp
			$$ = append($$,$1)
		}
		|
		multi_any_value any_value {
			//multi value
			$$ = append($$,$2)
		}

any_value	:
		string_or_var
		{
			//any_value: string_or_var
			$$ = $1
		}
		|
		FLOAT
		{
			//any_value: FLOAT
			$$ = $1
		}
		|
		INTEGER
		{
			//any_value: INTEGER
			$$ = $1
		}


string_or_var	:
		variable {
		 //string_or_var : variable
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
		STRING {
			//string_or_var : STRING
			if isTemplate($1.(string)) {
				$$,_ = executeTemplate( $1.(string) , GlobalVars)
			 }else{

			}
		}

variable	:
		'{''{' IDENTIFIER '}''}'{
			//getting variable
			$$.name = $3.(string)
			$$.value = query($3.(string),GlobalVars)
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


boolean_exp	:
		'(' boolean_exp ')'
		{
		 	//boolean_ex: '(' boolean_exp ')'
		  	$$ = $2
		}
		|
		TRUE
		{
			$$ = true
		}
		|
		FALSE
		{
			$$ = false
		}
		|
		boolean_exp AND boolean_exp
		{
			//boolean_ex: boolean_exp AND boolean_exp
		   	$$ = IsTrue($1) && IsTrue($3)
		}
		|
		boolean_exp OR boolean_exp
		{
			//boolean_ex: boolean_exp OR boolean_exp
		   	$$ = IsTrue($1) || IsTrue($3)
		}
		|
		any_value operator any_value
		{
			//boolean_ex: any_value operator any_value, oh conflicts start here :(
			operator_result := runop($1,$2,$3)
			$$ = operator_result
		}
		| expr_opr
expr_opr:
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

// arithmetic things
expr	:    '(' expr ')'
		{ $$  =  $2 }
	|    expr '+' expr
		{ $$,_  = add ( $1 , $3 ) }
	|    expr '-' expr
		{ $$,_  =  subtract($3 , $1) }
	|    expr '*' expr
		{ $$,_  =  multiply($3 , $1) }
	|    expr '/' expr
		{ $$,_  =  divide($3 , $1) }
	|    expr '%' expr
		{ $$,_  =  mod($3 , $1) }
	|    number
	;


number	: INTEGER
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
	variable
	{
	    	//number: variable
		$$ =  $1.value
	}



%%



type lex struct {
    tokens []Token
}

func (l *lex) Lex(lval *yySymType) int {
    if len(l.tokens) == 0 {
        return 0
    }

    v := l.tokens[0]
    l.tokens = l.tokens[1:]
    lval.val = v.Val
    return v.Type
}

func (l *lex) Error(e string) {
    log.Fatal(e)
}

//TODO: use channels here.
func Parse(text string) *lex {

yyXErrors[yyXError{ 1,-1 }] = "deneme error"
yyXErrors[yyXError{ 2,-1 }] = "deneme error"
yyXErrors[yyXError{ 3,-1 }] = "deneme error"
yyXErrors[yyXError{ 4,-1 }] = "deneme error"
yyXErrors[yyXError{ 5,-1 }] = "deneme error"
yyXErrors[yyXError{ 6,-1 }] = "deneme error"
yyXErrors[yyXError{ 7,-1 }] = "deneme error"
yyXErrors[yyXError{ 8,-1 }] = "deneme error"
yyXErrors[yyXError{ 9,-1 }] = "deneme error"
yyXErrors[yyXError{ 10,-1 }] = "10 state error"
yyXErrors[yyXError{ 11,-1 }] = "11 state error"
yyXErrors[yyXError{ 12,-1 }] = "12 state error"
yyXErrors[yyXError{ 13,-1 }] = "13 state error"
yyXErrors[yyXError{ 14,-1 }] = "14 state error"
yyXErrors[yyXError{ 15,-1 }] = "15 state error"
yyXErrors[yyXError{ 16,-1 }] = "16 state error"
yyXErrors[yyXError{ 17,-1 }] = "17 state error"
yyXErrors[yyXError{ 18,-1 }] = "18 state error"
yyXErrors[yyXError{ 19,-1 }] = "19 state error"
yyXErrors[yyXError{ 20,-1 }] = "20 state error"
yyXErrors[yyXError{ 21,-1 }] = "21 state error"
yyXErrors[yyXError{ 22,-1 }] = "22 state error"
yyXErrors[yyXError{ 23,-1 }] = "23 state error"
yyXErrors[yyXError{ 24,-1 }] = "24 state error"
yyXErrors[yyXError{ 25,-1 }] = "25 state error"
yyXErrors[yyXError{ 26,-1 }] = "26 state error"
yyXErrors[yyXError{ 27,-1 }] = "27 state error"
yyXErrors[yyXError{ 28,-1 }] = "28 state error"
yyXErrors[yyXError{ 29,-1 }] = "29 state error"
yyXErrors[yyXError{ 30,-1 }] = "30 state error"
yyXErrors[yyXError{ 31,-1 }] = "31 state error"
yyXErrors[yyXError{ 32,-1 }] = "32 state error"
yyXErrors[yyXError{ 33,-1 }] = "33 state error"
yyXErrors[yyXError{ 34,-1 }] = "34 state error"
yyXErrors[yyXError{ 35,-1 }] = "35 state error"
yyXErrors[yyXError{ 36,-1 }] = "36 state error"
yyXErrors[yyXError{ 37,-1 }] = "37 state error"
yyXErrors[yyXError{ 38,-1 }] = "38 state error"
yyXErrors[yyXError{ 39,-1 }] = "39 state error"
yyXErrors[yyXError{ 40,-1 }] = "40 state error"
yyXErrors[yyXError{ 41,-1 }] = "41 state error"
yyXErrors[yyXError{ 42,-1 }] = "42 state error"
yyXErrors[yyXError{ 43,-1 }] = "43 state error"
yyXErrors[yyXError{ 44,-1 }] = "44 state error"
yyXErrors[yyXError{ 45,-1 }] = "45 state error"
yyXErrors[yyXError{ 46,-1 }] = "46 state error"
yyXErrors[yyXError{ 47,-1 }] = "47 state error"
yyXErrors[yyXError{ 48,-1 }] = "48 state error"
yyXErrors[yyXError{ 49,-1 }] = "49 state error"
yyXErrors[yyXError{ 50,-1 }] = "50 state error"
yyXErrors[yyXError{ 51,-1 }] = "51 state error"
yyXErrors[yyXError{ 52,-1 }] = "52 state error"
yyXErrors[yyXError{ 53,-1 }] = "53 state error"
yyXErrors[yyXError{ 54,-1 }] = "54 state error"
yyXErrors[yyXError{ 55,-1 }] = "55 state error"
yyXErrors[yyXError{ 56,-1 }] = "56 state error"
yyXErrors[yyXError{ 57,-1 }] = "57 state error"
yyXErrors[yyXError{ 58,-1 }] = "58 state error"
yyXErrors[yyXError{ 59,-1 }] = "59 state error"
yyXErrors[yyXError{ 60,-1 }] = "60 state error"
yyXErrors[yyXError{ 61,-1 }] = "61 state error"
yyXErrors[yyXError{ 62,-1 }] = "62 state error"
yyXErrors[yyXError{ 63,-1 }] = "63 state error"
yyXErrors[yyXError{ 64,-1 }] = "64 state error"
yyXErrors[yyXError{ 65,-1 }] = "65 state error"
yyXErrors[yyXError{ 66,-1 }] = "66 state error"
yyXErrors[yyXError{ 67,-1 }] = "67 state error"
yyXErrors[yyXError{ 68,-1 }] = "68 state error"
yyXErrors[yyXError{ 69,-1 }] = "69 state error"
yyXErrors[yyXError{ 70,-1 }] = "70 state error"
yyXErrors[yyXError{ 71,-1 }] = "71 state error"
yyXErrors[yyXError{ 72,-1 }] = "72 state error"
yyXErrors[yyXError{ 73,-1 }] = "73 state error"
yyXErrors[yyXError{ 74,-1 }] = "74 state error"
yyXErrors[yyXError{ 75,-1 }] = "75 state error"
yyXErrors[yyXError{ 76,-1 }] = "76 state error"
yyXErrors[yyXError{ 77,-1 }] = "77 state error"
yyXErrors[yyXError{ 78,-1 }] = "78 state error"
yyXErrors[yyXError{ 79,-1 }] = "79 state error"
yyXErrors[yyXError{ 80,-1 }] = "80 state error"
yyXErrors[yyXError{ 81,-1 }] = "81 state error"
yyXErrors[yyXError{ 82,-1 }] = "82 state error"
yyXErrors[yyXError{ 83,-1 }] = "83 state error"
yyXErrors[yyXError{ 84,-1 }] = "84 state error"
yyXErrors[yyXError{ 85,-1 }] = "85 state error"
yyXErrors[yyXError{ 86,-1 }] = "86 state error"
yyXErrors[yyXError{ 87,-1 }] = "87 state error"
yyXErrors[yyXError{ 88,-1 }] = "88 state error"
yyXErrors[yyXError{ 89,-1 }] = "89 state error"

	s := NewScanner(strings.NewReader(strings.TrimSpace(text)))
	tokens := make(Tokens,0)

	for {
		token := s.Scan()
		//fmt.Println(token)
		if token.Type == EOF || token.Type == -1 {
			break
		}
		tokens = append(tokens, token)
	}
	 l := &lex{ tokens }
	 return l
}

func Reset(){
   	GlobalVars = map[string]interface{}{}
   	State = NewStats()
}

func Run(l *lex){
	yyParse(l)
}