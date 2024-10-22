package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/yael-castro/payments/internal/app/closings/business"
)

func NewClosingsStore(db *sql.DB) (business.RepositoryClosingStore, error) {
	if db == nil {
		return nil, errors.New("db is nil")
	}

	return repositoryClosingsStore{
		db: db,
	}, nil
}

type repositoryClosingsStore struct {
	db *sql.DB
}

func (r repositoryClosingsStore) GetClosings(ctx context.Context, ds business.ClosingIDs) ([]business.Closing, error) {
	//TODO implement me
	return []business.Closing{}, nil
}

func (r repositoryClosingsStore) SaveClosing(ctx context.Context, closing *business.Closing) error {
	//TODO implement me
	return nil
}
