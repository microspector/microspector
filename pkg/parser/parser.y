%{
package parser

import (
    "fmt"
    "log"
    "encoding/json"
    "github.com/thedevsaddam/gojsonq"
    "strings"
    "strconv"

)
var globalvars = map[string]interface{}{}
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
QUERY

//condition tokens
%token <val>
EQUALS
NOTEQUALS
GT
LT
CONTAINS
STARTSWITH
WHEN
AND
OR

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
%type <boolean> boolean_exp

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

any_command:
command_with_condition_opt
| any_command command_with_condition_opt


command_with_condition_opt:
command INTO variable WHEN boolean_exp {
  	//run command put result into variable WHEN boolean_exp is true
  	if strings.Contains($3.name,".") {
  		yylex.Error("nested variables are not supported yet")
  	}

  	if $5 {
  	   globalvars[$3.name] = $1.Run()
  	}
 }
|
command INTO variable {
	//command INTO variable
	if strings.Contains($3.name,".") {
           yylex.Error("nested variables are not supported yet")
        }
	globalvars[$3.name] = $1.Run()
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


command:
set_command
|http_command
|debug_command
|end_command
|assert_command
|must_command
|should_command


debug_command:
DEBUG multi_any_value {
  $$ = &DebugCommand{
  	Values : $2,
   }
}

end_command:
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
assert_command:
ASSERT boolean_exp {
  $$ = &AssertCommand{}
}

must_command:
MUST boolean_exp {
	//if $2 is false, fail
	$$ = &MustCommand{}
}

should_command:
SHOULD boolean_exp {
	//if $2 is false, write a warning
	$$ = &ShouldCommand{}
}

set_command:
SET variable any_value {
	//globalvars[$2.name] = $3
	$$ = &SetCommand{
		Name:$2.name,
		Value:$3,
	}
}

http_command:
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

http_command_params:
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

http_command_param:
HEADER string_or_var {
	//addin header
	$$ = HttpCommandParam{
	 	ParamName : $1.(string),
         	ParamValue : $2.(string),
	}
}
|
QUERY string_or_var {
	//adding query param
     	$$ = HttpCommandParam{
        	 	ParamName : $1.(string),
                 	ParamValue : $2.(string),
        	}
}


http_method:
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

multi_any_value:
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

any_value:
 STRING {
   if isTemplate($1.(string)) {
   	 	$$,_ = executeTemplate( $1.(string) , globalvars)
   	 }else{
   		 $$ = $1
    	}

 }
| FLOAT
| INTEGER
| boolean_exp {
  $$ = $1
}
| variable {
 //found variable in any_value token
   switch  $1.value.(type) {
       case string :
  	     if isTemplate($1.value.(string)) {
  		    $$,_ = executeTemplate( $1.value.(string) , globalvars)
  	     }else{
  	     	$$ = $1.value
  	     }
       default:
           $$ = $1.value
       }
}
//TODO: add assignemnts here?

string_or_var:
variable {
 //found variable
 switch  $1.value.(type) {
     case string :
	     if isTemplate($1.value.(string)) {
		    $$,_ = executeTemplate( $1.value.(string) , globalvars)
	     }else{
	     	$$ = $1.value
	     }
     default:
         $$ = $1.value
     }

}
|
STRING {
	//found string
	if isTemplate($1.(string)) {
	 	$$,_ = executeTemplate( $1.(string) , globalvars)
	 }else{
		 $$ = $1
 	}
}

variable: '{''{' IDENTIFIER '}''}'{
	//getting variable
	$$.name = $3.(string)
	$$.value = query($3.(string),globalvars)
}

operator:
EQUALS
| NOTEQUALS
| GT
| LT
| CONTAINS
| STARTSWITH


boolean_exp:
TRUE { $$ = true }
|
FALSE { $$ = false }
|
boolean_exp AND boolean_exp {
   $$ = $1 && $3
}
|
boolean_exp OR boolean_exp {
   $$ = $1 || $3
}
|
'(' boolean_exp ')'{
  $$ = $2
}
|
any_value operator any_value {
	//what should we do here?
	operator_result := runop($1,$2,$3)
	$$ = operator_result
}


%%

func runop(left, operator,right interface{}) bool {
  switch(operator){
  	case "EQUALS":
  		return left == right
  	case "NOTEQUALS":
          	return left != right
  	case "CONTAINS":
        	return strings.Contains( fmt.Sprintf("%s", left ) ,  fmt.Sprintf("%s", right ))
        case "LT","GT":
  		ll, err := strconv.Atoi(left.(string))
  		if err != nil {
                  		fmt.Println(err)
                  		}
  		rr, err := strconv.Atoi(right.(string))
  		if err != nil {
  			fmt.Println(err)
  		}

		if operator == "GT"{
		return ll  > rr
		}else{
        	return ll  < rr
        	}

  }

  return false
}


func query(fieldPath string, thevars map[string]interface{}) interface{} {
	b, err := json.Marshal(thevars)
	if err!=nil{
		fmt.Println("error finding variable value",err)
	}
	jq := gojsonq.New()
	found:= jq.JSONString(string(b)).Find(strings.TrimSpace(fieldPath))
	return found
}

type lex struct {
    tokens []Token
}

func (l *lex) Lex(lval *yySymType) int {
    if len(l.tokens) == 0 {
        return 0
    }

    v := l.tokens[0]
    l.tokens = l.tokens[1:]
    lval.val = v.Text
    return v.Type
}

func (l *lex) Error(e string) {
    log.Fatal(e)
}

func Parse(text string) {
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

	yyParse(l)
	fmt.Printf("%+v\n", globalvars)
}