package main

import (
    "net/http"
    "flag"
)

func main() {
    path := flag.String("path", "data", "PATH to serve")
    port := flag.String("port", "8000", "PORT to listen on")
    flag.Parse()
    panic(http.ListenAndServe(":" + *port, http.FileServer(http.Dir(*path))))
}