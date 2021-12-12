package controllers

import "github.com/gin-gonic/gin"

func Response(c *gin.Context, body gin.H, statusCode int) {
	ctx := c.Request.Context()

	select {
		case <- ctx.Done():
			return

		default:
			c.JSON(statusCode, body)
	}
}