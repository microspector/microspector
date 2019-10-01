# Microspector
Micro Service Inspector

Microspector is a scripting language designed to test microservices and RESTful APIs in a sexy way.

```
$ microspector --help

Usage of ./microspector:
  -file string
        task file path
  -folder string
        tasks folder path
  -verbose
        print out logs
  -version
        prints version

```

## Scripting

Writing a script in microspector is as simple as
```bash
SET  Url  "https://microspector.com/test.json"
HTTP GET  Url   HEADER "User-Agent:My super duper API test tool" INTO  ApiResult 
MUST  ApiResult.Json.boolean  == true
HTTP POST "https://hooks.slack.com/services/SLACK_TOKEN" 
     HEADER "User-Agent:Microspector
     Content-type: application/json"
     BODY '{ "text":"Oh!:( Microspector API returns false in json.boolean, you broke it!" }'
     INTO  SlackResult  
     WHEN  ApiResult.Json.boolean  != true
```

### Variables
Variables can be used like `VariableName`, `{{ VariableName }}` OR  `$VariableName`. All supported 
in order to combine with any other templating. Brackets are just to allow reserved keywords being used as variable name. 
For example `set $set 1` and `set {{ set }} 1` will set a variable named `set` while `set set 1` failing because set is a command
Variables can be accessed and set in same way. 
Accessing supports nested variables like ` HttpResult.Json.message  ` but setting does not yet.


### Commands
A Microspector script is a basic set of commands. Every command takes some params and does the rest accordingly. 

Currently supported commands are:

1. [SET](#set) 
2. [HTTP](#http) 
3. [ASSERT](#assert) 
4. [MUST](#must) 
5. [SHOULD](#should) 
6. [DEBUG](#debug)
7. [END](#end)
8. [INCLUDE](#include)
9. [SLEEP](#sleep)


#### SET
Set command is used to define and set a variable

Example:
```bash
SET $Url "https://microspector.com"
SET  Url  "https://microspector.com"
SET  MyElements  [1,2,3,4,5,6,7,8]
```


#### HTTP
Http command takes method, url, header and body in order to make an http request

Full Example:
```bash
HTTP POST "https://hooks.slack.com/services/SLACK_TOKEN" 
     HEADER "User-Agent:Microspector
     Accept: application/json
     Content-Type: application/json"
     BODY '{ "text":"Hello World!" }'
     INTO  Result 
```

Basic Example:
```bash
HTTP GET "https://microspector.com/test.json" INTO  result 
```
this command basically tries to fetch the url and put an &HttpResult into result variable
an HttpResult type has a few handy property that can be used for future commands like;

```go
type HttpResult struct {
	Took          int64
	Content       string
	Json          interface{}
	Headers       map[string]string
	StatusCode    int
	ContentLength int
	Error         string
}
```

#### ASSERT
Assert is an assertion command, takes an expression and does a thruty check to mark it is failed or succeeded. Different assertion commands are just to categorize the failures.  

```bash
ASSERT  result.Json.boolean equals true
ASSERT  result.Json.boolean #both works cuz assertion does a thruty check
```

#### MUST
Must is an assertion command, takes an expression and does a thruty check to mark it is failed or succeeded. 

```bash
MUST  result.StatusCode == 200
```

#### SHOULD
Should is an assertion command, takes an expression and does a thruty check to mark it is failed or succeeded  

```bash
SHOULD  result.Took  < 900 
```

#### DEBUG
Debug is used to printout variables with its structure, following example will just print the content to the stdout

```bash
DEBUG  result.Content 
```

#### END
End takes optional boolean expression. It just skips if thruty fails. When its used without parameter or the given expression passes thruty check it ends the execution.
```bash
END WHEN result.Json.boolean equals false
END result.Json.boolean equals false
```

#### INCLUDE
Include takes a file path as a parameter and parses in at runtime in the same context.
```bash
INCLUDE "tasks/sub-commands.msf"
```

#### SLEEP
Sleep takes an integer value in milliseconds and blocks the execution until then
```bash
SLEEP 500
```

### Async Commands
Any command can be run in background using `ASYNC` keyword like;
```bash
ASYNC HTTP POST "https://hooks.slack.com/services/SLACK_TOKEN" 
     HEADER "User-Agent:Microspector
     Accept: application/json
     Content-Type: application/json"
     BODY '{ "text":"Hello World!" }'
```
with `async`, commands does not allow `INTO` they just work and no return for them, async commands are good for making callbacks in background

# Contributing
Any contributions are more than welcome. Create issues as proposal if you have any suggestions. Or even better,
just fork, do your changes and send pull requests. Just make sure all tests pass using `make test` command before sending a pull request.

## TODO
- [x] ~~Support Including files~~
- [x] ~~Support Sleep~~
- [x] ~~Support `IS` operator to check type of a var in json~~
- [x] ~~Support arrays~~
- [ ] Support `IN` operator to  check elements in arrays
- [ ] Print a better summary of the execution
- [ ] Make stats reachable in script
- [ ] Support setting nested variables

## Known issues
- [x] ~~any integer type converts into float64 during querying items from variables because of a weird behavior of~~ (Unmarshalling)[https://golang.org/pkg/encoding/json/#Unmarshal]