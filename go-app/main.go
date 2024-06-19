//go:generate go get -u github.com/valyala/quicktemplate/qtc
//go:generate qtc -dir=templates
package main

import (
	"encoding/json"
	"fmt"
	gojson "github.com/goccy/go-json"
	"github.com/libp2p/go-reuseport"
	_ "github.com/valyala/quicktemplate"
	"go-app/templates"
	"html/template"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

func main() {
	tpl := template.Must(template.New("").ParseFiles("app.tpl"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		b, err := os.ReadFile("input.json")
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var data any
		err = json.Unmarshal(b, &data)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, "<html><body><ul>")
		err = tpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, "</ul></body></html>")
	})
	http.HandleFunc("/qtpl", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		b, err := os.ReadFile("input.json")
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var data []string
		err = json.Unmarshal(b, &data)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, "<html><body><ul>")
		fmt.Fprintln(w, templates.HTML(data))
		fmt.Fprintln(w, "</ul></body></html>")
	})
	http.HandleFunc("/qtpl_gojson", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		b, err := os.ReadFile("input.json")
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var data []string
		err = gojson.Unmarshal(b, &data)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, "<html><body><ul>")
		fmt.Fprintln(w, templates.HTML(data))
		fmt.Fprintln(w, "</ul></body></html>")
	})

	listener, err := reuseport.Listen("tcp4", ":8080")
	if err != nil {
		log.Panic(err)
	}
	http.Serve(listener, nil)
}
