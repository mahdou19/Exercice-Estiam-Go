package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		routeInfo := mux.CurrentRoute(r)
		var route string
		if routeInfo != nil {
			route, _ = routeInfo.GetPathTemplate()
		}

		next.ServeHTTP(w, r)

		log.Printf("[%s] %s %s %s", time.Since(start), r.Method, route, r.URL.Path)
	})
}
