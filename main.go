package main

import (
	"net/http"
)

// START - Handle Request Using Struct and ServeHTTP Method
type fooHandler struct {
	Message string
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(f.Message))
}
// END - Handle Request Using Struct and ServeHTTP Method


// START - Handle Request Using Function
func barHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Grr!"))
}


func main() {
	http.Handle("/foo", &fooHandler{Message: "Meow!"})
	http.HandleFunc("/bar", barHandler)

	http.ListenAndServe(":5000", nil)
}