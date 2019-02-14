# newsapi-go
A simple and lightweight client for [News API](https://newsapi.org/docs/).

## About
This is a client library written in Go for the News API v2. The functions in it mirror the endpoints exposed by the API, as seen in the [documentation](https://newsapi.org/docs/endpoints). You simply pass the query parameters into the function you need and a JSON response of type `[]byte` will be returned.

## Install
Just like with any other Go package, you can use:
```
go get github.com/thealexcons/newsapi-go
```

## Available Functions/Endpoints
Below are the three endpoints made available by the API. Most (if not all) of the parameters are optional, so you can simply pass in an empty string or 0 if the parameter is an int.
```go
GetTopHeadlines(query string, sources string, language string, country string, category string, pageSize int, page int) []byte

GetEverything(query string, sources string, domains string, excludeDomains string, from string, to string, language string, sortBy string, pageSize int, page int) []byte

GetSources(category string, language string, country string) []byte
```
As you can see, all three functions return `[]byte`. You can then deal with this JSON however you like.

## Example Usage
Here is an example program using the [Top Headlines endpoint](https://newsapi.org/docs/endpoints/top-headlines):
```go
package main

import (
    "github.com/thealexcons/newsapi-go"
    "fmt"
)

func main() {
    
    client := newsapi.Client{ ApiKey: "XXXXXXXXXXXXXXXX" }  // your API key goes here
    
    resp := client.GetTopHeadlines("man city", "", "en", "", "sports", 5, 1)
    
    fmt.Prinln(string(resp))

}
```

## Support
Feel free to make suggestions or provide feedback regarding the code. Please note that this is an unofficial library, so any questions you may have about the API itself must be directed to the great people behind News API. Thanks.
