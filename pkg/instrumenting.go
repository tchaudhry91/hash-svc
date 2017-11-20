package hasher

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/metrics"
	"time"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	next           HashService
}

func (mw instrumentingMiddleware) HashSHA256(ctx context.Context, input string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "HashSHA256", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	output, err = mw.next.HashSHA256(ctx, input)
	return
}

// NewInstrumentingMiddleware is a constructor for the instrumenting middleware used by the HashService
func NewInstrumentingMiddleware(requestCount metrics.Counter, requestLatency metrics.Histogram, next HashService) instrumentingMiddleware {
	return instrumentingMiddleware{requestCount, requestLatency, next}
}
