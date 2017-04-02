package controllers

import (
	"../core/redis"
	"../services/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/context"
	"net/http"
)

func AddController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	requestStory := new(models.Scen)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestStory)

	user := context.Get(r, "Username").(string)
	fmt.Printf("Username: %v\n", user)

	name := requestStory.Name
	scen := requestStory.Text

	fmt.Printf("Name: %v\n", name)
	fmt.Printf("Scen: %v\n", scen)

	redisConn := redis.Pconnect(1)
	defer redisConn.Close()
	redisConn.HsetValue(user, name, scen)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"done\"}"))
}
