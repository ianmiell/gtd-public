package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

func dashboardTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("In dashboardTaskHandler")
	data := PageData{Title: "Show Dashboard"}
	data.Result = getDashboardOutput()
	//log.Println(data.Result)
	log.Println("Rendering dashboard.html")
	renderTemplate(w, "dashboard.html", data)
}

func pullTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("In pullTaskHandler")
	output, err := runGtdCommand("pull")
	if err != nil {
		output += "\nError: " + err.Error()
		data := PageData{Title: "Show Dashboard"}
		data.Result = formatOutputWithAnchors(output)
		log.Println("Rendering dashboard.html after error in pullTaskHandler")
		renderTemplate(w, "dashboard.html", data)
	}
	log.Println(output)
	dashboardTaskHandler(w,r)
}

func pushTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("In pushTaskHandler")
	output, err := runGtdCommand("push")
	if err != nil {
		output += "\nError: " + err.Error()
		data := PageData{Title: "Show Dashboard"}
		data.Result = formatOutputWithAnchors(output)
		log.Println("Rendering dashboard.html after error in pushTaskHandler")
		renderTemplate(w, "dashboard.html", data)
	}
	log.Println(output)
	dashboardTaskHandler(w,r)
}

// Data gatherer components
func getDashboardOutput() template.HTML {
	output, err := runGtdCommand("dashboard")
	if err != nil {
		output += "\nError: " + err.Error()
	}
	return formatOutputWithAnchors(output)
}

// Service functions
func formatOutputWithAnchors(output string) template.HTML {
	var builder strings.Builder
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			builder.WriteString("<br>")
			continue
		}
		// Detect if line starts with a number
		fields := strings.Fields(trimmed)
		if len(fields) > 0 && isNumeric(fields[0]) {
			taskID := fields[0]
			anchor := `<a href="/gtd/show?taskid=` + taskID + `">` + template.HTMLEscapeString(taskID) + `</a>`
			// Replace taskID in the original line with anchor
			escapedLine := template.HTMLEscapeString(line)
			linkedLine := strings.Replace(escapedLine, taskID, anchor, 1)
			builder.WriteString(linkedLine)
		} else {
			builder.WriteString(template.HTMLEscapeString(line))
		}
		builder.WriteString("<br>")
	}
	//log.Println(builder.String())
	return template.HTML(builder.String()) // safe because we escape non-anchor content
}

func isNumeric(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

