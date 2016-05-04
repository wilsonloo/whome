// whome
package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var g_global_cfg map[string]string

func dump_request(r *http.Request) {

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

type Session struct {
	ModuleDir    string
	TopCommonTag string
}

var g_session Session

type TemplateEngine struct {
}

func (this *TemplateEngine) Dispatch(w http.ResponseWriter, r *http.Request, path string) {

	// parse top segment
	top_content := EchoView("top.html")

	// get html file
	body_content := EchoView(path)

	ret_content := top_content + body_content
	io.WriteString(w, ret_content)
}

var gTemplateEngine = &TemplateEngine{}

func myServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	dump_request(r)

	g_session.ModuleDir = "home"
	g_session.TopCommonTag = "id_top_common_tag_home"

	path := r.URL.Path
	if path == "/" || path == "/index/" || path == "/index.html" {
		g_session.ModuleDir = "/"
		g_session.TopCommonTag = "id_top_common_tag_home"
		path = "/" + g_session.ModuleDir + "/index.html"

	} else if path == "/about_me/" || path == "/about_me.html" || path == "/about_me/about_me.html" {
		g_session.ModuleDir = "about_me"
		g_session.TopCommonTag = "id_top_common_tag_about_me"
		path = "/" + g_session.ModuleDir + "/about_me.html"

	} else if path == "/lab/" || path == "/lab.html" || path == "/lab/lab.html" {
		g_session.ModuleDir = "lab"
		g_session.TopCommonTag = "id_top_common_tag_lab"
		path = "/" + g_session.ModuleDir + "/lab.html"

	} else {
		os.Exit(3333)
	}

	gTemplateEngine.Dispatch(w, r, path)
}

func init_app_cfg() {
	g_global_cfg = make(map[string]string)
	g_global_cfg["ROOT"] = "www"
	g_global_cfg["html_tmpl"] = g_global_cfg["ROOT"] + "/html_tmpl"

	// app full path
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)

	// app dir
	path_dir := filepath.Dir(path)

	g_global_cfg["APP_ROOT_DIR"] = path_dir

	g_session = Session{}
}

func main() {

	init_app_cfg()

	http.HandleFunc("/", myServeHTTP)

	www_root := g_global_cfg["APP_ROOT_DIR"] + "/" + g_global_cfg["ROOT"]
	fs_handler := http.FileServer(http.Dir(www_root + "/res/css/"))
	http.Handle("/css/", http.StripPrefix("/css/", fs_handler))

	fs_handler = http.FileServer(http.Dir(www_root + "/res/img/"))
	http.Handle("/images/", http.StripPrefix("/images/", fs_handler))

	fs_handler = http.FileServer(http.Dir(www_root + "/res/img/"))
	http.Handle("/favicon.ico", http.StripPrefix("/favicon.ico", fs_handler))

	// my_http_server := &MyHttpServer{}
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func EchoView(view_name string, args ...interface{}) string {

	// get html file
	real_path := g_global_cfg["html_tmpl"] + "/" + view_name

	// parse it
	t := template.New("new2")
	file_content, err := ioutil.ReadFile(real_path)
	CheckError(err)
	template.Must(t.Parse(string(file_content)))

	ret_buf := bytes.NewBufferString("")
	t.Execute(ret_buf, g_session)
	fmt.Println("the ret_buf size is ", ret_buf.Len())
	ret_string := ret_buf.String()

	// check if has nested templates
	pref_flag := "[[include"
	pref_flag_len := len(pref_flag)
	fmt.Println("=== for page: ", real_path)
	fmt.Println("content is ", len(ret_string))

	for {

		start_pos := strings.Index(ret_string, pref_flag)
		if start_pos == -1 {
			break
		}

		end_pos := strings.LastIndex(ret_string, "]]")
		if end_pos == -1 {
			break
		}

		// check if has target template as [[include xxx]]
		if start_pos+pref_flag_len >= end_pos {
			break
		}

		// get the nested template file
		nested_view := ret_string[start_pos+pref_flag_len : end_pos]
		fmt.Println("nested view is ", nested_view)
		nested_view = strings.Trim(nested_view, " \t\"")
		fmt.Println("nested view trimed is ", nested_view)
		nested_string := EchoView(nested_view)

		// replace old
		old := ret_string[start_pos : end_pos+2]
		fmt.Println("old is ", old)
		ret_string = strings.Replace(ret_string, old, nested_string, -1)
	}
	return ret_string
}

/*
func SayHello(w http.ResponseWriter, req *http.Request) {
	// info := Info{"个人网站", "克莱普斯", "http://www.sample.com/"}
	tmpl, _ := template.ParseFiles("tmpl.html")
	tmpl.Execute(w, nil)
}

func main2() {
	http.HandleFunc("/", SayHello)
	http.ListenAndServe(":9090", nil)
}
*/

func CheckError(err error) {
	if err != nil {
		log.Fatalln(err.Error())
		os.Exit(1)
	}
}
