package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/yael-castro/payments/internal/app/payments/business"
	"net/http"
)

func ErrorHandler(errorHandler gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := c.Errors.Last()

		var paymentError business.PaymentError

		if !errors.As(err, &paymentError) {
			errorHandler(c)
		}

		responseCode := http.StatusInternalServerError
		responseBody := gin.H{
			"code":  paymentError.Error(),
			"error": err.Error(),
		}

		switch paymentError {
		case business.ErrInvalidPaymentFilter:
			responseCode = http.StatusBadRequest
		}

		c.JSON(responseCode, responseBody)
	}
}
