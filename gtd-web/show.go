package main

import (
	"log"
	"net/http"
)

// Handlers for individual pages
func showTaskHandler(w http.ResponseWriter, r *http.Request) {
	var taskid string
	var output string
	var status string
	var priority string
    var err error
	log.Println("In showTaskHandler")
	data := PageData{Title: "Show Task"}
	if r.Method == http.MethodPost {
		taskid = r.FormValue("taskid")
		status = r.FormValue("status")
		priority = r.FormValue("priority")
	} else if r.Method == http.MethodGet {
		taskid = r.URL.Query().Get("taskid")
		status = r.URL.Query().Get("status")
		priority = r.URL.Query().Get("priority")
	}
	if status != "" {
		// TODO: Check values are acceptable
		runGtdCommand("move-status", status, taskid)
	}
	if priority != "" {
		// TODO: Check values are acceptable
		runGtdCommand("move-priority", priority, taskid)
	}
	if taskid != "" {
		output, err = runGtdCommand("show", taskid)
		if err != nil {
			output += "\nError: " + err.Error()
		}
		data.Result = formatOutputWithAnchors(output)
		data.TaskId = taskid
	}
	// TODO: assert taskid is a numeric
	// TODO: determine current priority and status (and then don't show those on the header)
	log.Println("Rendering show.html")
	renderTemplate(w, "show.html", data)
}
