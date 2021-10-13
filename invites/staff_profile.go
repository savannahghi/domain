package invites

type StaffProfile struct {
	ID *string

	UserID string // foreign key to user

	StaffNumber string

	Facilities []*Facility // TODO: needs at least one

	// A UI switcher optionally toggles the default
	// TODO: the list of facilities to switch between is strictly those that the user is assigned to
	DefaultFacilityID string // TODO: required, FK to facility

	// there is nothing special about super-admin; just the set of roles they have
	Roles []string // TODO: roles are an enum (controlled list), known to both FE and BE

	Addresses []*Address
}

type IRegisterStaffUser interface {
	// TODO: ensure default facility is set
	//		validation: ensure the staff profile has at least one facility
	//		ensure that the default facility is one of these
	// TODO: ensure the user exists...userID in profile
	RegisterStaffUser(user User, profile StaffProfile) (*User, *StaffProfile, error)
}

type IAddRoles interface {
	AddRoles(userID string, roles []string) (bool, error)
}

type IRemoveRole interface {
	RemoveRole(userID string, role string) (bool, error)
}

type IUpdateDefaultFacility interface {
	// TODO: the list of facilities to switch between is strictly those that the user is assigned to
	UpdateDefaultFacility(userID string, facilityID string) (bool, error)
}

type StaffProfileUsecases interface {
	IRegisterStaffUser
	IAddRoles
	IRemoveRole
	IUpdateLanguagePreferences
	IUpdateDefaultFacility
}

// TODO: CRUD for StaffProfile
