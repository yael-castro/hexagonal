package business

import "fmt"

type ClosingIDs []ClosingID

func (d ClosingIDs) Validate() (err error) {
	for _, id := range d {
		if err = id.Validate(); err != nil {
			return err
		}
	}

	return nil
}

type ClosingID uint64

func (c ClosingID) Validate() error {
	if c == 0 {
		return fmt.Errorf("missing order id to confirm the purchase (%w)", ErrInvalidClosingID)
	}

	return nil
}

type Closing struct {
	ID       uint64
	Payments []Payment
}

// IsTime indicates if is time to close the sales period
func (c Closing) IsTime() bool {
	return len(c.Payments) > 0
}

type Payment struct {
	ID    uint64
	Value float64
}
