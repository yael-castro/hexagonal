package command

import (
	"errors"
	"github.com/yael-castro/payments/internal/app/closings/business"
)

func ErrorFunc(errFunc func(error) int) func(error) int {
	return func(err error) int {
		var closingErr business.ClosingError

		if !errors.As(err, &closingErr) {
			return errFunc(err)
		}

		switch closingErr {
		case business.ErrInvalidClosingID:
			return 1
		}

		return errFunc(err)
	}
}
