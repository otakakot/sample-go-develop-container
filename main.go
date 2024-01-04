package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(Hello(r.UserAgent())))
	})

	slog.Info("http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}

func Hello(s string) string {
	return fmt.Sprintf("Hello, %s", s)
}
