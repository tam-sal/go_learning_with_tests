package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreetingHandler(res http.ResponseWriter, req *http.Request) {
	Greet(res, "world")
}

func main() {
	buff := bytes.Buffer{}
	Greet(&buff, "John")
	fmt.Println(buff.String())
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreetingHandler)))
}
