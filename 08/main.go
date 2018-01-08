package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template/*.html"))
}
func main() {
	http.HandleFunc("/", index)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":8081", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	p := "Eko Santoso"
	tpl.ExecuteTemplate(w, "index.html", p)
}
