package http

import (
	"net/http"
	"strconv"

	"github.com/sebasm/footballgo/api"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	football "github.com/sebasm/footballgo"
)

type LeagueHandler struct {
	*chi.Mux
	LeagueService football.LeagueService
	LeagueAPI     football.LeagueAPI
}

func NewLeagueHanlder() *LeagueHandler {
	h := &LeagueHandler{
		Mux: chi.NewRouter(),
	}
	h.Get("/{leagueID}", h.getLeague)
	h.Get("/import-league/{leagueID}", h.importLeague)
	return h
}

func (h *LeagueHandler) getLeague(w http.ResponseWriter, r *http.Request) {
	_, err := strconv.Atoi(chi.URLParam(r, "leagueID"))
	league, err := api.GetLeague("CL")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, "League not found.")
		return
	}
	asd, err := h.LeagueService.ImportLeague(league)

	render.JSON(w, r, asd)

}

func (h *LeagueHandler) importLeague(w http.ResponseWriter, r *http.Request) {
	chi.URLParam(r, "leagueID")
}
