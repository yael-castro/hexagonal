package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/yael-castro/payments/internal/app/payments/business"
)

func NewRepositoryPaymentSearcher(db *sql.DB) (business.RepositoryPaymentSearcher, error) {
	if db == nil {
		return nil, errors.New("db is required")
	}

	return repositoryPaymentSearcher{
		db: db,
	}, nil
}

type repositoryPaymentSearcher struct {
	db *sql.DB
}

func (r repositoryPaymentSearcher) SearchPayments(ctx context.Context, f *business.PaymentFilter) (payments []business.Payment, err error) {
	filter := NewPaymentFilter(*f)

	stmt, args, err := SelectPayments(filter)
	if err != nil {
		return
	}

	rows, err := r.db.QueryContext(ctx, stmt, args...)
	if err != nil {
		return
	}

	const expectedPayments = 10
	payments = make([]business.Payment, 0, expectedPayments)

	for rows.Next() {
		var payment Payment

		err = rows.Scan(
			payment.ID,
			payment.Value,
		)
		if err != nil {
			return
		}

		payments = append(payments, payment.ToBusiness())
	}

	return
}
