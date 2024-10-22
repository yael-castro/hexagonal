//go:build cmd

package container

import (
	"context"
	"database/sql"
	"github.com/yael-castro/payments/internal/app/closings/business"
	"github.com/yael-castro/payments/internal/app/closings/input/command"
	"github.com/yael-castro/payments/internal/app/closings/output/postgres"
	"log"
	"os"
)

func New() Container {
	return new(cmd)
}

type cmd struct {
	container
}

func (r *cmd) Inject(ctx context.Context, a any) (err error) {
	switch a := a.(type) {
	case *func(context.Context, ...string) int:
		return r.injectCommand(ctx, a)
	}

	return r.container.Inject(ctx, a)
}

func (r *cmd) injectCommand(ctx context.Context, cmd *func(context.Context, ...string) int) (err error) {
	// External dependencies
	errLogger := log.New(os.Stderr, "[ERROR] ", log.LstdFlags)
	// infoLogger := log.New(os.Stdout, "[INFO] ", log.LstdFlags)

	var db *sql.DB
	if err = r.Inject(ctx, &db); err != nil {
		return
	}

	// Secondary adapters
	closingsStore, err := postgres.NewClosingsStore(db)
	if err != nil {
		return
	}

	// Business logic
	closingsGenerator, err := business.NewClosingsGenerator(closingsStore)
	if err != nil {
		return
	}

	// Error handling
	errFunc := func(error) int {
		return 1
	}
	errFunc = command.ErrorFunc(errFunc)

	// Primary adapters
	*cmd = command.GenerateClosings(closingsGenerator, errLogger, errFunc)
	return
}
