package webapi

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	gin.SetMode(gin.ReleaseMode)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/view/cases", viewCasesHandler)
	}
	return r
}
