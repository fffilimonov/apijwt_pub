package controllers

import (
	"../core/redis"
	"../services"
	"../services/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func SignupController(w http.ResponseWriter, r *http.Request) {
	requestUser := new(models.User)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)

	fmt.Printf("Username: %v\n", requestUser.Username)

	redisConn := redis.Pconnect(2)
	defer redisConn.Close()
	hashMap, reer := redisConn.HgetAll(requestUser.Username)
	if reer != nil {
		panic(reer)
	}
	if _, ok := hashMap["Password"]; ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("exist"))
	} else {
		redisConn.HsetValue(requestUser.Username, "Password", requestUser.GetPassword())
		redisConn.HsetValue(requestUser.Username, "Active", "0")
		redisConn.HsetValue(requestUser.Username, "Seconds", "300")
		responseStatus, resp := services.SendActivation(requestUser)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(responseStatus)
		w.Write(resp)
	}
}
