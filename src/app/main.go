package main

import (
	. "../handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", MakeHandler(RootHandler))
	http.HandleFunc("/materials/", MakeHandler(MaterialsHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
