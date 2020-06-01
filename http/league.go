package http

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	football "github.com/sebasm/footballgo"
)

type LeagueHandler struct {
	*chi.Mux
	LeagueService football.LeagueService
}

func NewLeagueHanlder() *LeagueHandler {
	h := &LeagueHandler{
		Mux: chi.NewRouter(),
	}
	h.Get("/{leagueID}", h.getLeague)
	return h
}

func (h *LeagueHandler) getLeague(w http.ResponseWriter, r *http.Request) {
	leagueID, err := strconv.Atoi(chi.URLParam(r, "leagueID"))
	asd, err := h.LeagueService.League(leagueID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, "League not found.")
		return
	}

	render.JSON(w, r, asd)

}
