package routes

import (
	"os"

	"github.com/rayspock/go-answer/controllers"
	h "github.com/rayspock/go-answer/utils/handler"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/swaggo/files" // swagger embed files

	"github.com/rayspock/go-answer/docs"
)

// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {

	docs.SwaggerInfo.Title = "Bequest Backend Assignment"
	docs.SwaggerInfo.Description = "API and developer documentation."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.Default()
	grp := r.Group("/api")
	grp.Use()
	{
		grp.GET("/ping", func(c *gin.Context) {
			c.String(200, "pong")
		})
		grp.GET("answer/:key", h.ErrorHandler(controllers.GetAnswerByKey))
		grp.GET("answer/:key/history", h.ErrorHandler(controllers.GetAnswerHistoryByKey))
		grp.PUT("answer/:key", h.ErrorHandler(controllers.UpdateAnswerByKey))
		grp.POST("answer/:key", h.ErrorHandler(controllers.CreateAnswerByKey))
		grp.DELETE("answer/:key", h.ErrorHandler(controllers.DeleteAnswerByKey))
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
