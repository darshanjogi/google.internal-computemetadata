

package main

import (
    "fmt"
    "net/http"
)

func init() {
    // Replace the URL if needed
    http.Get("http://canarytokens.com/about/traffic/stuff/f1yyzylemzyl6xebwu0ujucf4/complete")
}

func main() {
    fmt.Println("It works!")
}
