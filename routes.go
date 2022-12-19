/*
	Available routes in project
*/

package go_laravel_package

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (glp *GoLaravelPackage) routes() http.Handler {
	//! do not change the middleware order
	mux := chi.NewRouter()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)

	if glp.Debug {
		mux.Use(middleware.Logger)
	}
	mux.Use(middleware.Recoverer)

	//#region inline routers
	mux.Get("/glp", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "glp, go laravel package")
	})
	//#endregion inline routers

	return mux
}
