package controller

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/stayup/q2/db"
)

var (
	m sync.Mutex
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	realPassword, err := db.QueryPassword(username)
	if realPassword == "" {
		w.Write([]byte("username not exists!"))
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if password != realPassword {
		w.Write([]byte("password not right!"))
		return
	}
	token := GenToken(username)
	db.UpdateToken(username, token)
}

func MD5(data []byte) string {
	_md5 := md5.New()
	_md5.Write(data)
	return hex.EncodeToString(_md5.Sum([]byte("")))
}

func GenToken(username string) string {
	ts := fmt.Sprintf("%v", time.Now().Unix())
	tokenPrefix := MD5([]byte(username + ts + "_tokensalt"))
	return tokenPrefix + ts[:8]
}

func PayHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	sum := r.Form.Get("sum")
	sum1, err := strconv.ParseFloat(sum, 64)
	if err != nil {
		fmt.Println("string converts to float error!")
	}
	db.CreateOrder(username, sum1)
	http.Redirect(w, r, "/user/affirm", http.StatusFound)
}

func AffirmHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	affirmed := r.FormValue("affirm")
	orderId := r.FormValue("order_id")
	username := r.FormValue("username")
	sum := r.Form.Get("sum")
	sum1, err := strconv.ParseFloat(sum, 64)
	if err != nil {
		fmt.Println("string converts to float error!")
	}
	if affirmed == "1" {
		m.Lock()
		db.CompleteOrder(orderId)
		db.UpdateUser(sum1, username)
		m.Unlock()
	}
	db.DeleteToken(username)
	w.Write([]byte("payment success!"))
	return
}
