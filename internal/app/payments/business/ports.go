package business

import (
	"context"
)

// Ports for drive adapters
type (
	// CasePaymentsSearcher defines a way to search payments
	CasePaymentsSearcher interface {
		SearchPayments(context.Context, *PaymentFilter) ([]Payment, error)
	}
)

// Ports for driven adapters
type (
	// RepositoryPaymentSearcher defines the common operations related to the Payment storage
	RepositoryPaymentSearcher interface {
		SearchPayments(context.Context, *PaymentFilter) ([]Payment, error)
	}
)
