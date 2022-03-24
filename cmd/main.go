package main

import (
	"fmt"
	"go-api/lib/utility"
	"net/http"
)

func handler(web http.ResponseWriter, req *http.Request) {
	web.WriteHeader(200)
	web.Header().Set("Content-Type", "text/html; charset=utf8")
	fmt.Fprintf(web, "<h1>Hello world!</h1>")
}

func main() {
	utility.LoadEnv()

	http.HandleFunc("/", handler)
	fmt.Println("server - http://localhost:3000")

	http.ListenAndServe(":3000", nil)
}
