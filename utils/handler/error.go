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
				c.AbortWithError(re.Code, re)
				return;
			}
			c.AbortWithError(http.StatusInternalServerError, err)
		}
	}
}