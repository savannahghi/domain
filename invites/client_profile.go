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

	TreatmentEnrollmentDate *time.Time // use for date of treatment enrollment

	ClientType string // TODO: enum; e.g PMTCT, OVC

	Active bool

	HealthRecordID *string // optional link to a health record e.g FHIR Patient ID

	// TODO: a client can have many identifiers; an identifier belongs to a client
	// (implement reverse relation lookup)
	Identifiers []*Identifier

	Addresses []*Address

	RelatedPersons []*RelatedPerson // e.g next of kin

	// client's currently assigned facility
	FacilityID string // TODO: FK

	TreatmentBuddyUserID string // TODO: optional, FK to User

	CHVUserID string // TODO: optional, FK to User

	ClientCounselled bool
}

type ClientProfileRegistrationPayload struct {
	// every client is a user first
	// biodata is linked to the user record
	// the client record is for bridging to other identifiers e.g patient record IDs
	UserID string // TODO: Foreign key to User

	ClientType string // TODO: enum; e.g PMTCT, OVC

	PrimaryIdentifier *Identifier // TODO: optional, default set if not givemn

	Addresses []*Address

	FacilityID string

	TreatmentEnrollmentDate *time.Time

	ClientCounselled bool

	// TODO: when returning to UI, calculate length of treatment (return as days for ease of use in frontend)
}

type IRegisterClient interface {
	// TODO: the input client profile must not have an ID set
	//		validate identifiers when creating
	//		if the enrollemnt date is not supplied, set it automatically
	//		default to the client profile being active right after creation
	//		create a patient on FHIR (HealthRecordID
	//		if identifers not supplied (e.g patient being created on app), set
	//			an internal identifier as the default. It should be updated later
	//			with the CCC number or other final identifier
	// TODO: ensure the user exists...supplied user ID
	// TODO: only register clients who've been counselled
	// TODO: consider: after successful registration, send invite link automatically
	RegisterClient(user User, profile ClientProfileRegistrationPayload) (*ClientProfile, error)
}

type IAddClientIdentifier interface {
	// TODO idType is an enum
	// TODO use idType and settings to decide if it's a primary identifier or not
	AddIdentifier(clientID string, idType string, idValue string, isPrimary bool) (*Identifier, error)
}

type IInactivateClient interface {
	// TODO Consider making reasons an enum
	InactivateClient(clientID string, reason string, notes string) (bool, error)
}

type IReactivateClient interface {
	ReactivateClient(clientID string, reason string, notes string) (bool, error)
}

type ITransferClient interface {
	// TODO: maintain log of past transfers, who did it etc
	TransferClient(
		clientID string,
		OriginFacilityID string,
		DestinationFacilityID string,
		Reason string, // TODO: consider making this an enum
		Notes string, // optional notes...e.g if the reason given is "Other"
	) (bool, error)
}

type IGetClientIdentifiers interface {
	GetIdentifiers(clientID string, active bool) ([]*Identifier, error)
}

type IInactivateClientIdentifier interface {
	InactivateIdentifier(clientID string, identifierID string) (bool, error)
}

type IAssignTreatmentSupporter interface {
	AssignTreatmentSupporter(
		clientID string,
		treatmentSupporterID string,
		treatmentSupporterType string, // TODO: enum, start with CHV and Treatment buddy
	) (bool, error)
}

type IUnassignTreatmentSupporter interface {
	UnassignTreatmentSupporter(
		clientID string,
		treatmentSupporterID string,
		reason string, // TODO: ensure these are in an audit log
		notes string, // TODO: Optional
	) (bool, error)
}

type IAddRelatedPerson interface {
	// add next of kin
	AddRelatedPerson(
		clientID string,
		relatedPerson *RelatedPerson,
	) (*RelatedPerson, bool)
}

type ClientProfileUseCases interface {
	IAddClientIdentifier
	IGetClientIdentifiers
	IInactivateClientIdentifier
	IRegisterClient
	IInactivateClient
	IReactivateClient
	ITransferClient
	IAssignTreatmentSupporter
	IUnassignTreatmentSupporter
	IAddRelatedPerson
}

// TODO: Client profile CRUD
//	client profile search...search across identifiers + human readable fields
//	list filters to include facility
//  fetch profile by user ID also...e.g to get profile after login
//	inline into the profile a calculated field with treatment duration
