package main

import (
	"fmt"
	"net/http"
)

type wandoHandler struct{}

func (wa *wandoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Wando Handler")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	http.Handle("/wando", &wandoHandler{})

	http.ListenAndServe(":3000", nil)
}
