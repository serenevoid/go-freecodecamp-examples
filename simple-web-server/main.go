package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
    fileserver := http.FileServer(http.Dir("./static/"))
    http.Handle("/", fileserver)
    http.HandleFunc("/form", formHandler)
    http.HandleFunc("/hello", helloHandler)

    fmt.Print("Starting simple web server at port 8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/hello" {
        http.Error(w, "404 not found", http.StatusNotFound)
        return
    }
    if r.Method != "GET" {
        http.Error(w, "Method not supported", http.StatusNotFound)
        return
    }
    fmt.Println(w, "Hello")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        log.Fatal(w, "ParseForm() err: %v", err)
        return
    }
    fmt.Fprint(w, "Post Request Successful")
    name := r.FormValue("name")
    address := r.FormValue("address")
    fmt.Fprintf(w, "Name: %s\n", name)
    fmt.Fprintf(w, "address: %s\n", address)
}
