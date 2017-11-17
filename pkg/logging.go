package hasher

import (
	"context"
	"github.com/go-kit/kit/log"
	"time"
)

type loggingMiddleware struct {
	logger log.Logger
	next   HashService
}

func (mw loggingMiddleware) HashSHA256(ctx context.Context, input string) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "hashsha256",
			"input", input,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.next.HashSHA256(ctx, input)
	return
}

// NewLoggingMiddleware is a constructor for the logging middleware to be used by the HashService
func NewLoggingMiddleware(logger log.Logger, next HashService) loggingMiddleware {
	return loggingMiddleware{logger, next}
}
