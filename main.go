package main

import (
	"fmt"
	"net/http"
	"os"
    "log"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html := fmt.Sprintf("<h1><p><center>%s</center></p></h1>", os.Getenv("WELCOME"))
		fmt.Fprint(w, html)
	})

    log.Fatal(http.ListenAndServe(":8080", nil))
}