package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
)

var colors = []string{"red", "green", "blue"}

func TestTextOutput(t *testing.T) {
	t.Run("text output", func(t *testing.T) {
		for _, color = range colors {
			request, _ := http.NewRequest(http.MethodGet, "/text", nil)
			response := httptest.NewRecorder()

			text(response, request)

			got := response.Body.String()
			want := strings.ToUpper(color)

			ok, _ := regexp.MatchString(fmt.Sprintf("%s:.*", want), got)
			if !ok {
				t.Errorf("got %q, want %q", got, want)
			}
		}
	})
}

func TestColoredOutput(t *testing.T) {
	t.Run("colored output", func(t *testing.T) {
		for _, color = range colors {
			request, _ := http.NewRequest(http.MethodGet, "/colored", nil)
			response := httptest.NewRecorder()

			colored(response, request)

			got := response.Body.String()
			want := strings.ToUpper(color)

			ok, _ := regexp.MatchString(fmt.Sprintf("%s:.*", want), got)
			if !ok {
				t.Errorf("got %q, want %q", got, want)
			}
		}
	})
}

func TestHtmlOutput(t *testing.T) {
	t.Run("html output", func(t *testing.T) {
		for _, color = range colors {
			request, _ := http.NewRequest(http.MethodGet, "/html", nil)
			response := httptest.NewRecorder()

			html(response, request)

			got := response.Body.String()
			want := color

			ok, _ := regexp.MatchString(fmt.Sprintf("background-color: %s", want), got)
			if !ok {
				t.Errorf("got %q, want %q", got, want)
			}
		}
	})
}

func TestAutoOutput(t *testing.T) {
	tests := []struct {
		browser string
		color   string
		want    string
	}{
		{
			browser: "curl/7.69.1",
			color:   "green",
			want:    "GREEN:",
		},
		{
			browser: "curl/6.0.0",
			color:   "blue",
			want:    "BLUE:",
		},
		{
			browser: "Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:81.0) Gecko/20100101 Firefox/81.0",
			color:   "red",
			want:    "background-color: red",
		},
		{
			browser: "Unknown_Browser/1.0",
			color:   "green",
			want:    "background-color: green",
		},
	}

	t.Run("auto output", func(t *testing.T) {
		for _, test := range tests {
			color = test.color
			request, _ := http.NewRequest(http.MethodGet, "/", nil)
			request.Header.Set("User-Agent", test.browser)
			response := httptest.NewRecorder()

			auto(response, request)

			got := response.Body.String()
			want := test.want

			ok, _ := regexp.MatchString(test.want, got)
			if !ok {
				t.Errorf("got %q, want %q", got, want)
			}
		}
	})
}
