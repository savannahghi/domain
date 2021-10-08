package invites

type User struct {
	ID *string // globally unique ID

	Username string // @handle, also globally unique; nickname

	DisplayName string // user's preferred display name

	// TODO Consider making the names optional in DB; validation in frontends
	FirstName  string // given name
	MiddleName *string
	LastName   string

	UserType string // TODO enum; e.g client, health care worker

	Gender string // TODO enum; genders; keep it simple

	Active bool

	Contacts []*Contact // TODO: validate, ensure

	PushToken string
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

type UserUseCases interface {
	IUserInvite
	IUserForget
	IRequestDataExport
}

// TODO: CRUD for users...including search
// search for users

// some users have expiring PINs e.g professionals
// user preferences i.e preferred languages, opt-ins, promotional messages etc
// hashing / masking of UIDs for anonymization e.g for data analysis
