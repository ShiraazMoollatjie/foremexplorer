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
	r.HandleFunc("/analytics/posts/timeofday", h.TimeOfDay).Methods("GET")
	r.HandleFunc("/analytics/posts/timeofweek", h.TimeOfWeek).Methods("GET")
	r.HandleFunc("/analytics/highestcomments", h.ComingSoon).Methods("GET")
	r.HandleFunc("/analytics/highestreactions", h.HighestReactions).Methods("GET")

	log.Println("starting server")
	log.Fatal(http.ListenAndServe(":8000", r))
}
