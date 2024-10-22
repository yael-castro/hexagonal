package http

import "github.com/yael-castro/payments/internal/app/payments/business"

type PaymentFilter struct {
	Q    string `form:"q"`
	Page int    `form:"page"`
	Size int    `form:"size"`
}

func (p *PaymentFilter) ToBusiness() *business.PaymentFilter {
	return &business.PaymentFilter{
		Keyword: p.Q,
		Page:    uint16(p.Page),
		Size:    uint16(p.Size),
	}
}
