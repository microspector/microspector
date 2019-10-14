# Microspector
Micro Service Inspector

Microspector is a scripting language designed to test microservices and RESTful APIs in a sexy way.

```
$ ./microspector --help

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
10. [CMD](#cmd)
11. [ECHO](#echo)


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
HTTP command also has `nofollow` keyword return the first response instead of following redirection and `insecure` to skip ssl verifications

```bash
HTTP POST "https://hooks.slack.com/services/SLACK_TOKEN" 
     HEADER "User-Agent:Microspector
     Accept: application/json
     Content-Type: application/json"
     BODY '{ "text":"Hello World!" }'
     NOFOLLOW # or FOLLOW to follow redirection
     INSECURE # ignore ssl errors for this call
     INTO  Result 
```



```go
type HttpResult struct {
	Took          int64
	Content       string
	Json          interface{}
	Headers       map[string]string
	StatusCode    int
	ContentLength int
	Certificate   Certificate {
		NotAfter int64 // cert expire date in unix format
	}
	Error         string
}
```

#### ASSERT
Assert is an assertion command, takes an expression and does a truthy check to mark it is failed or succeeded. Different assertion commands are just to categorize the failures.  
Assertion commands also get a string at the end of the command to save it as messages in the assertion category to collect failure messages at the end
```bash
ASSERT  result.Json.boolean equals true "boolean field in json body is not true"
ASSERT  result.Json.boolean #both works cuz assertion does a truthy check
```

#### MUST
Must is an assertion command, takes an expression and does a truthy check to mark it is failed or succeeded. 

```bash
MUST  result.StatusCode == 200 "it did not return 200, returned {{ .result.StatusCode }}"
```

#### SHOULD
Should is an assertion command, takes an expression and does a truthy check to mark it is failed or succeeded  

```bash
SHOULD  result.Took  < 900  "loading microspector is too slow"
```

#### DEBUG
Debug is used to printout variables with its structure, following example will just print the content to the stdout

```bash
DEBUG  result.Content 
```

#### END
End takes optional boolean expression. It just skips if truthy fails. When its used without parameter or the given expression passes truthy check it ends the execution.
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

#### CMD
CMD takes first argument as executable path and others as param and simply runs it in os/exec
```bash
CMD 'echo' 'microspector' into output
MUST output equals 'microspector' "output does not equal microspector"
```

#### ECHO
Echo takes first parameter which must be a string or can be converted to a string as format and other as parameters. echo supports go style formatting
```bash
set name "microspector"
set since 2019
echo "name is %s it works since %d" name since
```

##### Async Commands
Any command can be run in background using `ASYNC` keyword like;
```bash
ASYNC HTTP POST "https://hooks.slack.com/services/SLACK_TOKEN"  HEADER "User-Agent:Microspector\nAccept: application/json\nContent-Type: application/json" BODY '{ "text":"Hello World!" }'
```
with `async`, commands does not allow `INTO` they just work and no return for them, async commands are good for making callbacks in background


#### Functions
Microspector supports [functions](pkg/templating/funcs.go) calls but not defining them yet, It has some builtin helper functions like;
```bash

len()
openssl_rand(len int,string "hex|base64")
rand(min, max int)
now()
timestamp()
unix( time )
hash_md5(any)
hash_sha256(any)
url_encode(string)

Example:

set rand_hex openssl_rand(32,"hex") # 87d2f76b2a38b9ee6b8c05875804c6d0cef5eb17b9d9f8c5bab3083be74e5fa8
set hash_md5(rand_hex) # b487c7f25df1575cdf73fa3a213c4026

```
These functions are also template functions, we will make other go template builtin functions and some more reachable in this context.
We are not considering allowing function defining yet since it is not a programming language but we will extend helper functions or allowing plugins to extend functions in future

#### LOOP
Microspector supports LOOPS through any type of array
```bash
set array [1,2,3,4,5,[6.1,6.2,6.3,6.4,6.5]]

loop a in array
    echo "a:%d\n" a when a is integer
    loop b in a when a is array
       echo "b:%.1f\n"  b
    endloop 
endloop

```

Output:
```bash
a:1
a:2
a:3
a:4
a:5
b:6.1
b:6.2
b:6.3
b:6.4
b:6.5
```

#### Reaching stats at runtime
At the end of the execution, microspector provides some stats about assertions called `State`
```go
package parser

type State struct {
    Assert  struct {
                Fail     int
                Success  int
                Messages []string
            }
    Must    struct {
                Fail     int
                Success  int
                Messages []string
            }
    Should  struct {
                Fail     int
                Success  int
                Messages []string
            }
}

```
where you have failed count for must,should and assert commands and the messages. Now, they are reachable at runtime too like;
```bash
set x false
must x equals true
set MustFailedCount State.Must.Failed
```
it is restricted to set a variable called `State` in favor context state.

# [Contributing](CONTRIBUTING.md)

## TODO
- [ ] Support JWT
- [ ] Support date-time operations and duration units
- [ ] Print a better summary of the execution
- [ ] Support setting nested variables

## Known issues
-  