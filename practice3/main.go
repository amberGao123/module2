package main

import (
	"fmt"
	"log"
	"net/http"
)

/*Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出*/
func main() {

	http.HandleFunc("/requestInfo", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request url %s, host %s, status: %d", r.URL.Path, r.Host, http.StatusOK)
		fmt.Fprintf(w, "request url %s, host %s, status: %d", r.URL.Path, r.Host, http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(":8082", nil))
}
