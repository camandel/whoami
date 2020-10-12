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

func TestSimpleOutput(t *testing.T) {
	t.Run("simple output", func(t *testing.T) {
		for _, color = range colors {
			request, _ := http.NewRequest(http.MethodGet, "/simple", nil)
			response := httptest.NewRecorder()

			simple(response, request)

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
			request, _ := http.NewRequest(http.MethodGet, "/", nil)
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
