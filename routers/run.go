package routers

import (
	"../controllers"
	"../core/authentication"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetRunRoutes(router *mux.Router) *mux.Router {
	router.Handle("/api/run",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.RunController),
		)).Methods("Post")

	return router
}
