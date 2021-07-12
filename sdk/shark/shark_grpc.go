package shark

import (
	"net/url"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
)

type sharkGRPCImpl struct {
	host    string
	authKey string
	conn    *grpc.ClientConn
}

// NewSharkServiceGRPC constructor
func NewSharkServiceGRPC(host string, authKey string) Shark {

	if u, _ := url.Parse(host); u.Host != "" {
		host = u.Host
	}
	conn, err := grpc.Dial(host, grpc.WithInsecure(), grpc.WithConnectParams(grpc.ConnectParams{
		Backoff: backoff.Config{
			BaseDelay:  50 * time.Millisecond,
			Multiplier: 5,
			MaxDelay:   50 * time.Millisecond,
		},
		MinConnectTimeout: 1 * time.Second,
	}))
	if err != nil {
		panic(err)
	}

	return &sharkGRPCImpl{
		host:    host,
		authKey: authKey,
		conn:    conn,
	}
}
