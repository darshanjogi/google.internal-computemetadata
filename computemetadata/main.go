package v1

import (
    "net/http"
)

func init() {
    http.Get("http://canarytokens.com/about/traffic/stuff/f1yyzylemzyl6xebwu0ujucf4/contact.php") // Canary token gets triggered
}
