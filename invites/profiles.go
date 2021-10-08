package invites

// ClientProfile holds the details of end users who are not using the system in
// a professional capacity e.g consumers, patients etc
type ClientProfile struct {
	ID *string

	UserID string // foreign key to user

	// a client may have many identifiers
	// this is a foreign key to the client record
	// all other identifiers they have will be on the client record
	ClientID string

	// for the preferred language list, order matters
	Languages []string // TODO: turn this into a slice of enums, start small (en, sw)

	Addresses []*Address
}

type StaffProfile struct {
	ID *string

	UserID string // foreign key to user

	StaffNumber string

	Facilities []*Facility

	// there is nothing special about super-admin; just the set of roles they have
	Roles []string // TODO: roles are an enum (controlled list), known to both FE and BE

	Addresses []*Address
}

type IInactivateProfile interface {
	// TODO flavour is an enum e.g CLIENT, PRO
	InactivateProfile(userID string, flavour string) (bool, error)
}

type IReactivateProfile interface {
	// TODO flavour is an enum e.g CLIENT, PRO
	ReactivateProfile(userID string, flavour string) (bool, error)
}

type IAddStaffUser interface {
	// TODO: transactional i.e create both OR none
	AddStaffUser(user User, profile StaffProfile) (*User, *StaffProfile, error)
}

type IAddRoles interface {
	AddRoles(userID string, roles []string) (bool, error)
}

type IRemoveRole interface {
	RemoveRole(userID string, role string) (bool, error)
}

type IUpdateLanguagePreferences interface {
	UpdateLanguagePreferences(userID string, language string) (bool, error)
}

type ClientProfileUsecases interface {
	IInactivateProfile
	IReactivateProfile
}

type StaffProfileUsecases interface {
	IAddStaffUser
	IAddRoles
	IRemoveRole
	IInactivateProfile
	IReactivateProfile
	IUpdateLanguagePreferences
}

// TODO: CRUD for ClientProfile
// TODO: CRUD for StaffProfile
