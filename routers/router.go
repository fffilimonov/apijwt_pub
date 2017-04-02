package routers

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.KeepContext = true
	router = SetAddRoutes(router)
	router = SetRunRoutes(router)
	router = SetDashboardRoutes(router)
	router = SetLoginRoutes(router)
	router = SetSignupRoutes(router)
	router = SetActivateCaptchaRoutes(router)
	router = SetResultRoutes(router)
	router = SetResultsRoutes(router)
	return router
}
