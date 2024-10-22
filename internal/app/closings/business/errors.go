package business

import "strconv"

// Supported values for ClosingError
const (
	_ ClosingError = iota
	ErrInvalidClosingID
)

type ClosingError uint8

func (e ClosingError) Error() string {
	return "PAYMENTS_CLOSINGS_" + strconv.FormatUint(uint64(e), 10)
}
