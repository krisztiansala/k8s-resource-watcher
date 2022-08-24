package logging

import (
	"log"
	"net/http"
)

type StatusRecorder struct {
	http.ResponseWriter
	Status int
}

func (r *StatusRecorder) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

func WithLogging(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		recorder := &StatusRecorder{
			ResponseWriter: w,
			Status:         200,
		}
		handler.ServeHTTP(recorder, r)
		log.Printf("%s %s %d %s", r.RemoteAddr, r.Method, recorder.Status, r.URL)
	})
}
