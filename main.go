package main

import (
	"SFGPL-End/biz/route"
	"github.com/gin-contrib/cors"
	"net/http"
	"time"
)

func main() {

	r := route.CreatedRouter()
	r.Use(cors.Default())

	r.Static("/file", "./static")
	s := &http.Server{
		Addr:           "localhost:18080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		panic(err.Error())
		return
	}

}
