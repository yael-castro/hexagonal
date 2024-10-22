package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRoutes(e *gin.Engine, h Handler) {
	errorHandler := ErrorHandler(func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, gin.H{})
	})

	e.GET("/v1/payments", h.GetPayments, errorHandler)
}
