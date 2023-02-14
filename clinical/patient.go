package clinical

type SimplePatientRegistrationInput struct {
	FirstName               string                    `json:"firstName"`
	LastName                string                    `json:"lastName"`
	OtherNames              *string                   `json:"otherNames"`
	IdentificationDocuments []*IdentificationDocument `json:"identificationDocuments"`
	BirthDate               string                    `json:"birthDate"`
	PhoneNumbers            []string                  `json:"phoneNumbers"`
	Gender                  string                    `json:"gender"`
}

type IdentificationDocument struct {
	DocumentType   IDDocumentType `json:"documentType"`
	DocumentNumber string         `json:"documentNumber"`
}

// IDDocumentType is an internal code system for identification document types.
type IDDocumentType string

// ID type constants
const (
	// IDDocumentTypeNationalID ...
	IDDocumentTypeNationalID IDDocumentType = "national_id"
	// IDDocumentTypePassport ...
	IDDocumentTypePassport IDDocumentType = "passport"
	// IDDocumentTypeAlienID ...
	IDDocumentTypeAlienID IDDocumentType = "alien_id"
)

// AllIDDocumentType is a list of known ID types
var AllIDDocumentType = []IDDocumentType{
	IDDocumentTypeNationalID,
	IDDocumentTypePassport,
	IDDocumentTypeAlienID,
}

// PatientGenderEnum is a FHIR enum
type PatientGenderEnum string

const (
	// PatientGenderEnumMale ...
	PatientGenderEnumMale PatientGenderEnum = "male"
	// PatientGenderEnumFemale ...
	PatientGenderEnumFemale PatientGenderEnum = "female"
	// PatientGenderEnumOther ...
	PatientGenderEnumOther PatientGenderEnum = "other"
	// PatientGenderEnumUnknown ...
	PatientGenderEnumUnknown PatientGenderEnum = "unknown"
)

// FHIRPatient definition: demographics and other administrative information about an individual or animal receiving care or other health-related services.
type FHIRPatient struct {
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	ID string `json:"id,omitempty"`

	// Whether this patient record is in active use.
	// Many systems use this property to mark as non-current patients, such as those that have not been seen for a period of time based on an organization's business rules.
	//
	// It is often used to filter patient lists to exclude inactive patients
	//
	// Deceased patients may also be marked as inactive for the same reasons, but may be active for some time after death.
	Active bool `json:"active,omitempty"`

	// A name associated with the individual.
	Name string `json:"name,omitempty"`

	// A contact detail (e.g. a telephone number or an email address) by which the individual may be contacted.
	PhoneNumber []string `json:"phoneNumber,omitempty"`

	// Administrative Gender - the gender that the patient is considered to have for administration and record keeping purposes.
	Gender *PatientGenderEnum `json:"gender,omitempty"`

	// The date of birth for the individual.
	BirthDate string `json:"birthDate,omitempty"`
}

type PatientPayload struct {
	PatientRecord   *FHIRPatient `json:"patientRecord,omitempty"`
	HasOpenEpisodes bool         `json:"hasOpenEpisodes,omitempty"`
}
