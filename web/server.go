package web

import "net/http"

func Serve() {
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(_escFS(false))))
	//http.HandleFunc("/", srv.indexHandler)
	http.ListenAndServe(":9001", nil)
}
