package invites

type Address struct {
	ID string

	Type       string // TODO: enum; postal, physical or both
	Text       string // actual address, can be multi-line
	Country    string // TODO: enum
	PostalCode string
	County     string // TODO: counties belong to a country
	Active     bool
}

// TODO: CRUD for this
