package handler

import (
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
	w.Write([]byte("hello"))
}

func Bye(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/bye.html")
}
