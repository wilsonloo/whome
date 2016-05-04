// whome
package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var g_global_cfg map[string]string

func dump_request(r *http.Request) {
	r.ParseForm()

	fmt.Println("request map: ", r.Form)
	fmt.Println("path: ", r.URL.Path)
	fmt.Println("schema: ", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, ";"))
	}
}

func tmpl_writer(w io.Writer, format string, val ...interface{}) {
	var output string
	output = "sdfsdfsdf"

	w.Write([]byte(output))
}

type Person struct {
	Age  int
	Name string
}

type TemplateEngine struct {
}

func (this *TemplateEngine) Dispatch(w http.ResponseWriter, r *http.Request, path string) {
	// todo get html file
	log.Println("casting real path...")
	real_path := g_global_cfg["html_templ"] + path

	// todo parse it
	log.Println("parsing template...")
	t, err := template.ParseFiles(real_path)
	CheckError(err)

	// todo write to client
	log.Println("writing to client...")
	p := Person{
		Age:  3333,
		Name: "namesdfsd",
	}
	err = t.Execute(w, p)
	CheckError(err)
}

var g_template_engine = &TemplateEngine{}

func myServeHTTP(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
	if path == "/" {
		path = "/index.html"
	}

	g_template_engine.Dispatch(w, r, path)
}

func main() {
	g_global_cfg = make(map[string]string)
	g_global_cfg["ROOT"] = "www"
	g_global_cfg["html_templ"] = g_global_cfg["ROOT"] + "/html_tmpl"

	http.HandleFunc("/", myServeHTTP)

	// file_server := http.StripPrefix("/favicon.ico/", http.FileServer(http.Dir("www/res/image/")))
	// http.Handle("/favicon.ico/", file_server)

	// my_http_server := &MyHttpServer{}
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func SayHello(w http.ResponseWriter, req *http.Request) {
	// info := Info{"个人网站", "克莱普斯", "http://www.sample.com/"}
	tmpl, _ := template.ParseFiles("tmpl.html")
	tmpl.Execute(w, nil)
}

func main2() {
	http.HandleFunc("/", SayHello)
	http.ListenAndServe(":9090", nil)
}

func CheckError(err error) {
	if err != nil {
		log.Fatalln(err.Error())
		os.Exit(1)
	}
}
