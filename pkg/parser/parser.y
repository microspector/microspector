%{
package parser

import (
    "fmt"
    "log"
    "encoding/json"
    "github.com/thedevsaddam/gojsonq"
    "strings"
   "github.com/tufanbarisyildirim/microspector/pkg/command"

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
%type <http_command_params> http_command_params
%type <http_command_param> http_command_param
%type <boolean> boolean_exp condition

%union{
	val interface{}
	str string
	integer int
	boolean bool
	flt int64
	bytes []byte
	variable struct{
		name string
		value interface{}
	}
	http_command_params []command.HttpCommandParam
	http_command_param  command.HttpCommandParam
}

%start any_command

%%

any_command:
command_with_condition_opt
| any_command command_with_condition_opt


command_with_condition_opt:
command | command WHEN boolean_exp


command:
set_command
|http_command
|must_command

must_command:
MUST boolean_exp
{
fmt.Println("MUST",$2)
	//if $2 is false, fail
}
| SHOULD boolean_exp
{
fmt.Println("SHOULD",$2)
	//if $2 is false, write a warning
}

set_command:
SET variable any_value {
	globalvars[$2.name] = $3
}
| SET variable {
	yylex.Error("syntax error, please set a valuable type to variable")
}

http_command:
HTTP http_method string_or_var http_command_params INTO variable {
  //call http with header here.
  fmt.Println($1,$2,$3,$4,$5,$6)
  fmt.Println("we will ",$2,$3,"with ",$4," and put results into ",$6.name)
}
| HTTP http_method string_or_var INTO variable {
  //call http put result into variable
}
| HTTP http_method string_or_var http_command_params  {
  //call http with header here.
}
| HTTP http_method string_or_var {
  //just call http here.
}
 HTTP http_method {
   yylex.Error("http command needs a url")
}
| HTTP {
   yylex.Error("http command needs a method")
}

http_command_params:
http_command_param {
	if $$ == nil {
	  $$ = make([]command.HttpCommandParam,0)
	}
	$$ = append($$,$1)
}
| http_command_params http_command_param  {
	if $$ == nil {
	  $$ = make([]command.HttpCommandParam,0)
	}

	$$ = append($$,$2)
}

http_command_param:
HEADER string_or_var {
	//addin header
	$$ = command.HttpCommandParam{
	 	ParamName : $1.(string),
         	ParamValue : $2.(string),
	}
}
|
QUERY string_or_var {
	//adding query param
     	$$ = command.HttpCommandParam{
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


any_value:
string_or_var
{
	$$ = $1
}
| FLOAT { $$ = $1 }
| INTEGER { $$ = $1 }
| boolean_exp {
//boolean value
 	$$ = $1
}
//TODO: add assignemnts here?

string_or_var:
variable {
 //found variable
 fmt.Println($1)
 $$ = $1.value
}
| STRING {
//found string
 fmt.Println($1)
 $$ = $1
}

variable: '{''{' IDENTIFIER '}''}'{
	//getting variable
	$$.name = $3.(string)
	$$.value = query($3.(string),globalvars)
}

operator:
EQUALS
| GT
| LT
| CONTAINS
| STARTSWITH
| AND
| OR


boolean_exp:
TRUE {
//found true
 $$ = true
}
| FALSE {
 $$ = false
}
|
condition {
  $$ = $1
}
| variable {
 $$ = isTrue($1)
}



condition :
any_value operator any_value{
	//what should we do here?
	$$ = runop($1,$2,$3)
}
| '(' condition ')' {
   	$$ = $2
}


%%

func runop(left, operator,right interface{}) bool {
  switch(operator){
  	case "EQUALS":
  		return left == right
  	case "CONTAINS":
        	return strings.Contains(left.(string), right.(string))
  }

  return false
}


func query(fieldPath string, thevars map[string]interface{}) interface{} {
	b, _ := json.Marshal(thevars)
	return gojsonq.New().JSONString(string(b)).Find(fieldPath)
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
		fmt.Printf("Token: %s, Text:%s\n", token.TypeName(), token.Text)
		if token.Type == EOF || token.Type == -1 {
			break
		}
		tokens = append(tokens, token)
	}

	 l := &lex{ tokens }

	yyParse(l)
	fmt.Println(globalvars)
}