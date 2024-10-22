package postgres

import (
	"database/sql"
	"github.com/yael-castro/payments/internal/app/payments/business"
)

type Payment struct {
	ID    sql.NullInt64
	Value sql.NullFloat64
}

func (p Payment) ToBusiness() business.Payment {
	return business.Payment{
		ID:    uint64(p.ID.Int64),
		Value: p.Value.Float64,
	}
}

func NewPaymentFilter(filter business.PaymentFilter) PaymentFilter {
	return PaymentFilter{
		Keyword: sql.NullString{
			String: filter.Keyword,
			Valid:  len(filter.Keyword) > 0,
		},
		Page: sql.NullInt64{
			Int64: int64(filter.Page),
			Valid: filter.Page > 0,
		},
		Size: sql.NullInt64{
			Int64: int64(filter.Size),
			Valid: filter.Size > 0,
		},
	}
}

type PaymentFilter struct {
	Keyword sql.NullString
	Page    sql.NullInt64
	Size    sql.NullInt64
}
