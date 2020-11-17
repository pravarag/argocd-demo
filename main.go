package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", ServerMain)
	http.ListenAndServe(":8080", nil)
	
}

func ServerMain(w.http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, %s!", r.URL.Path[1:])
}
