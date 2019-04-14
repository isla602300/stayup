package main

import (
	"fmt"
	"net/http"

	"github.com/stayup/q2/controller"
)

func main() {
	http.HandleFunc("/user/login", controller.LoginHandler)
	http.HandleFunc("/user/pay", controller.PayHandler)
	http.HandleFunc("/user/affirm", controller.AffirmHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to start server, error:%v", err)
	}
}
