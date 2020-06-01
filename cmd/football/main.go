package main

import (
	"log"
	"net/http"

	httpm "github.com/sebasm/footballgo/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/sebasm/footballgo/postgres"
)

// func Routes() *chi.Mux {

// 	router := chi.NewRouter()

// 	router.Use(
// 		render.SetContentType(render.ContentTypeJSON),
// 		middleware.Logger,
// 		middleware.RedirectSlashes,
// 		middleware.Recoverer,
// 	)

// 	router.Route("/v1", func(r chi.Router) {
// 		r.Mount("/api/book", httpm.BookRoutes())
// 		// r.Mount("/api/user", routes.UserRoutes())
// 	})

// 	return router

// }

func main() {

	db, err := postgres.Open("")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	leagueService := &postgres.LeagueService{DB: db}

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
