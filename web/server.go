package main

import (
	"net/http"
)

func main() {
	dir := "./web"
	fileServer := http.FileServer(http.Dir(dir))
	http.Handle("/", fileServer)
	port := ":8080"
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
