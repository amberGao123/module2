package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	DEFAULT_VERSION = "demo"
)

/*读取当前系统的环境变量中的 VERSION 配置，并写入 response header*/
func main() {

	http.HandleFunc("/header_version", func(w http.ResponseWriter, r *http.Request) {
		rheader := r.Header
		for k, v := range rheader {
			w.Header().Set(k, strings.Join(v, ", "))
		}

		version := os.Getenv("VERSION")
		if version == "" {
			version = DEFAULT_VERSION
		}
		w.Header().Set("VERSION", version)

		io.WriteString(w, "headers with version ok!\n")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
