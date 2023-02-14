package clinical

import (
	"time"
)

// ObservationStatusEnum is a FHIR enum
type ObservationStatusEnum string

const (
	// ObservationStatusEnumRegistered ...
	ObservationStatusEnumRegistered ObservationStatusEnum = "registered"
	// ObservationStatusEnumPreliminary ...
	ObservationStatusEnumPreliminary ObservationStatusEnum = "preliminary"
	// ObservationStatusEnumFinal ...
	ObservationStatusEnumFinal ObservationStatusEnum = "final"
	// ObservationStatusEnumAmended ...
	ObservationStatusEnumAmended ObservationStatusEnum = "amended"
	// ObservationStatusEnumCorrected ...
	ObservationStatusEnumCorrected ObservationStatusEnum = "corrected"
	// ObservationStatusEnumCancelled ...
	ObservationStatusEnumCancelled ObservationStatusEnum = "cancelled"
	// ObservationStatusEnumEnteredInError ...
	ObservationStatusEnumEnteredInError ObservationStatusEnum = "entered_in_error"
	// ObservationStatusEnumUnknown ...
	ObservationStatusEnumUnknown ObservationStatusEnum = "unknown"
)

type FHIRObservation struct {
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	ID *string `json:"id,omitempty"`

	Status *ObservationStatusEnum `json:"status,omitempty"`

	PatientID string `json:"patientID,omitempty"`

	EncounterID string `json:"encounterID,omitempty"`

	// The time or time-period the observed value is asserted as being true. For biological subjects - e.g. human patients - this is usually called the "physiologically relevant time". This is usually either the time of the procedure or of specimen collection, but very often the source of the date/time is not known, only the date/time itself.
	EffectiveDateTime *time.Time `json:"effectiveDateTime,omitempty"`

	// The information determined as a result of making the observation, if the information has a simple value.
	ObservationValue *string `json:"observationValue,omitempty"`

	ObservationType string `json:"observationType,omitempty"`
}

type FHIRObservationInput struct {
	Status *ObservationStatusEnum `json:"status,omitempty"`

	PatientID string `json:"patientID,omitempty"`

	EncounterID string `json:"encounterID,omitempty"`

	// The time or time-period the observed value is asserted as being true. For biological subjects - e.g. human patients - this is usually called the "physiologically relevant time". This is usually either the time of the procedure or of specimen collection, but very often the source of the date/time is not known, only the date/time itself.
	EffectiveDateTime *time.Time `json:"effectiveDateTime,omitempty"`

	// The information determined as a result of making the observation, if the information has a simple value.
	ObservationValue *string `json:"observationValue,omitempty"`

	ObservationType string `json:"observationType,omitempty"`
}
