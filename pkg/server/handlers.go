package server

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
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

type timeOfWeekResp struct {
	Stats struct {
		BestHourOfTheWeek int
		DayOfWeek         int
		TimeOfDay         int
		ReactionCount     int
	}
	Data map[int]int
}

func (h handlers) TimeOfWeek(w http.ResponseWriter, r *http.Request) {

	al, err := db.ListArticles(h.state.DB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error"))
	}

	t0 := time.Now()
	rc := make(map[int]int)
	for _, a := range al {
		var doy int
		if int(a.PublishedAt.Weekday()) == 0 {
			doy = a.PublishedAt.Hour()
		} else {
			doy = a.PublishedAt.Hour() * int(a.PublishedAt.Weekday())
		}

		r, ok := rc[doy]
		if !ok {
			rc[doy] = a.PublicReactionsCount
		} else {
			rc[doy] = a.PublicReactionsCount + r
		}
	}
	log.Printf("iterating through all articles took %s", time.Since(t0))

	var bestTimeOfWeek, bestCount int
	for k, v := range rc {
		if v > bestCount {
			bestTimeOfWeek = k
			bestCount = v
		}
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", " ")
	enc.Encode(timeOfWeekResp{
		Stats: struct {
			BestHourOfTheWeek int
			DayOfWeek         int
			TimeOfDay         int
			ReactionCount     int
		}{
			BestHourOfTheWeek: bestTimeOfWeek,
			DayOfWeek:         bestTimeOfWeek / 24,
			TimeOfDay:         bestTimeOfWeek % 24,
			ReactionCount:     bestCount,
		},
		Data: rc,
	})
}

func (h handlers) HighestReactions(w http.ResponseWriter, r *http.Request) {
	al, err := db.ListArticles(h.state.DB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error"))
	}

	sort.Slice(al, func(i, j int) bool {
		return al[i].PublicReactionsCount > al[j].PublicReactionsCount
	})

	enc := json.NewEncoder(w)
	enc.SetIndent("", " ")
	enc.Encode(struct {
		Posts []db.Article
	}{
		Posts: al[:100],
	})
}

func (h handlers) HighestComments(w http.ResponseWriter, r *http.Request) {
	al, err := db.ListArticles(h.state.DB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error"))
	}

	sort.Slice(al, func(i, j int) bool {
		return al[i].CommentsCount > al[j].CommentsCount
	})

	enc := json.NewEncoder(w)
	enc.SetIndent("", " ")
	enc.Encode(struct {
		Posts []db.Article
	}{
		Posts: al[:100],
	})
}
