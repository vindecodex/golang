package main

import (
	"golang/mysql_crud/controllers"
	"net/http"
)

func main() {
	/********************************
	route /files/ ==> ./
	*********************************/
	mux := http.NewServeMux()
	mux.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("./"))))
	mux.HandleFunc("/", controllers.Welcome)
	http.ListenAndServe(":9000", mux)
}
