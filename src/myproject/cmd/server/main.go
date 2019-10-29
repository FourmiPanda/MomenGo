package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

// You need to do a ng build before launching the server
var httpPort = "2020"
// local file : MomenGo/src/myproject/cmd/server/main.go
// ng folder  : MomenGo/src/client/dist/client
var folderDist = "../../../client/dist/client" // ng build output folder
func serverHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := os.Stat(folderDist + r.URL.Path); err != nil {
		http.ServeFile(w, r, folderDist+"/index.html")
		return
	} else {
		log.Println(err)
	}
	http.ServeFile(w, r, folderDist+r.URL.Path)
}

func main() {
	r := mux.NewRouter()
	r.NotFoundHandler = r.NewRoute().HandlerFunc(serverHandler).GetHandler()
	http.Handle("/lol", r)
	err := http.ListenAndServe(":"+httpPort, nil)
	if err != nil {
		log.Println(err)
	}
}