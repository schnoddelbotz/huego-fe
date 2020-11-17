package web

import "net/http"

func Serve() error {
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(_escFS(false))))
	//http.HandleFunc("/", srv.indexHandler)
	return http.ListenAndServe(":9001", nil)
}
