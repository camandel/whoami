package main

import (
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

// ANSI colors
const (
	Red   = "\033[31m"
	Green = "\033[32m"
	Blue  = "\033[34m"
	Reset = "\033[0m"
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
		{{range .IPAddr}}
		<h1>{{.}}</h1>
		{{end}}
	</body>
</html>`

// DefaultColor is used to set the default background color
var DefaultColor = "red"
var hostname = ""
var ipAddr []string
var color = DefaultColor

// HTML output
func html(w http.ResponseWriter, r *http.Request) {

	data := struct {
		Title    string
		Color    string
		Hostname string
		IPAddr   []string
	}{
		Title:    "whoami",
		Color:    color,
		Hostname: hostname,
		IPAddr:   ipAddr,
	}

	fmt.Fprintf(os.Stdout, "I'm %s %s - %s app (html)\n", hostname, ipAddr, color)
	t, err := template.New("webpage").Parse(tpl)
	if err != nil {
		log.Fatal("Error loading html template")
	}
	err = t.Execute(w, data)
	if err != nil {
		log.Fatal("Error parsing html template")
	}
}

// simple output
func simple(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(os.Stdout, "I'm %s %s - %s app (simple)\n", hostname, ipAddr, color)
	fmt.Fprintf(w, "%s: I'm %s %s\n", strings.ToUpper(color), hostname, ipAddr)
}

// colored output
func colored(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(os.Stdout, "I'm %s %s - %s app (colored)\n", hostname, ipAddr, color)
	ansiColor := ""
	switch color {
	case "red":
		ansiColor = Red
	case "green":
		ansiColor = Green
	case "blue":
		ansiColor = Blue
	}
	fmt.Fprintf(w, fmt.Sprintf("%s%s: I'm %s %s%s\n", ansiColor, strings.ToUpper(color), hostname, ipAddr, Reset))
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	color = os.Getenv("COLOR")
	if color == "" {
		color = DefaultColor
	}

	hostname, _ = os.Hostname()

	// get IP addresses
	addrs, err := net.InterfaceAddrs()
	if err == nil {
		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					ipAddr = append(ipAddr, ipnet.IP.String())
				}
			}
		}
	}

	http.HandleFunc("/", html)
	http.HandleFunc("/simple", simple)
	http.HandleFunc("/colored", colored)

	log.Fatal(http.ListenAndServe(":"+port, nil))

	fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)
}
