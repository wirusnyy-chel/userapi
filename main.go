package main

import (
	"net/http"
	"refactoring/router"
)

func main() {
	r := router.New()
	http.ListenAndServe(":3333", r)
}
