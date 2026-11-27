package main

import (
	"net/http"
)

// func HAdmit(w http.ResponseWriter, request *http.Request) {
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	w.Write([]byte())

// }

// func mh(w http.ResponseWriter, request *http.Request) {
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	w.Write([]byte(formadmint))

// }

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	//	http.HandleFunc("/admit", HAdmit)
	http.ListenAndServe(":8080", nil)
}
