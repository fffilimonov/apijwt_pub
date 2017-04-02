package routers

import (
	"../controllers"
	"github.com/gorilla/mux"
)

func SetSignupRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/api/signup", controllers.SignupController).Methods("POST")
	return router
}
