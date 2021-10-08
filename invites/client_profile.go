package invites

import "time"

type Identifier struct {
	ID             *string // globally unique identifier
	ClientID       string  // TODO: FK to client
	IdentifierType string  // TODO: Enum; start with basics e.g CCC number, ID number
	IdentifierUse  string  // TODO: Enum; e.g official, temporary, old (see FHIR Person for enum)

	// TODO: Validate identifier value against type e.g format of CCC number
	// TODO: Unique together: identifier value & type i.e the same identifier can't be used for more than one client
	IdentifierValue     string // the actual identifier e.g CCC number
	Description         string
	ValidFrom           *time.Time
	ValidTo             *time.Time
	Active              bool
	IsPrimaryIdentifier bool
}

// ClientProfile holds the details of end users who are not using the system in
// a professional capacity e.g consumers, patients etc.
//It is a linkage model e.g to tie together all of a person's identifiers
// and their health record ID
type ClientProfile struct {
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
	// (implement reverse relation lookup)
	Identifiers []*Identifier

	Addresses []*Address
}

// TODO: Behavior; add/register client profile
//		validate identifiers when creating
// 		registration payload should be front-end friendly **
// TODO: Inactivate/reactivate client...? reason for inactivate e.g transfer out?
// TODO: Transfer client from facility to facility
// TODO: Behavior; calculate length of treatment
// TODO: Behavior; add treatment buddy (other user)
// TODO: Behavior; remove treatment buddy
// TODO: log in (Pro) to a facility?
// TODO: next of kin / related person
// TODO: client facility...at any time the client has one facility
//	list only active clients OR provide filter facility
// TODO: treatment buddy (optional)...another user
// TODO: next of kin and relationship...next of kin struct...related person

type IAddClientIdentifier interface {
	// TODO idType is an enum
	// TODO use idType and settings to decide if it's a primary identifier or not
	AddIdentifier(clientID string, idType string, idValue string) (*Identifier, error)
}

type IGetClientIdentifiers interface {
	GetIdentifiers(clientID string, active bool) ([]*Identifier, error)
}

type IInactivateClientIdentifier interface {
	InactivateIdentifier(identifierID string) (bool, error)
}

type ClientProfileUseCases interface {
	IAddClientIdentifier
	IGetClientIdentifiers
	IInactivateClientIdentifier
}

// TODO: Client profile CRUD
//	client profile search...search across identifiers + human readable fields
//	list filters to include facility
