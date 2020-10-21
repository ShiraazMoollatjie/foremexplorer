package server

import (
	"encoding/json"
	"net/http"

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

func (h handlers) ComingSoonHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(comingSoonResponse{
		Message: "Busy being implemented",
	})
}
