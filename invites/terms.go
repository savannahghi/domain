package invites

import "time"

type Terms struct {
	ID        string
	Text      string
	Flavour   string
	ValidFrom time.Time
	ValidTo   time.Time
}

type IGetCurrentTerms interface {
	GetCurrentTerms(flavour string) (string, error)
}

type TermsUseCases interface {
	IGetCurrentTerms
}

// TODO Crud for terms
