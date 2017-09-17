package main

import (
    "fmt"
    "net/http"
    "strings"
    "log"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    fmt.Println(r.Form)
    fmt.Println("path: ", r.URL.Path)
    fmt.Println(r.Form["url_long"])

    for k, v := range r.Form {
        fmt.Println("key: ", k)
        fmt.Println("val: ", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello x1ah")
}

func main() {
    http.HandleFunc("/", sayHelloName)
    err := http.ListenAndServe(":8888", nil)
    if err != nil {
        log.Fatal("Listenandserver: ", err)
    }
}
