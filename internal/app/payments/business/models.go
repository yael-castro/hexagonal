package business

import (
	"fmt"
	closings "github.com/yael-castro/payments/internal/app/closings/business"
)

type PaymentFilter struct {
	Keyword string
	Page    uint16
	Size    uint16
}

func (f *PaymentFilter) Validate() error {
	if f == nil {
		return fmt.Errorf("%w: missing payment filter to improve the search", ErrInvalidPaymentFilter)
	}

	return nil
}

// Aliases
type (
	Payment = closings.Payment
)
