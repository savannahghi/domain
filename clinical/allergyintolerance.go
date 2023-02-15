package clinical

import (
	"time"
)

// AllergyIntoleranceCriticalityEnum is a FHIR enum
type AllergyIntoleranceCriticalityEnum string

const (
	// AllergyIntoleranceCriticalityEnumLow ...
	AllergyIntoleranceCriticalityEnumLow AllergyIntoleranceCriticalityEnum = "low"
	// AllergyIntoleranceCriticalityEnumHigh ...
	AllergyIntoleranceCriticalityEnumHigh AllergyIntoleranceCriticalityEnum = "high"
	// AllergyIntoleranceCriticalityEnumUnableToAssess ...
	AllergyIntoleranceCriticalityEnumUnableToAssess AllergyIntoleranceCriticalityEnum = "unable_to_assess"
)

// AllergyIntoleranceReactionSeverityEnum is a FHIR enum
type AllergyIntoleranceReactionSeverityEnum string

const (
	// AllergyIntoleranceReactionSeverityEnumMild ...
	AllergyIntoleranceReactionSeverityEnumMild AllergyIntoleranceReactionSeverityEnum = "mild"
	// AllergyIntoleranceReactionSeverityEnumModerate ...
	AllergyIntoleranceReactionSeverityEnumModerate AllergyIntoleranceReactionSeverityEnum = "moderate"
	// AllergyIntoleranceReactionSeverityEnumSevere ...
	AllergyIntoleranceReactionSeverityEnumSevere AllergyIntoleranceReactionSeverityEnum = "severe"
)

type FHIRAllergyIntoleranceInput struct {
	PatientID string `json:"patientID,omitempty"`

	// Estimate of the potential clinical harm, or seriousness, of the reaction to the identified substance.
	Criticality AllergyIntoleranceCriticalityEnum `json:"criticality,omitempty"`

	// The encounter when the allergy or intolerance was asserted.
	EncounterID string `json:"encounterID,omitempty"`

	// Estimated or actual date,  date-time, or age when allergy or intolerance was identified.
	OnsetDateTime *time.Time `json:"onsetDateTime,omitempty"`

	Severity *AllergyIntoleranceReactionSeverityEnum `json:"severity,omitempty"`

	// Identification of the specific substance (or pharmaceutical product) considered to be responsible for the Adverse Reaction event.
	// Note: the substance for a specific reaction may be different from the substance identified as the cause of the risk, but it must be consistent with it.
	// For instance, it may be a more specific substance (e.g. a brand medication) or a composite product that includes the identified substance. It must be clinically safe to only process the 'code' and ignore the 'reaction.substance'.  If a receiving system is unable to confirm that AllergyIntolerance.reaction.substance falls within the semantic scope of AllergyIntolerance.code, then the receiving system should ignore AllergyIntolerance.reaction.substance.
	SubstanceCode int `json:"substanceCode,omitempty"`
}

type FHIRAllergyIntolerance struct {
	ID *string `json:"id,omitempty"`

	PatientID string `json:"patientID,omitempty"`

	// Estimate of the potential clinical harm, or seriousness, of the reaction to the identified substance.
	Criticality AllergyIntoleranceCriticalityEnum `json:"criticality,omitempty"`

	// The encounter when the allergy or intolerance was asserted.
	EncounterID string `json:"encounterID,omitempty"`

	// Estimated or actual date,  date-time, or age when allergy or intolerance was identified.
	OnsetDateTime *time.Time `json:"onsetDateTime,omitempty"`

	Severity *AllergyIntoleranceReactionSeverityEnum `json:"severity,omitempty"`

	// Identification of the specific substance (or pharmaceutical product) considered to be responsible for the Adverse Reaction event.
	// Note: the substance for a specific reaction may be different from the substance identified as the cause of the risk, but it must be consistent with it.
	// For instance, it may be a more specific substance (e.g. a brand medication) or a composite product that includes the identified substance. It must be clinically safe to only process the 'code' and ignore the 'reaction.substance'.  If a receiving system is unable to confirm that AllergyIntolerance.reaction.substance falls within the semantic scope of AllergyIntolerance.code, then the receiving system should ignore AllergyIntolerance.reaction.substance.
	SubstanceCode int `json:"substanceCode,omitempty"`
}
