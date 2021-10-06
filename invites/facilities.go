package invites

// Facility models the details of healthcare facilities that are on the platform.
//
// e.g CCC clinics, Pharmacies.
//
type Facility struct {
	ID          *string // globally unique when set
	Name        string  // unique within this structure
	Code        string  // MFL Code for Kenyan facilities, globally unique
	Active      bool
	County      string // TODO: Controlled list of counties
	Description string
}

type FacilityPage struct {
	Facilities   []*Facility
	Count        int
	CurrentPage  int
	NextPage     *int
	PreviousPage *int
}

type FilterParam struct {
	Name     string
	DataType string // TODO: Ideally a controlled list i.e enum
	Data     string // TODO: Clear spec on validation e.g dates must be ISO 8601
}

func (f FilterParam) Validate() error {
	// TODO: Very strict validation of data <-> data type
	// 	     this is a good candidate for TDD with unit tests
	// TODO: make sure this is always called before filter params are used
	return nil
}

// TODO Read items per page from settings

type FacilityUseCases interface {
	List(
		// search
		searchTerm *string,

		// filter
		filter []*FilterParam,

		// paginate
		page int,
	) (*FacilityPage, error)

	// TODO Ensure blank ID when creating
	Create(facility Facility) (*Facility, error)
}
