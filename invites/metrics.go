package invites

import "time"

type Metric struct {
	ID string // ensures we don't re-save the same metric; opaque; globally unique

	Type string // TODO Metric types should be a controlled list i.e enum

	// this will vary by context
	// should not identify the user (there's a UID field)
	// focus on the actual event
	Payload map[string]string

	Timestamp time.Time

	// a user identifier, can be hashed for anonymity
	// with a predictable one way hash
	UserID string
}

type IMetricsCollect interface {
	// TODO Check identifier; ensure idempotency
	Collect(metric Metric) error
}

type MetricsUsecases interface {
	IMetricsCollect
}
