package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", readme)
	http.ListenAndServe(":8087", nil)

}

func readme(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "src/myproject/internal/web")
}
