/**
*	NAME: Aaron Meek
*	DATE: 2022 - 08 - 15
*
*	This contains the routing logic
 */

package main

import (
	"net/http"

	"github.com/urhumantoast/NewPersonalWebsite/pkg/config"
	"github.com/urhumantoast/NewPersonalWebsite/pkg/handlers"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/contact-complete", handlers.Repo.ContactComplete)
	mux.Get("/proj-main", handlers.Repo.ProjMain)
	mux.Get("/proj-app", handlers.Repo.ProjApp)
	mux.Get("/proj-emb", handlers.Repo.ProjEmb)
	mux.Get("/proj-elc", handlers.Repo.ProjElc)
	mux.Get("/placeholder", handlers.Repo.Placeholder)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
