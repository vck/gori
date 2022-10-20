package main

import (
    "fmt"
    "net/http"
    "log"
    "io/ioutil"
    "github.com/gorilla/mux"
)


func Handler(writer http.ResponseWriter, req *http.Request) {
    body, _ := ioutil.ReadAll(req.Body)
    writer.Write([]byte(body))
}

func main() {
    r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    r.HandleFunc("/", Handler).Methods("POST")

    // Bind to a port and pass our router in
    log.Fatal(http.ListenAndServe(":9000", r))
}
