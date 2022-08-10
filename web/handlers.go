package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
)

type HomePage struct {
	Name string
}
func homeHandler(w http.ResponseWriter, r *http.Request, ps *httprouter.Params)  {
	p := &HomePage{Name: "interesting1113"}
	t, e := template.ParseFiles("./template/home.html")
	if e != nil {
		log.Printf("Parsing template home.html error: %s", e)
		return
	}

	t.Execute(w, p)
}
