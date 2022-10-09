package main

import (
	"log"
	"net/http"
)

/*当访问 localhost/healthz 时，应返回 200*/
func main() {

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("200"))
	})

	log.Fatal(http.ListenAndServe(":8083", nil))
}
