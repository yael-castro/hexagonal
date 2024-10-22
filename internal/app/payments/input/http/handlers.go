package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/yael-castro/payments/internal/app/payments/business"
	"net/http"
)

func NewHandler(searcher business.CasePaymentsSearcher) (Handler, error) {
	if searcher == nil {
		return Handler{}, errors.New("searcher must not be nil")
	}

	return Handler{}, nil
}

type Handler struct {
	paymentsSearcher business.CasePaymentsSearcher
}

func (h Handler) GetPayments(c *gin.Context) {
	var filter PaymentFilter

	err := c.BindQuery(&filter)
	if err != nil {
		_ = c.Error(err)
		return
	}

	ctx := c.Request.Context()

	payments, err := h.paymentsSearcher.SearchPayments(ctx, filter.ToBusiness())
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, payments)
}
