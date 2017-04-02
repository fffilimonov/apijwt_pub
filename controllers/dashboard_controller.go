package controllers

import (
	"../core/redis"
	"../services/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/context"
	"net/http"
)

func DashboardController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	user := context.Get(r, "Username").(string)
	fmt.Printf("Username: %v\n", user)

	redisConn := redis.Pconnect(1)
	defer redisConn.Close()
	redisConnUser := redis.Pconnect(2)
	defer redisConnUser.Close()

	hashMap, reer := redisConn.HgetAll(user)
	if reer != nil {
		panic(reer)
	}

	userHashMap, reeru := redisConnUser.HgetAll(user)
	if reer != nil {
		panic(reeru)
	}

	scensList := []models.Scen{}

	resDashboard := new(models.Dashboard)
	resDashboard.Seconds = userHashMap["Seconds"]

	for k, v := range hashMap {
		scensList = append(scensList, models.Scen{k, v})
	}

	resDashboard.Scens = scensList

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resDashboard)
}
