// whome
package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

/*
* router
 */
type MyHttpServer struct {
}

func (this *MyHttpServer) Serve(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println("request map: ", r.Form)
	fmt.Println("path: ", r.URL.Path)
	fmt.Println("schema: ", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, ";"))
	}

	fmt.Fprintf(w, "hello gerryyange\n")
}

func main() {
	fmt.Println("Hello World!")

	my_http_server := &MyHttpServer{}

	err := http.ListenAndServe(":9090", my_http_server)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
