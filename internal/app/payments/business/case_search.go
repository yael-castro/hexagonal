package business

import (
	"context"
	"errors"
)

func NewPaymentsSearcher(repositorySearcher RepositoryPaymentSearcher) (CasePaymentsSearcher, error) {
	if repositorySearcher == nil {
		return nil, errors.New("repository searcher is nil")
	}

	return casePaymentsSearcher{
		repositorySearcher: repositorySearcher,
	}, nil
}

type casePaymentsSearcher struct {
	repositorySearcher RepositoryPaymentSearcher
}

func (c casePaymentsSearcher) SearchPayments(ctx context.Context, filter *PaymentFilter) ([]Payment, error) {
	err := filter.Validate()
	if err != nil {
		return []Payment{}, nil
	}

	return c.repositorySearcher.SearchPayments(ctx, filter)
}
