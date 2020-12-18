package handler

import (
	"log"
	"net/http"

	"github.com/rayspock/go-answer/utils/exception"

	"github.com/gin-gonic/gin"
)

//ErrorHandler ... Error handler for each http request
func ErrorHandler(f func(*gin.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := f(c)
		if err != nil {
			log.Println(err)
			if re, ok := err.(*exception.RequestError); ok {
				er := exception.HTTPError{
					Code:    re.Code,
					Message: re.Error(),
				}
				c.AbortWithStatusJSON(re.Code, er)
				return;
			}
			er := exception.HTTPError{
				Code:    http.StatusInternalServerError,
				Message: "Internal Server Error",
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, er)
		}
	}
}