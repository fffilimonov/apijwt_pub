package routers

import (
	"../controllers"
	"../core/authentication"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetResultRoutes(router *mux.Router) *mux.Router {
	router.Handle("/api/result/{id}",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.ResultController),
		)).Methods("Get")

	return router
}

func SetResultsRoutes(router *mux.Router) *mux.Router {
	router.Handle("/api/results",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.ResultsController),
		)).Methods("Get")

	return router
}
