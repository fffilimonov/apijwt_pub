package routers

import (
	"../controllers"
	"github.com/gorilla/mux"
)

func SetLoginRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/api/login", controllers.LoginController).Methods("POST")
	return router
}
