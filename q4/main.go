package main

import (
	"fmt"
	"net/http"

	"github.com/stayup/q4/controller"
)

func main() {
	http.HandleFunc("/subscribe", controller.SubscribeHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to start server, error:%v", err)
	}
}