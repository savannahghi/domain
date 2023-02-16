package clinical

import (
	"time"
)

// ConditionInput is the input type for Condition
type ConditionInput struct {
	PatientID string `json:"patientID,omitempty"`

	EncounterID string `json:"encounterID,omitempty"`

	OnsetDate *time.Time `json:"onsetDate,omitempty"`

	Note string `json:"note,omitempty"`

	// Will go to the `CODE` struct field
	Condition string `json:"condition,omitempty"`
}

// Condition definition: a clinical condition, problem, diagnosis, or other event, situation, issue, or clinical concept that has risen to a level of concern.
type Condition struct {
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	ID *string `json:"id,omitempty"`

	RecordedDate *time.Time `json:"recordedDate,omitempty"`

	PatientID string `json:"patientID,omitempty"`

	Condition string `json:"condition,omitempty"`

	EncounterID string `json:"encounterID,omitempty"`

	OnsetDate *time.Time `json:"onsetDate,omitempty"`

	Note string `json:"note,omitempty"`
}
