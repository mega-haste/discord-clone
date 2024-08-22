package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"upchat.com/server/controllers/api"
)

func initRoutes(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "Api are in /api",
		})
	})

	apiRouter := router.Group("/api")
	apiRouter.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
	}))
	apiRouter.GET("/", api.Index)
	apiRouter.GET("/global/messages", api.GetGlobalMessages)
	apiRouter.POST("/login", api.PostLogin)
	apiRouter.GET("/login", api.GetLogin)
}
