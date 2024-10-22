package business

import (
	"context"
)

// Ports for drive adapters
type (
	// CaseClosingsGenerator defines a way to make closings
	CaseClosingsGenerator interface {
		// GenerateClosings make closings with ...
		GenerateClosings(context.Context, ClosingIDs) error
	}
)

// Ports for driven adapters
type (
	// RepositoryClosingStore defines the common operations related to the closings storage
	RepositoryClosingStore interface {
		GetClosings(context.Context, ClosingIDs) ([]Closing, error)
		SaveClosing(context.Context, *Closing) error
	}
)
