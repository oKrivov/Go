package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Println("started", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		log.Println("finished", time.Since(start))
	})
}

func hello(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)
	fmt.Fprintln(w, "hello")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	handler := loggingMiddleware(mux)

	fmt.Println("server started :8080")
	http.ListenAndServe(":8080", handler)
}
