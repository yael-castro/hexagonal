package business

import (
	"context"
	"errors"
)

func NewClosingsGenerator(repositoryStore RepositoryClosingStore) (CaseClosingsGenerator, error) {
	if repositoryStore == nil {
		return nil, errors.New("closing store is nil")
	}

	return caseClosingsGenerator{
		repositoryStore: repositoryStore,
	}, nil
}

type caseClosingsGenerator struct {
	repositoryStore RepositoryClosingStore
}

func (c caseClosingsGenerator) GenerateClosings(ctx context.Context, closingIDs ClosingIDs) (err error) {
	err = closingIDs.Validate()
	if err != nil {
		return
	}

	closings, err := c.repositoryStore.GetClosings(ctx, closingIDs)
	if err != nil {
		return err
	}

	errs := make([]error, 0, len(closings))

	for i := range closings {
		errs = append(errs, c.generateClosing(ctx, &closings[i]))
	}

	return errors.Join(err)
}

func (c caseClosingsGenerator) generateClosing(ctx context.Context, closing *Closing) error {
	if closing.IsTime() {
		return nil // Nothing to do. It is not time for close
	}

	return c.repositoryStore.SaveClosing(ctx, closing)
}
