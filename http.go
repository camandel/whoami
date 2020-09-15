package main

import (
	"fmt"
	"html/template"
	"log"
	"net"
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
		{{range .IPAddr}}
		<h1>{{.}}</h1>
		{{end}}
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

	// get IP addresses
	var ipAddr []string
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

	fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stdout, "I'm %s %s - %s app\n", hostname, ipAddr, color)
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
