//go:build http

package container

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/yael-castro/payments/internal/app/payments/business"
	"github.com/yael-castro/payments/internal/app/payments/input/http"
	"github.com/yael-castro/payments/internal/app/payments/output/postgres"
)

func New() Container {
	return new(handler)
}

type handler struct {
	container
}

func (h *handler) Inject(ctx context.Context, a any) error {
	switch a := a.(type) {
	case **gin.Engine:
		return h.injectHandler(ctx, a)
	}

	return h.container.Inject(ctx, a)
}

func (h *handler) injectHandler(ctx context.Context, e **gin.Engine) (err error) {
	// External dependencies
	var db *sql.DB

	if err = h.Inject(ctx, &db); err != nil {
		return err
	}

	//infoLogger := log.New(os.Stdout, "[INFO] ", log.LstdFlags)
	//errLogger := log.New(os.Stderr, "[ERROR] ", log.LstdFlags)

	// Secondary adapters
	repositorySearcher, err := postgres.NewRepositoryPaymentSearcher(db)
	if err != nil {
		return
	}

	// Business logic
	searcher, err := business.NewPaymentsSearcher(repositorySearcher)
	if err != nil {
		return
	}

	// Primary adapters
	handlers, err := http.NewHandler(searcher)
	if err != nil {
		return
	}

	// Building mux
	engine := gin.New()

	// Setting routes
	http.SetRoutes(engine, handlers)

	// Initializing http.Handler
	*e = engine
	return
}
