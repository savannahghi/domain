package invites

type StaffProfile struct {
	ID *string

	UserID string // foreign key to user

	StaffNumber string

	Facilities []*Facility

	// there is nothing special about super-admin; just the set of roles they have
	Roles []string // TODO: roles are an enum (controlled list), known to both FE and BE

	Addresses []*Address
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

type StaffProfileUsecases interface {
	IAddStaffUser
	IAddRoles
	IRemoveRole
	IUpdateLanguagePreferences
}

// TODO: CRUD for StaffProfile
