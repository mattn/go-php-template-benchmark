package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
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
	http.ListenAndServe(":8081", nil)
}
