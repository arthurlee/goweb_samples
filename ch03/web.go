package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("------------------------")
	fmt.Println("form:", r.Form)
	fmt.Println("method:", r.Method)
	fmt.Println("url:", r.URL)
	fmt.Println("host:", r.Host)
	fmt.Println("path:", r.URL.Path)
	fmt.Println("scheme:", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println(k, " = ", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello from go!\n")
}

func main() {
	http.HandleFunc("/", sayHelloName)

	port := ":9090"
	log.Println("Listen to port ", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
