package main

import "fmt"
import "net/http"

func SimpleIndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s!", r.URL.Path[1:])
}

func HttpFileHandler(response http.ResponseWriter, request *http.Request) {
	//fmt.Fprintf(w, "Hi from e %s!", r.URL.Path[1:])
	http.ServeFile(response, request, "Index.html")
}

func main() {

	fmt.Println("Server Starting")
	http.HandleFunc("/", SimpleIndexHandler)
	http.HandleFunc("/index", HttpFileHandler)

	//http.HandleFunc("/", indexTemplateHandler)

	http.ListenAndServe(":8080", nil)
}
