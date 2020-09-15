package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// HTML template
const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
        <title>{{.Title}}</title>
        <style>
        body {
            background-color: {{.Color}};
            color: white;
            padding-top: 50px;
            text-align: center;
            font-family: sans-serif, serif;
        }
        </style>
	</head>
	<body>
		<h1>I'm {{.Hostname}}</h1>
	</body>
</html>`

// DefaultColor is used to set the default background color
var DefaultColor = "red"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	color := os.Getenv("COLOR")
	if color == "" {
		color = DefaultColor
	}

	hostname, _ := os.Hostname()

	data := struct {
		Title    string
		Color    string
		Hostname string
	}{
		Title:    "whoami",
		Color:    color,
		Hostname: hostname,
	}

	fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stdout, "I'm %s (%s app)\n", hostname, color)
		t, err := template.New("webpage").Parse(tpl)
		if err != nil {
			log.Fatal("Error loading html template")
		}
		err = t.Execute(w, data)
		if err != nil {
			log.Fatal("Error parsing html template")
		}
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
