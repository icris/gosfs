package main

import (
    "net/http"
    "flag"
    "fmt"
)

func main() {
    path := flag.String("path", ".", "PATH to serve")
    port := flag.String("port", "8000", "PORT to listen on")
    flag.Parse()
    fmt.Printf("listening on port %v...\n", *port)
    panic(http.ListenAndServe(":" + *port, http.FileServer(http.Dir(*path))))
}