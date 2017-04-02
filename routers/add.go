package routers

import (
	"../controllers"
	"../core/authentication"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetAddRoutes(router *mux.Router) *mux.Router {
	router.Handle("/api/add",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.AddController),
		)).Methods("Post")

	return router
}
