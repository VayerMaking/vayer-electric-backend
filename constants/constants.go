package constants

import (
	"time"

	"vayer-electric-backend/env"
)

var (
	ShutdownTimeout         = time.Duration(env.SHUTDOWN_TIMEOUT) * time.Second
	RequestTimeout          = time.Duration(env.REQUEST_TIMEOUT) * time.Second
	StatsdFlushInterval = time.Duration(env.STATSD_FLUSH) * time.Millisecond
)

