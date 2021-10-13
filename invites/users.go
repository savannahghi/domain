package invites

import "time"

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

	// for the preferred language list, order matters
	Languages []string // TODO: turn this into a slice of enums, start small (en, sw)

	PushTokens []string

	// when a user logs in successfully, set this
	LastSuccessfulLogin *time.Time

	// whenever there is a failed login (e.g bad PIN), set this
	// reset to null / blank when they succeed at logging in
	LastFailedLogin *time.Time

	// each time there is a failed login, **increment** this
	// set to zero after successful login
	FailedLoginCount int

	// calculated each time there is a failed login
	NextAllowedLogin *time.Time

	TermsAccepted   bool
	AcceptedTermsID string // foreign key to version of terms they accepted
}

type AuthCredentials struct {
	User *User

	RefreshToken string
	IDToken      string
	ExpiresIn    time.Time
}

// UserPIN is used to store users' PINs and their entire change history.
type UserPIN struct {
	UserID string // TODO: At the DB, this should be indexed

	HashedPIN string
	ValidFrom time.Time
	ValidTo   time.Time

	// TODO: Compute this each time an operation involving the PIN is carried out
	// 	in order to make routine things e.g login via PIN fast
	IsValid bool // TODO: Consider a composite or partial DB index with UserID, IsValid, flavour
	Flavour string
}

type IUserInvite interface {

	// TODO: flavour is an enum; client or pro app
	// TODO: send invite link via e.g SMS
	//    the invite deep link: opens the app if installed OR goes to the store if not installed
	//    a first time PIN is set and sent to the user
	//    this PIN must be changed on first use
	//    this PIN can be used only once
	//	  **encode** first use PIN and user ID into invite link
	//	  i.e not a generic invite link
	// TODO: generate first time PIN, must change, link to user
	// TODO: set the PIN valid to to the current moment so that the user is forced to change upon login
	// TODO determine communication channel for invite (e.g SMS) from settings
	Invite(userID string, flavour string) (bool, error)
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
	Forget(userID string, pin string, flavour string) (bool, error)
}

// IRequestDataExport allows a user to request data known about them
// Mostly for legal compliance.
// The first impl. will simply create a task (for manual follow up) and acknowledge
type IRequestDataExport interface {
	RequestDataExport(userID string, pin string, flavour string) (bool, error)
}

type ISetUserPIN interface {
	// SetUserPIN sets a user's PIN.
	// It can be used to set a PIN for the first time.
	// It can be used to change the PIN.
	// It can also be used to change a PIN e.g on first login after invite or
	// after expiry.
	// TODO: auditable
	// TODO: Consult CLIENT_PIN_VALIDITY_DAYS and PRO_PIN_VALIDITY DAYS env/setting to set expiry
	// TODO: flavour is an enum...same enum used in profile e.g Client, Pro
	// TODO: ensure that old PINs are not re-used
	//	this presumes that we keep a record of **hashed** PINs per user
	// TODO Each time a PIN is set, recalculate valid to / valid from and update the
	//	cached IsActive value as appropriate i.e latest PIN active, others inactive
	//
	// PINs should not be re-used (compare hashed PINs)
	// TODO: the user pin table has validity and each new PIN that is set should be a new
	// entry in the table; and also invalidate past PINs.
	// it means that the same table can be used to check for PIN reuse.
	// TODO: all PINs are hashed
	SetUserPIN(userID string, pin string, confirm string, flavour string) (bool, error)
}

type ResetPIN interface {
	// ResetPIN can be used by admins or healthcare workers to generate and send
	// a new PIN for a client or other user.
	// The new PIN is generated automatically and set to expire immediately so
	// that a PIN change is forced on next login.
	// TODO: Notify user after PIN reset
	ResetPIN(userID string, flavour string) (bool, error)
}

// IVerifyPIN is used e.g to check the PIN when accessing sensitive content
type IVerifyPIN interface {
	VerifyPIN(userID string, flavour string, pin string) (bool, error)
}

type ILogin interface {
	// ...
	// when successful: return the user object
	// when not successful: nil user, error **code**, error
	// error codes should be standardized (enum)
	// the second param: intended for the clients (mobile, web) to understand
	// the third param: a technical error that can be handled in Go e.g logged
	// TODO: After verifying PIN, check PIN valid to
	//	if in future; allow login
	// 	if in past; require change
	//	require change: communicate to mobile/web client via error code (second return value)
	//  ONLY create access token/cookie etc AFTER all checks pass
	// TODO: error codes (second param) need to be a controlled list (enum) that is
	// 	synchronized between the frontend clients, Go code and GraphQL schema.
	//	it needs to be discussed by mobile + backend devs together.
	// TODO Only allow active users to log in
	// TODO For successful logins, reset last failed login and failed login count; set last successful login
	// TODO For failed logins:
	//	increment failed login count
	//	update last failed login timestamp
	//	set next allowed login timestamp
	//	use the failed login count (post increment) as the exponent to calculate the duration/interval
	//		to add in order to get the next allowed login timestamp
	//	the base (for the exponential backoff calculation) is a setting (env + default)
	//	default this base to 4...but override to 3 for a start in env
	// TODO: Only users who have accepted terms can login
	// TODO: Update metrics e.g login count, failed login count, successful login count etc
	Login(userID string, pin string, flavour string) (*AuthCredentials, string, error)
}

type IReviewTerms interface {

	// ReviewTerms can be used to accept or review terms
	ReviewTerms(userID string, accepted bool, flavour string) (bool, error)
}

type IAnonymizedIdentifier interface {
	// GetAnonymizedUserIdentifier is used to get an opaque (but **stable**) user
	// identifier for events, analytics etc
	GetAnonymizedUserIdentifier(userID string, flavour string) (string, error)
}

type IAddPushToken interface {
	AddPushtoken(userID string, flavour string) (bool, error)
}

type IRemovePushtoken interface {
	RemovePushToken(userID string, flavour string) (bool, error)
}

type IUpdateLanguagePreferences interface {
	UpdateLanguagePreferences(userID string, language string) (bool, error)
}

type UserUseCases interface {
	IUserInvite
	IUserForget
	ISetUserPIN
	ILogin
	IRequestDataExport
	IReviewTerms
	IAnonymizedIdentifier
	IAddPushToken
	IRemovePushtoken
	IUpdateLanguagePreferences
}

// TODO: CRUD for users...including search
// TODO: do not implement create user...if abused, we'd end up with a situation where there is a user but no profile
// search for users
// filtering user lists: e.g by criteria such as age (needed e.g to find survey cohorts)
