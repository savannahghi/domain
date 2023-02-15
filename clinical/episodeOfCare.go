package clinical

// EpisodeOfCareStatusEnum is a  enum
type EpisodeOfCareStatusEnum string

const (
	// EpisodeOfCareStatusEnumPlanned ...
	EpisodeOfCareStatusEnumPlanned EpisodeOfCareStatusEnum = "planned"
	// EpisodeOfCareStatusEnumWaitlist ...
	EpisodeOfCareStatusEnumWaitlist EpisodeOfCareStatusEnum = "waitlist"
	// EpisodeOfCareStatusEnumActive ...
	EpisodeOfCareStatusEnumActive EpisodeOfCareStatusEnum = "active"
	// EpisodeOfCareStatusEnumOnhold ...
	EpisodeOfCareStatusEnumOnhold EpisodeOfCareStatusEnum = "onhold"
	// EpisodeOfCareStatusEnumFinished ...
	EpisodeOfCareStatusEnumFinished EpisodeOfCareStatusEnum = "finished"
	// EpisodeOfCareStatusEnumCancelled ...
	EpisodeOfCareStatusEnumCancelled EpisodeOfCareStatusEnum = "cancelled"
	// EpisodeOfCareStatusEnumEnteredInError ...
	EpisodeOfCareStatusEnumEnteredInError EpisodeOfCareStatusEnum = "entered_in_error"
)

// EpisodeOfCare definition: an association between a patient and an organization / healthcare provider(s) during which time encounters may occur. the managing organization assumes a level of responsibility for the patient during this time.
type EpisodeOfCare struct {
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	ID *string `json:"id,omitempty"`
	// planned | waitlist | active | onhold | finished | cancelled.
	Status *EpisodeOfCareStatusEnum `json:"status,omitempty"`

	PatientID string `json:"patientID,omitempty"`

	// Will be extracted from the request headers
	OrganizationID string `json:"organizationID,omitempty"`
}

// EpisodeOfCareInput is the input type for EpisodeOfCare
type EpisodeOfCareInput struct {

	// planned | waitlist | active | onhold | finished | cancelled.
	Status *EpisodeOfCareStatusEnum `json:"status,omitempty"`

	PatientID string `json:"patientID,omitempty"`
}
