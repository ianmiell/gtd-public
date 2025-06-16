package main

import (
	"log"
	"net/http"
)

func searchTaskHandler(w http.ResponseWriter, r *http.Request) {
	var term string
	log.Println("In searchTaskHandler")
	data := PageData{Title: "Search Tasks"}
	if r.Method == http.MethodPost {
		term = r.FormValue("term")
	} else if r.Method == http.MethodGet {
		term = r.URL.Query().Get("term")
	}
	if term != "" {
		output, err := runGtdCommand("search", term)
		if err != nil {
			output += "\nError: " + err.Error()
		}
		data.Result = formatOutputWithAnchors(output)
		data.SearchTerm = term
	}
	log.Println("Rendering search.html")
	renderTemplate(w, "search.html", data)
}
