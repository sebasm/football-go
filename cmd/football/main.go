package main

import (
	"log"
	"net/http"

	football "github.com/sebasm/footballgo"
	httpm "github.com/sebasm/footballgo/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/sebasm/footballgo/postgres"
)

func main() {

	db, err := postgres.Open("")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var leagueService football.LeagueService = &postgres.LeagueService{DB: db}

	h := httpm.NewLeagueHanlder()
	h.LeagueService = leagueService

	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api/league", h)
		// r.Mount("/api/user", routes.UserRoutes())
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}
