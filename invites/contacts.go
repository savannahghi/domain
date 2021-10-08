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

// TODO: CRUD for contacts
