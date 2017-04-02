package controllers

import (
	"../core/redis"
	"../services/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"net/http"
)

func ResultController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Printf("result ID: %v\n", id)

	redisConn := redis.Pconnect(0)
	defer redisConn.Close()
	rres, reer := redisConn.GetValue(id)
	if reer != nil {
		if reer == redis.ErrNil {
			w.WriteHeader(http.StatusGone)
		} else {
			panic(reer)
		}
	} else {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("[" + rres + "]"))
	}
}

func ResultsController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	user := context.Get(r, "Username").(string)

	fmt.Printf("results user: %v\n", user)

	redisConn := redis.Pconnect(3)
	defer redisConn.Close()
	hashMap, reer := redisConn.HgetAll(user)

	if reer != nil {
		if reer == redis.ErrNil {
			w.WriteHeader(http.StatusGone)
		} else {
			panic(reer)
		}
	} else {

		resList := []models.Result{}

		for k, v := range hashMap {
			resList = append(resList, models.Result{k, "[" + v + "]"})
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resList)
	}
}
