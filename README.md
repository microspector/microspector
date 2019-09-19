# Microspector
Micro Service Inspector

Microspector is scripting language designed to test RESTFul APIs in a sexy way.

It is as simple as
```
SET {{ Url }} "https://microspector.com/test.json"
HTTP GET {{ Url }} INTO {{ ApiResult }}
MUST {{ ApiResult.Json.boolean }} == true
```

### Variables
Variables can be used like `{{ VariableName }}` OR  `$VariableName` both are supported.


### Commands
A Microspector script is a basic set of commands. Every command takes some params does the rest accordingly currently supported commands are:

1. [SET](#set) 
2. [HTTP](#http) 
3. [MUST](#must) 
4. [SHOULD](#should) 
5. [ASSERT](#should) 
6. [DEBUG](#debug)
7. [END](#end)


#### SET
Set command is used to define and set a variable

Example:
```bash
SET $Url "https://microspector.com"
SET {{ Url }} "https://microspector.com"
```


#### HTTP
Http command takes method, url, header and body in order to make an http request

Example:
```bash
HTTP GET "https://microspector.com/test.json" INTO {{ result }}
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

#### MUST
Must is an assertion command, takes an expression and does a thruty check to mark it is failed or succeeded. Different assertion commands are just to categorize the failures.

```bash
MUST {{ result.StatusCode }} EQUALS 200
```

#### SHOULD
Must is an assertion command, takes an expression and does a thruty check to mark it is failed or succeeded  

```bash
SHOULD {{ result.Took }} < 900 
```

#### ASSERT
Must is an assertion command, takes an expression and does a thruty check to mark it is failed or succeeded  

```bash
ASSERT {{ result.Json.boolean }} EQUALS true
```

#### DEBUG
Debug is used to printout variables with its structure, following example will just print the content to the stdout

```bash
DEBUG {{ result.Content }}
```

#### END
End takes optional boolean expression it just skips if thruty fails when its used without parameter or given expression passes thruty it ends the execution.
```bash
END WHEN {{ result.Json.boolean }} EQUALS false
```


# Contributing
Any contributions are more than welcome. Create issues as proposal if you have any suggestions. Or even better,
just fork, do your changes and send pull requests. Just make sure all tests passed using `make test` command before sending a pull request.

## TODO
- [ ] Support `IS` operator to check type of a var in json
- [ ] Support arrays