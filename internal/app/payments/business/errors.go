package business

import "strconv"

// Supported values for PaymentError
const (
	_ PaymentError = iota
	ErrInvalidPaymentFilter
)

type PaymentError uint8

func (e PaymentError) Error() string {
	return "PAYMENTS_PAYMENTS_" + strconv.FormatUint(uint64(e), 10)
}
