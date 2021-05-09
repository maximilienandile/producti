package server

import "github.com/gin-gonic/gin"

func (s *Server) CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Cache-Control")
		c.Next()
	}
}
