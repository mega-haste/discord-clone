package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"upchat.com/server/socket"

	"upchat.com/server/model"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panicln("Error while leading .env file.", err)
	}

	model.Init()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
	}))

	initRoutes(router)

	router.GET("/ws", func(ctx *gin.Context) {
		socket.SocketServer.Handle(ctx.Writer, ctx.Request)
	})

	fmt.Println("Server is starting!!")
	router.Run(":8080")
}
