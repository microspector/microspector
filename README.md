# Microspector
Micro Service Inspector

Microspector is scripting language designed to test RESTFul APIs in a sexy way.

It is as simple as
```
SET {{ Url }} "https://microspector.com/test.json"
HTTP GET {{ Url }} INTO {{ ApiResult }}
MUST {{ ApiResult.Json.boolean }} == true
```

