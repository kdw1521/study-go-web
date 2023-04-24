package main

import (
	"go-web-server/wandoapp"
	"net/http"
)

func main() {

	http.ListenAndServe(":3000", wandoapp.NewHttpHandler())
}
