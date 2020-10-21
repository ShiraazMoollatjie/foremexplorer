package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/ShiraazMoollatjie/foremexplorer/pkg/db"
	"github.com/ShiraazMoollatjie/foremexplorer/pkg/state"
)

type handlers struct {
	state *state.State
}

func NewHandlers(s *state.State) handlers {
	return handlers{
		state: s,
	}
}

type comingSoonResponse struct {
	Message string
}

func (h handlers) ComingSoon(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(comingSoonResponse{
		Message: "Busy being implemented",
	})
}

type dayOfYearResp struct {
	Stats struct {
		BestDayOfYear int
		ReactionCount int
	}
	Data map[int]int
}

func (h handlers) DayOfYear(w http.ResponseWriter, r *http.Request) {

	al, err := db.ListArticles(h.state.DB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error"))
	}

	t0 := time.Now()
	rc := make(map[int]int)
	for _, a := range al {
		doy := a.PublishedAt.YearDay()
		r, ok := rc[doy]
		if !ok {
			rc[doy] = a.PublicReactionsCount
		} else {
			rc[doy] = a.PublicReactionsCount + r
		}
	}
	log.Printf("iterating through all articles took %s", time.Since(t0))

	var bestDay, bestCount int
	for k, v := range rc {
		if v > bestCount {
			bestDay = k
			bestCount = v
		}
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", " ")
	enc.Encode(dayOfYearResp{
		Stats: struct {
			BestDayOfYear int
			ReactionCount int
		}{
			BestDayOfYear: bestDay,
			ReactionCount: bestCount,
		},
		Data: rc,
	})
}

type timeOfDayResp struct {
	Stats struct {
		BestHourOfDay int
		ReactionCount int
	}
	Data map[int]int
}

func (h handlers) TimeOfDay(w http.ResponseWriter, r *http.Request) {

	al, err := db.ListArticles(h.state.DB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error"))
	}

	t0 := time.Now()
	rc := make(map[int]int)
	for _, a := range al {
		doy := a.PublishedAt.Hour()
		r, ok := rc[doy]
		if !ok {
			rc[doy] = a.PublicReactionsCount
		} else {
			rc[doy] = a.PublicReactionsCount + r
		}
	}
	log.Printf("iterating through all articles took %s", time.Since(t0))

	var bestHour, bestCount int
	for k, v := range rc {
		if v > bestCount {
			bestHour = k
			bestCount = v
		}
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", " ")
	enc.Encode(timeOfDayResp{
		Stats: struct {
			BestHourOfDay int
			ReactionCount int
		}{
			BestHourOfDay: bestHour,
			ReactionCount: bestCount,
		},
		Data: rc,
	})
}
