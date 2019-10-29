package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", readme)
	fmt.Println("HTTP Server listening on :8087")
	http.ListenAndServe(":8087", nil)
}

func readme(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "src/myproject/internal/web")
}
