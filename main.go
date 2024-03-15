package main

import (
	"fmt"
	"rankbp-be/config"
	"rankbp-be/model"
	"rankbp-be/router"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(config.Config.AppMode)
	srv := router.NewServer()
	model.InitModel()

	if err := srv.ListenAndServe(); err != nil {
		fmt.Printf("fail to init server: %s\n", err.Error())
		panic(err)
	}

}
