<!--
In Go templates, checking {{if .TaskId}} works if .TaskId is a non-zero value (i.e., not empty, nil, false, or 0). However, if .TaskId might be completely missing from the context (not just empty), a safer approach is:

Solution:
go
Copy
Edit
{ {if isset . "TaskId"}}
Unfortunately, Go templates don't have a built-in isset function, so you need to define it in your Go code like this:

go
Copy
Edit
import "text/template"

// Helper function to check if a key exists in a map
func isset(m map[string]interface{}, key string) bool {
    _, exists := m[key]
        return exists
        }

        // Register the function
        tmpl := template.New("example").Funcs(template.FuncMap{
            "isset": isset,
            })
            Then you can use:

            html
            Copy
            Edit
            { {if isset . "TaskId"}}
            This ensures .TaskId is present in the context before checking its value.
            -->
