package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var count int = 1

func main() {
	fmt.Printf("====== This is http server ======\n")
	flag.Set("v", "4")
	// glog.V(2).Info("starting http server...")
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/getenv", Getenv)
	http.HandleFunc("/getstatuscode", Getstatuscode)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("------ This is roothandler[%d] ------\n", count)
	user := r.URL.Query().Get("user")
	if user != "" {
		io.WriteString(w, fmt.Sprintf("Hello [%s]\n", user))
	} else {
		io.WriteString(w, "Hello [Stranger]\n")
	}
	io.WriteString(w, "========== Details of the http request header:============\n")
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s = %s\n", k, v))
	}
	count++
}

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("------ This is healthz[%d] ------\n", count)
	io.WriteString(w, "200\n")
	count++
}

func Getenv(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("------ This is getenv[%d] ------\n", count)
	var goversion string
	goversion = os.Getenv("GOVERSION")
	if goversion != "" {
		io.WriteString(w, fmt.Sprintf("goversion is: %s\n", goversion))
	} else {
		io.WriteString(w, "version is nil\n")
	}
	fmt.Println(goversion)
	count++
}

func Getstatuscode(w http.ResponseWriter, r *http.Request) {
	rsp, err := http.Get("http://127.0.0.1:8090")
	if err != nil {
		fmt.Println("httpget rsp error is: ", err)
		return
	}
	io.WriteString(w, fmt.Sprintf("getstatuscode is: %d", rsp.StatusCode))
	fmt.Println("statuscode is: %d", rsp.StatusCode)
}
