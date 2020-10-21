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

	r.HandleFunc("/", h.ComingSoon).Methods("GET")
	r.HandleFunc("/analytics/posts/dayofyear", h.DayOfYear).Methods("GET")
	r.HandleFunc("/analytics/posts/timeofday", h.ComingSoon).Methods("GET")
	r.HandleFunc("/analytics/posts/timeofweek", h.ComingSoon).Methods("GET")
	r.HandleFunc("/analytics/tags/highestcomments", h.ComingSoon).Methods("GET")
	r.HandleFunc("/analytics/tags/highestreactions", h.ComingSoon).Methods("GET")

	log.Println("starting server")
	log.Fatal(http.ListenAndServe(":8000", r))
}
