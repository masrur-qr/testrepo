package handlers

import (
	"docker/app/controller"

	"github.com/gin-gonic/gin"
)

func Routes()  {
	route := gin.Default()

	route.GET("/data/list",controller.DataList)
	route.POST("/data/add",controller.DataPost)

	route.Run(":2424")
}