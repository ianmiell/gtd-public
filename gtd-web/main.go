package main

import (

	"html/template"
	"log"
	"net/http"
	"os/exec"
	"os"
	"strings"
	"runtime"

	"github.com/gorilla/mux"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

type PageData struct {
	Title  string
	Result template.HTML
	TaskId string
	SearchTerm string
}

// Optional: define an error for unsupported platforms
type UnsupportedPlatformError struct {
	OS string
}

func (e *UnsupportedPlatformError) Error() string {
	return "unsupported operating system: " + e.OS
}

func runGtdCommand(args ...string) (string, error) {
	var gtdPath string
	switch runtime.GOOS {
		case "darwin":
			gtdPath = "../bin/gtd"
		case "linux":
			gtdPath = "../bin/gtd"
		default:
			return "", &UnsupportedPlatformError{OS: runtime.GOOS}
	}
	cmd := exec.Command(gtdPath, args...)
	cmd.Env = append(os.Environ(), "GTD_DISPLAY_CONTEXT=web")
	out, err := cmd.CombinedOutput()
	return strings.TrimSpace(string(out)), err
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	log.Println("Rendering template: " + tmpl)
	if err := templates.ExecuteTemplate(w, tmpl, data); err != nil {
		log.Println("Error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Main function
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/gtd/newtask", newTaskHandler).Methods("GET", "POST")
	r.HandleFunc("/gtd/show", showTaskHandler).Methods("GET","POST")
	r.HandleFunc("/gtd/dashboard", dashboardTaskHandler).Methods("GET", "POST")
	r.HandleFunc("/gtd/search", searchTaskHandler).Methods("GET", "POST")
	r.HandleFunc("/gtd/pull", pullTaskHandler).Methods("GET")
	r.HandleFunc("/gtd/push", pushTaskHandler).Methods("GET")

	log.Println("Starting server at :9090")
	log.Fatal(http.ListenAndServe(":9090", r))
}
