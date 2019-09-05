package controllers

import (
	"../core/redis"
	"../services/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/context"
	"github.com/pborman/uuid"
	"net/http"
	"strconv"
	"os/exec"
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

		command := []string{
			"run",
			"-d",
			"-e",
			"USER_ID=" + user,
			"-e",
			"SCEN_ID=" + scen,
			"-e",
			"RES_ID=" + UUID,
			"-e",
			"Browser=" + browser,
			"--add-host=redis-scens:172.17.0.1",
			"-v",
			"/dev/shm:/dev/shm",
			"fffilimonov/bohrium",
		}

	    cmd := exec.Command("docker", command...)
	    stdout, err := cmd.Output()

		fmt.Printf("Err cmd: %v\n", err)
		fmt.Printf("Ok cmd: %v\n", string(stdout))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(UUID))
	} else {
		w.WriteHeader(http.StatusPaymentRequired)
	}
}
