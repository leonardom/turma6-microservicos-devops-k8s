package main

import (
	"fmt"
	"log"
	"net/http"
	"greeting"
)

func handler(res http.ResponseWriter, req *http.Request) {
	message := "Code.education Rocks!"
	fmt.Fprint(res, greeting.Greeting(message))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}