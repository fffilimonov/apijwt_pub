package controllers

import (
	"../core/redis"
	"../hyper"
	"../services/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/context"
	"github.com/pborman/uuid"
	"net/http"
	"strconv"
)

func RunController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	requestStory := new(models.Run)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestStory)

	user := context.Get(r, "Username").(string)
	scen := requestStory.Scen
	browser := requestStory.Browser
	UUID := uuid.New()

	fmt.Printf("Username: %v\n", user)
	fmt.Printf("Scen: %v\n", scen)
	fmt.Printf("Browser: %v\n", browser)
	fmt.Printf("Res: %v\n", UUID)

	redisConn := redis.Pconnect(2)
	defer redisConn.Close()
	hashMap, _ := redisConn.HgetAll(user)
	currentValueStr := hashMap["Seconds"]
	currentValue, _ := strconv.Atoi(currentValueStr)
	fmt.Printf("Seconds: %v\n", currentValue)

	if currentValue > 0 {
		redisConn3 := redis.Pconnect(3)
		defer redisConn3.Close()
		redisConn3.HappendValue(user, UUID, "{\"Staring\": true}")
		go hyper.StartBohrium(user, scen, UUID, browser)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(UUID))
	} else {
		w.WriteHeader(http.StatusPaymentRequired)
	}
}
