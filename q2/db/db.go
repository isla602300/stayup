package db

import (
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

var (
	orderCount = 0
)

func init() {
	db, _ = sql.Open("mysql", "root:isla602300@tcp(localhost:3306)/gamepay?charset=utf8")
	db.SetMaxOpenConns(1000)
	err := db.Ping()
	if err != nil {
		fmt.Printf("failed to connect to mysql, error: %v", err)
		os.Exit(1)
	}
}

func QueryPassword(username string) (string, error) {
	stmt, err := db.Prepare("select user_name, password from tbl_user where user_name=? limit 1")
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	defer stmt.Close()
	password := ""
	err = stmt.QueryRow(username).Scan(&password)
	return password, err
}

func UpdateUser(sum float64, username string) (err error) {
	stmt, err := db.Prepare("replace into tbl_user_game(`coins`) values (?) where user_name=?")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer stmt.Close()
	_, err1 := stmt.Exec(int(sum*100), username)
	if err1 != nil {
		fmt.Println(err1.Error())
		return
	}
	return
}

func UpdateToken(username string, token string) {
	stmt, err := db.Prepare("replace into tbl_user_token(`user_name`, `user_token`) values (?,?)")
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	defer stmt.Close()
	_, err1 := stmt.Exec(username, token)
	if err1 != nil {
		fmt.Println(err1.Error())
		return
	}
	return
}

func DeleteToken(username string) {
	stmt, err := db.Prepare("delete from tbl_user_token where user_name=?")
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	defer stmt.Close()
	_, err1 := stmt.Exec(username)
	if err1 != nil {
		fmt.Println(err1.Error())
		return
	}
	return
}

func CreateOrder(username string, sum float64) {
	stmt, err := db.Prepare("insert ignore into tbl_order (`order_id`,`user_name`,`sum`,`status`) values (?,?,?,0)")
	if err != nil {
		fmt.Printf("failed to add order, error:%v", err)
		return
	}
	defer stmt.Close()
	i1 := strconv.Itoa(rand.Intn(10))
	i2 := strconv.Itoa(rand.Intn(10))
	order_id := fmt.Sprintf("%x", time.Now().UnixNano()) + i1 + i2
	_, err1 := stmt.Exec(order_id, username, sum)
	if err1 != nil {
		fmt.Printf("%v", err1)
		return
	}
}

func CompleteOrder(order_id string) {
	stmt, err := db.Prepare("update tbl_order set status = 1 where order_id = ?")
	if err != nil {
		fmt.Printf("failed to complete order, error:%v", err)
		return
	}
	defer stmt.Close()
	_, err1 := stmt.Exec(order_id)
	if err1 != nil {
		fmt.Printf("%v", err1)
		return
	}
}
