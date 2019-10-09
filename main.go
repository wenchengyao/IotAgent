package main

import (
	"log"

	"./agent"
	"./cmd"
	"./provider/shunzhou"
	"./provider/xixun"
	"github.com/gin-gonic/gin"
)

var Client cmd.Cli

type Req struct {
}

func main() {
	xixun.Init(7077)
	shunzhou.Init(7078)
	agent.Init("L1000000001", "127.0.0.1", 19999)

	r := gin.Default()
	r.POST("/people", OperateDevice)
	//	r.GET("/device/", GetDevice)
	//	r.GET("/device/:id", GetDevice)
	//	r.PUT("/people/:id", UpdateDevice)
	//	r.DELETE("/people/:id", DeleteDevice)
	r.Run(":8080")
}

func GetDevice(c *gin.Context) {
	//	id := c.Params.ByName("id")
	log.Fatalln("查询指定设备")
}

func OperateDevice(c *gin.Context) {

}
