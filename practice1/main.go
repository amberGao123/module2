package main

import (
	"io"
	"log"
	"net/http"
	"strings"
)

/*接收客户端 request，并将 request 中带的 header 写入 response header*/
func main() {

	http.HandleFunc("/header", func(w http.ResponseWriter, r *http.Request) {
		rheader := r.Header
		for k, v := range rheader {
			w.Header().Set(k, strings.Join(v, ", "))
		}
		io.WriteString(w, "headers ok!\n")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
