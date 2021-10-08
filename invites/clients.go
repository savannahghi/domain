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

type BioData struct {
	FirstName   string
	MiddleName  *string
	LastName    *string
	DateOfBirth *time.Time
	Gender      string // TODO: enum
}

// Client is a linkage model e.g to tie together all of a person's identifiers
// and their health record ID
type Client struct {
	ID string // globally unique identifier; synthetic i.e has no encoded meaning

	ClientType string // TODO: enum; e.g PMTCT, OVC

	Active bool

	EnrollmentDate *time.Time // use for date of treatment enrollment

	Address string

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
// TODO: Inactivate/reactivate client
// TODO: Client CRUD
//	list only active clients OR provide filter facility
// TODO: Behavior; calculate length of treatment

type ClientUseCases interface {
}

// most important contacts: phone
// county
// facility / clinic
// treatment buddy (optional)
// next of kin and relationship
// action: invite to client app
// action: add to group(s)
// search for a patient...search criteria
// inactivating and reactivating a client
// ? transfer out of our care ?
