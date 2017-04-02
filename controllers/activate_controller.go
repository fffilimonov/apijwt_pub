package controllers

import (
	"../core/redis"
	b64 "encoding/base64"
	"fmt"
	"github.com/dpapathanasiou/go-recaptcha"
	"github.com/gorilla/mux"
	"net/http"
)

func processRequest(r *http.Request) (result bool) {
	recaptcha.Init("")
	result = false
	recaptchaResponse, responseFound := r.Form["g-recaptcha-response"]
	if responseFound && recaptchaResponse[0] != "" {
		result = recaptcha.Confirm(r.Header.Get("X-Forwarded-For"), recaptchaResponse[0])
	}
	return
}

func ActivateCaptchaController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	vars := mux.Vars(r)
	user := vars["user"]
	tmp, _ := b64.URLEncoding.DecodeString(user)
	userEnc := string(tmp)
	fmt.Printf("Username: %v\n", userEnc)

	err := r.ParseForm()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("bad form\n"))
	} else {
		if processRequest(r) {
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
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("bad captcha\n"))
		}
	}
}
