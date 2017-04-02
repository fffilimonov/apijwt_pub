package main

import (
	"./core/redis"
	"./routers"
	"./settings"
	"github.com/codegangsta/negroni"
	"net/http"
)

//go:generate go run scripts/embed.go

func main() {
	settings.Init()

	redis.Pinit()
	defer redis.Pclose()

	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":5000", n)
}
