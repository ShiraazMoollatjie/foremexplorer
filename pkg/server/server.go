package server

import (
	"log"
	"net/http"

	"github.com/ShiraazMoollatjie/foremexplorer/pkg/state"
	"github.com/gorilla/mux"
)

func ServeHttp(s *state.State) {
	r := mux.NewRouter()
	h := NewHandlers(s)

	r.HandleFunc("/", h.ComingSoonHandler).Methods("GET")
	r.HandleFunc("/analytics/posts/dayofyear", h.ComingSoonHandler).Methods("GET")
	r.HandleFunc("/analytics/posts/dayofyear", h.ComingSoonHandler).Methods("GET")
	r.HandleFunc("/analytics/posts/timeofday", h.ComingSoonHandler).Methods("GET")
	r.HandleFunc("/analytics/posts/timeofweek", h.ComingSoonHandler).Methods("GET")
	r.HandleFunc("/analytics/tags/highestcomments", h.ComingSoonHandler).Methods("GET")
	r.HandleFunc("/analytics/tags/highestreactions", h.ComingSoonHandler).Methods("GET")

	log.Println("starting server")
	log.Fatal(http.ListenAndServe(":8000", r))
}
