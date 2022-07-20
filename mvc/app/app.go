package app

import (
	"net/http"

	"github.com/walbety/go-microservice/mvc/controller"
)

func StartApp() {
	http.HandleFunc("/user", controller.GetUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
