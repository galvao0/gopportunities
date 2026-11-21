package router

import (
	docs "github.com/galvao0/gopportunities/docs"
	"github.com/galvao0/gopportunities/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(r *gin.Engine) {

	//Initialize Handler
	handler.Init()

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := r.Group(basePath)
	{
		v1.GET("/opening", handler.ShowOpeningHadler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
	}

	// Initialize Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
