package invites

import "time"

// TODO: introduce into all tables a base model with created, updated, created by, updated by
//	and logic to fill/update these

// AuditLog is used to record all senstive changes
//
// e.g
//	- changing a client's treatment buddy
// 	- changing a client's facility
// 	- deactivating a client
//	- changing a client's assigned community health volunteer
//
// Rules of thumb: is there a need to find out what/when/why something
// occured? Is a mistake potentially serious? Is there potential for
// fraud?
type AuditLog struct {
	ID         string
	TimeStamp  time.Time
	RecordType string // e.g auditing facility changes etc
	Notes      string

	// this will vary by context
	// should not identify the user (there's a UID field)
	// focus on the actual event
	Payload map[string]string
}

// TODO: CRUD for audit logs
