package main

import (
	"net/http"
)

// Handlers for individual pages
func newTaskHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{Title: "New Task"}
	// TODO get taskId
	if r.Method == http.MethodPost {
		subject := r.FormValue("subject")
		output, err := runGtdCommand("nt", subject)
		if err != nil {
			output += "\nError: " + err.Error()
		}
		data.Result = getDashboardOutput()
		renderTemplate(w, "dashboard.html", data)
	} else {
		renderTemplate(w, "newtask.html", data)
	}
}
