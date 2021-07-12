package shark

import (
	"net/http"
	"time"

	"pkg.agungdp.dev/candi/candiutils"
)

type sharkRESTImpl struct {
	host    string
	authKey string
	httpReq candiutils.HTTPRequest
}

// NewSharkServiceREST constructor
func NewSharkServiceREST(host string, authKey string) Shark {

	return &sharkRESTImpl{
		host:    host,
		authKey: authKey,
		httpReq: candiutils.NewHTTPRequest(
			candiutils.HTTPRequestSetRetries(5),
			candiutils.HTTPRequestSetSleepBetweenRetry(500*time.Millisecond),
			candiutils.HTTPRequestSetHTTPErrorCodeThreshold(http.StatusBadRequest),
			candiutils.HTTPRequestSetBreakerName("shark"),
		),
	}
}
