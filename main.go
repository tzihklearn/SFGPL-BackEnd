package main

import (
	"SFGPL-End/biz/dal"
	"SFGPL-End/biz/route"
	"github.com/gin-contrib/cors"
	"net/http"
	"time"
)

func main() {

	dal.Init()
	//创建路由
	r := route.CreatedRouter()
	//配置跨域处理
	r.Use(cors.Default())

	//配置服务监听端口等信息
	s := &http.Server{
		Addr:           "localhost:18080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	//开启服务
	err := s.ListenAndServe()
	if err != nil {
		panic(err.Error())
		return
	}

}
