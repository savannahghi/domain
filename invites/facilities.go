package invites

// Facility models the details of healthcare facilities that are on the platform.
//
// e.g CCC clinics, Pharmacies.
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
	DataType string // TODO: Ideally a controlled list i.e enum (MFL code, Active, County )
	Value    string // TODO: Clear spec on validation e.g dates must be ISO 8601. This is the actual data being filtered
}

func (f FilterParam) Validate() error {
	// Validate enums
	// TODO: Very strict validation of data <-> data type
	// 	     this is a good candidate for TDD with unit tests
	// TODO: make sure this is always called before filter params are used
	return nil
}

// TODO Read items per page from settings

type IFacilityList interface {
	// TODO Document: callers should specify active
	List(
		// search
		searchTerm *string,

		// filter
		filter []*FilterParam,

		// paginate
		page int,
	) (*FacilityPage, error)
}

type IFacilityRetrieve interface {
	Retrieve(id string) (*Facility, error)
}

type IFacilityCreate interface {
	// TODO Ensure blank ID when creating
	// TODO Since `id` is optional, ensure pre-condition check
	Create(facility Facility) (*Facility, error)
}

type IFacilityUpdate interface {
	// TODO Pre-condition: ensure `id` is set and valid
	Update(facility Facility) (*Facility, error)
}

type IFacilityDelete interface {
	// TODO Ensure delete is idempotent
	Delete(id string) (bool, error)
}

type IFacilityInactivate interface {
	// TODO Toggle active boolean
	Inactivate(id string) (*Facility, error)
}

type IFacilityReactivate interface {
	Reactivate(id string) (*Facility, error)
}

type FacilityUseCases interface {
	IFacilityList
	IFacilityRetrieve
	IFacilityCreate
	IFacilityUpdate
	IFacilityDelete
	IFacilityInactivate
	IFacilityReactivate
}
