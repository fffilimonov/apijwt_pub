package controllers

import (
	"../core/redis"
	b64 "encoding/base64"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)


func ActivateCaptchaController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	vars := mux.Vars(r)
	user := vars["user"]
	tmp, _ := b64.URLEncoding.DecodeString(user)
	userEnc := string(tmp)
	fmt.Printf("Username: %v\n", userEnc)

	redisConn := redis.Pconnect(2)
	defer redisConn.Close()
	hashMap, reer := redisConn.HgetAll(userEnc)
	if reer != nil {
		panic(reer)
	}
	if _, ok := hashMap["Password"]; ok {
		redisConn.HsetValue(userEnc, "Active", "1")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("not exist\n"))
	}
}
