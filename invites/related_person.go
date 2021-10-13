package invites

import "time"

type RelatedPerson struct {
	ID string

	Active           bool
	RelatedTo        string // TODO: FK to client
	RelationshipType string // TODO: enum
	FirstName        string
	LastName         string
	OtherName        string // TODO: optional
	Gender           string // TODO: enum

	DateOfBirth *time.Time // TODO: optional
	Addresses   []*Address // TODO: optional
	Contacts    []*Contact // TODO: optional
}

// TODO: need CRUD for this; to ensure viewing and editing of client profiles in the UI works
