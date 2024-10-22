package command

import (
	"context"
	"flag"
	"github.com/yael-castro/payments/internal/app/closings/business"
	"log"
	"strconv"
	"strings"
)

func GenerateClosings(maker business.CaseClosingsGenerator, errLogger *log.Logger, errFunc func(error) int) func(context.Context, ...string) int {
	return func(ctx context.Context, args ...string) int {
		// Declaring flags
		str := ""
		flags := flag.NewFlagSet(args[0], flag.ContinueOnError)

		// Parse command arguments
		flags.StringVar(&str, "closings", "", "csv with the specified closing IDs (optional)")

		err := flags.Parse(args[1:])
		if err != nil {
			flags.Usage()
			return errFunc(err)
		}

		// Decoding closing IDs
		ids := strings.Split(str, ",")
		closingIDs := make(business.ClosingIDs, 0, len(ids))

		for _, id := range ids {
			var closingID uint64

			closingID, err = strconv.ParseUint(id, 10, 64)
			if err != nil {
				flags.Usage()
				return errFunc(err)
			}

			closingIDs = append(closingIDs, business.ClosingID(closingID))
		}

		// Dispatch business case
		err = maker.GenerateClosings(ctx, closingIDs)
		if err != nil {
			errLogger.Println(err)
			return errFunc(err)
		}

		return errFunc(err)
	}
}
