package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		detectedErrors := c.Errors
		// check if there is at least one error in the gin
		// context
		if len(detectedErrors) > 0 {
			for k, v := range detectedErrors {
				log.Printf("[ERROR N°%d] %s", k, v)
			}
			statusCode := c.Writer.Status()
			if statusCode == 0 || statusCode == http.StatusOK {
				c.Status(http.StatusInternalServerError)
			}
			c.Abort()
			return
		}
		// No error continue
	}
}
