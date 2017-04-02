package routers

import (
	"../controllers"
	"../core/authentication"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetDashboardRoutes(router *mux.Router) *mux.Router {
	router.Handle("/api/dashboard",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.DashboardController),
		)).Methods("GET")

	return router
}
