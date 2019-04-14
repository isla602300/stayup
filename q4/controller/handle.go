package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/redigo/redis"
	rPool "github.com/stayup/q4/redis"
)

var (
	totalNum = 0
)

func SubscribeHandler(w http.ResponseWriter, r *http.Request) {
	// if time.Now().Hour() < 8 {
	// 	w.Write([]byte("sorry, you can not subscribe before 8 am!"))
	// 	return
	// }
	if totalNum == 1000 {
		w.Write([]byte("sorry, there is no vacancy available for you today!"))
		return
	}
	// 1. 解析用户请求参数
	r.ParseForm()
	name := r.Form.Get("name")
	IdNumber := r.Form.Get("idnumber")
	// 2. 获得redis连接池中的一个连接
	rConn := rPool.RedisPool().Get() //Get()从连接池中取出连接
	defer rConn.Close()
	// 3.查询redis中是否已经存在数据
	data, err := rConn.Do("HGET", "subscribe", "id_"+IdNumber)
	if err != nil && err != redis.ErrNil {
		fmt.Println(err.Error())
	}
	if data != nil {
		w.Write([]byte("you have subscribed before!"))
		return
	}
	// 4. 更新redis缓存状态
	rConn.Do("HSET", "subscribe", "id_"+IdNumber, name)
	totalNum++
	// 5. 返回处理结果到客户端
	w.Write([]byte("subscribe success!"))
	// 6. 清空一天的Redis
	if time.Now().Hour() == 7 && totalNum != 0 {
		totalNum = 0
		err = rConn.Flush()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
