package main

import (
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

func main() {
    router := mux.NewRouter()
 
    router.HandleFunc("/", ConvertDate)

    log.Fatal(http.ListenAndServe(":8000", router))
}