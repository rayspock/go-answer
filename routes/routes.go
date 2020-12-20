package routes

import (
	"fmt"
	"log"
	"os"

	"github.com/rayspock/go-answer/controllers"
	h "github.com/rayspock/go-answer/utils/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware

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
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", os.Getenv("PORT"))
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.Default()

	log.Println("Setup CORS...")
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:4000"}
	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true
	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS")

	grp := r.Group("/api")
	grp.Use(cors.New(corsConfig))
	{
		grp.GET("/ping", func(c *gin.Context) {
			c.String(200, "pong")
		})
		grp.GET("answer", h.ErrorHandler(controllers.GetAllAnswer))
		grp.GET("answer/:key", h.ErrorHandler(controllers.GetAnswerByKey))
		grp.GET("answer/:key/history", h.ErrorHandler(controllers.GetAnswerHistoryByKey))
		grp.PUT("answer/:key", h.ErrorHandler(controllers.UpdateAnswerByKey))
		grp.POST("answer", h.ErrorHandler(controllers.CreateAnswerByKey))
		grp.DELETE("answer/:key", h.ErrorHandler(controllers.DeleteAnswerByKey))
	}
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
