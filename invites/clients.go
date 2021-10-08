package invites

import "time"

// TODO MPI: Client Unique ID

type Identifier struct {
	ID             string // globally unique identifier
	ClientID       string // TODO: FK to client
	IdentifierType string // TODO: Enum; start with basics e.g CCC number, ID number
	IdentifierUse  string // TODO: Enum; e.g official, temporary, old (see FHIR Person for enum)

	// TODO: Validate identifier value against type e.g format of CCC number
	IdentifierValue     string // the actual identifier e.g CCC number
	Description         string
	ValidFrom           *time.Time
	ValidTo             *time.Time
	Active              bool
	IsPrimaryIdentifier bool
}

// TODO: Behavior; validate identifier

// Client is a linkage model e.g to tie together all of a person's identifiers
// and their health record ID
type Client struct {
	ID string // globally unique identifier; synthetic i.e has no encoded meaning

	// every client is a user first
	// biodata is linked to the user record
	// the client record is for bridging to other identifiers e.g patient record IDs
	UserID string // TODO: Foreign key to User

	EnrollmentDate *time.Time // use for date of treatment enrollment

	ClientType string // TODO: enum; e.g PMTCT, OVC

	Active bool

	HealthRecordID *string // optional link to a health record e.g FHIR Patient ID

	// TODO: a client can have many identifiers; an identifier belongs to a client
	// reverse relation
	Identifiers []*Identifier
}

// TODO: Behavior; add/register client
// TODO: registration payload should be front-end friendly **
// TODO: Behavior; ability to look up a client's identifiers (reverse relations)
// TODO: Behavior; add an identifier to a client...embed identifier here
// TODO: Behavior; list a client's idenfiers, optional filter for status
// TODO: Behavior; Modify (e.g retire/inactivate) a client's identifier
// TODO: Behavior; search | should search across identifiers & human readable fields e.g name
//	for search filters: include facility
// TODO: Inactivate/reactivate client...? reason for inactivate e.g transfer out?
// TODO: Transfer client from facility to facility
// TODO: Health Record ID e.g FHIR Patient ID
// TODO: Behavior; calculate length of treatment
// TODO: Behavior; add treatment buddy (other user)
// TODO: Behavior; remove treatment buddy
// TODO: Consider address struct
// TODO: Should this be independent or part of the profile?
// TODO: log in (Pro) to a facility?
// TODO: next of kin / related person
// TODO: client facility...at any time the client has one facility
// TODO: Client CRUD
//	list only active clients OR provide filter facility
// TODO: treatment buddy (optional)...another user
// TODO: next of kin and relationship...next of kin struct...related person

type ClientUseCases interface {
}
