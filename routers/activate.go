package routers

import (
	"../controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetActivateCaptchaRoutes(router *mux.Router) *mux.Router {
	router.Handle("/api/activate/{user}",
		negroni.New(
			negroni.HandlerFunc(controllers.ActivateCaptchaController),
		)).Methods("POST")

	return router
}
