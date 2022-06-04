package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"todomvc/db"
	"todomvc/handler"
)
func initServer(engine *gin.Engine) *http.Server{
	return &http.Server{
		Addr:              ":8080",
		Handler:           engine,
		ReadTimeout:       10 *time.Second,
		WriteTimeout:      10 *time.Second,
		MaxHeaderBytes:    1 << 20,
	}
}

func initRouter(engine *gin.Engine) {

	r := engin.Group("api")
	{
		r.POST("add", handler.Add)
		r.POST("del", handler.Del)
		r.POST("update", handler.Update)
		r.POST("find", handler.Find)
	}
}


func main() {
	engine := gin.Default()
	engine.Use()
	initRouter(engine)
	server := initServer(engine)
	db.InitDB()
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
