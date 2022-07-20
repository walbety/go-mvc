package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/walbety/go-microservice/mvc/service"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {

	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)

		resp.Write([]byte("user_id must be a number"))
		return
	}

	user, err := service.GetUser(userId)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(err.Error()))
		return
	}

	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
