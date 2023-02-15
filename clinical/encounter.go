package clinical

// EncounterStatusEnum is a FHIR enum
type EncounterStatusEnum string

const (
	// EncounterStatusEnumPlanned ...
	EncounterStatusEnumPlanned EncounterStatusEnum = "planned"
	// EncounterStatusEnumArrived ...
	EncounterStatusEnumArrived EncounterStatusEnum = "arrived"
	// EncounterStatusEnumTriaged ...
	EncounterStatusEnumTriaged EncounterStatusEnum = "triaged"
	// EncounterStatusEnumInProgress ...
	EncounterStatusEnumInProgress EncounterStatusEnum = "in_progress"
	// EncounterStatusEnumOnleave ...
	EncounterStatusEnumOnleave EncounterStatusEnum = "onleave"
	// EncounterStatusEnumFinished ...
	EncounterStatusEnumFinished EncounterStatusEnum = "finished"
	// EncounterStatusEnumCancelled ...
	EncounterStatusEnumCancelled EncounterStatusEnum = "cancelled"
	// EncounterStatusEnumEnteredInError ...
	EncounterStatusEnumEnteredInError EncounterStatusEnum = "entered_in_error"
	// EncounterStatusEnumUnknown ...
	EncounterStatusEnumUnknown EncounterStatusEnum = "unknown"
)

type EncounterClass string

const (
	// Also referred to as outpatient
	EncounterClassAmbulatory         EncounterClass = "ambulatory"
	EncounterClassEmergency          EncounterClass = "emergency"
	EncounterClassField              EncounterClass = "field"
	EncounterClassHomeHealth         EncounterClass = "home health"
	EncounterClassInpatientEncounter EncounterClass = "inpatient encounter"
	EncounterCodeInpatientAcute      EncounterClass = "inpatient acute"
	EncounterCodeInpatientNonAcute   EncounterClass = "inpatient non-acute"
	EncounterCodePreOp               EncounterClass = "pre-op"
	EncounterCodeShortStay           EncounterClass = "short stay"
	EncounterCodeVirtual             EncounterClass = "virtual"
)

// FHIREncounterInput is the input type for Encounter
type FHIREncounterInput struct {
	// planned | arrived | triaged | in-progress | onleave | finished | cancelled +.
	Status EncounterStatusEnum `json:"status,omitempty"`

	// Concepts representing classification of patient encounter such as ambulatory (outpatient), inpatient, emergency, home health or others due to local variations.
	// This is a coding input --> Note this when mapping to FHIR
	Class EncounterClass `json:"class,omitempty"`

	PatientID string `json:"patientID,omitempty"`

	// EpisodeOfCareID Episode(s) of care that this encounter should be recorded against
	EpisodeOfCareID string `json:"episodeOfCareID,omitempty"`
}

// FHIREncounter definition: an interaction between a patient and healthcare provider(s) for the purpose of providing healthcare service(s) or assessing the health status of a patient.
type FHIREncounter struct {
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	ID *string `json:"id,omitempty"`

	// planned | arrived | triaged | in-progress | onleave | finished | cancelled +.
	Status EncounterStatusEnum `json:"status,omitempty"`

	// Concepts representing classification of patient encounter such as ambulatory (outpatient), inpatient, emergency, home health or others due to local variations.
	// This is a coding input --> Note this when mapping to FHIR
	Class EncounterClass `json:"class,omitempty"`

	PatientID string `json:"patientID,omitempty"`

	// EpisodeOfCareID Episode(s) of care that this encounter should be recorded against
	EpisodeOfCareID string `json:"episodeOfCareID,omitempty"`
}
