// Page 96
//
// test script:
//		curl -d "c=3" -d "url_long=111" -d "url_long=222" "http://127.0.0.1:9090/abc?a=1&b=2"

package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Scheme始终是空的
// 参考https://github.com/kataras/iris/blob/master/http.go中第327行的实现来间接获取

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
		fmt.Println(k, " = ", strings.Join(v, ", "))
	}
	fmt.Fprintf(w, "Hello from go!\n")
}

func main() {
	http.HandleFunc("/", sayHelloName)

	port := ":9090"
	log.Println("Listen to port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
