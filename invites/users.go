package invites

type Contact struct {
	ID string

	Type string // TODO enum

	Contact string // TODO Validate: phones are E164, emails are valid

	Active bool

	// a user may opt not to be contacted via this contact
	// e.g if it's a shared phone owned by a teenager
	OptedIn bool
}

type User struct {
	ID *string // globally unique ID

	Username string // @handle, also globally unique; nickname

	DisplayName string // user's preferred display name

	// TODO Consider making the names optional in DB; validation in frontends
	FirstName  string // given name
	MiddleName string
	LastName   string

	UserType string // TODO enum; e.g client, health care worker

	Gender string // TODO enum; genders; keep it simple

	Active bool

	Contacts  []*Contact // TODO: validate, ensure
	PushToken string
}

type PatientProfile struct {
	ID *string

	UserID string // foreign key to user
}

type StaffProfile struct {
	ID *string

	UserID string // foreign key to user

	StaffNumber string

	Facilities []*Facility

	// there is nothing special about super-admin; just the set of roles they have
	Roles []string // TODO: roles are an enum (controlled list), known to both FE and BE

	// TODO PIN Expiry
}

type IUserInvite interface {

	// TODO: app is an enum; client or pro app
	// TODO: send invite link via e.g SMS
	//    the invite deep link: opens the app if installed OR goes to the store if not installed
	//    a first time PIN is set and sent to the user
	//    this PIN must be changed on first use
	//    this PIN can be used only once
	// TODO: generate first time PIN, must change, link to user
	// TODO determine communication channel for invite (e.g SMS) from settings
	Invite(userID string, app string) (bool, error)
}

// IUserForget models the behavior needed to comply with privacy laws e.g GDPR
// and "forget me"
type IUserForget interface {

	// Forget inactivates the user record AND hashes all identifiable information
	// After this: the user should not be available on user lists or able to log in
	// After this: it should not be possible to re-identify the user
	// This is irreversible and the UX should ensure confirmation
	// Validate: A user can only forget themselves
	// Validate: PIN is correct
	Forget(userID string, pin string) (bool, error)
}

// IRequestDataExport allows a user to request data known about them
// Mostly for legal compliance.
// The first impl. will simply create a task (for manual follow up) and acknowledge
type IRequestDataExport interface {
	RequestDataExport(userID string, pin string) (bool, error)
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

type IInactivateProfile interface {
	// TODO profileType is an enum
	InactivateProfile(userID string, profileType string) (bool, error)
}

type IReactivateProfile interface {
	// TODO profileType is an enum
	ReactivateProfile(userID string, profileType string) (bool, error)
}

type UserUseCases interface {
	IUserInvite
	IUserForget
	IRequestDataExport
}

type PatientProfileUsecases interface {
	IInactivateProfile
	IReactivateProfile
}

type StaffProfileUsecases interface {
	IAddRoles
	IRemoveRole
	IInactivateProfile
	IReactivateProfile
}

// listing users
// search for users

// some users have expiring PINs e.g professionals
// user preferences i.e preferred languages, opt-ins, promotional messages etc

// fetch profile
// update profile
// hashing / masking of UIDs for anonymization e.g for data analysis
