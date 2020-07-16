package webapi

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	gin.SetMode(gin.ReleaseMode)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/start", startHandler)
		apiV1.GET("/cases", viewCasesHandler)
		apiV1.GET("/reports", viewReportsHandler)
		apiV1.GET("/report/:case_id", viewReportHandler)
		apiV1.GET("/excel", viewExcelHandler)
	}
	return r
}
