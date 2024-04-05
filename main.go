package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hello Argo!!!!</h1> <p>tudo certo</p>"))
	})
	http.ListenAndServe(":8080", nil)
}
