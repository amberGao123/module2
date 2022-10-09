package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/demo", func(w http.ResponseWriter, r *http.Request) {
		demo := r.Header.Get("Demo")
		w.Header().Set("demo", demo)

		version := os.Getenv("VERSION")
		w.Header().Set("VERSION", version)

		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			fmt.Fprintf(w, "userip: %q is not IP:port", r.RemoteAddr)
		}
		fmt.Fprintf(w, "userip: %v, statusCode %v", ip, http.StatusOK)
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("200"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
