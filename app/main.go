package main

import (
  "fmt"
  "net/http"
  "reflect"
  "runtime"
)

func hello (w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello")
}

func log (h http.HandlerFunc) http.HandlerFunc {
  return func (w http.ResponseWriter, r *http.Request) {
    name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
    fmt.Println("Hanler function called - " + name)
    h(w, r)
  }
}

func main() {
  server := http.Server{
	  Addr: ":1323",
  }

  http.Handle("/hello", log(hello))

  server.ListenAndServe()
}