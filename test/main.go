package main

import (
	"net/http"

	"github.com/GoWebFramework/sample-framework"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func main() {
	a := sample.Instance{
		Host: ":8080",
	}
	a.GET("/", hello)
	a.Run()
}
