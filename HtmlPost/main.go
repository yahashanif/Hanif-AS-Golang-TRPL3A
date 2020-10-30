package main

import (
	fn "Hanif-AS-Golang-TRPL3A/HtmlPost/function"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", fn.RouteIndexGet)
	http.HandleFunc("/process", fn.RouteSubmitPost)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
