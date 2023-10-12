package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Up and runing")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World, request sended to: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":80", nil)
}
