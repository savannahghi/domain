package clinical

// MedicationStatementStatusEnum indicates the status of a medication statement
type MedicationStatementStatusEnum string

const (
	// MedicationStatementStatusEnumActive is The medication is still being taken.
	MedicationStatementStatusEnumActive MedicationStatementStatusEnum = "active"

	// MedicationStatementStatusEnumInActive is The medication is no longer being taken.
	MedicationStatementStatusEnumInActive MedicationStatementStatusEnum = "inactive"

	// MedicationStatementStatusEnumEnteredInError is Some of the actions that are implied by the medication statement may have occurred.
	MedicationStatementStatusEnumEnteredInError MedicationStatementStatusEnum = "entered-in-error"

	// MedicationStatementStatusEnumIntended is The medication may be taken at some time in the future.
	MedicationStatementStatusEnumIntended MedicationStatementStatusEnum = "intended"

	// MedicationStatementStatusEnumStopped is Actions implied by the statement have been permanently halted, before all of them occurred
	MedicationStatementStatusEnumStopped MedicationStatementStatusEnum = "stopped"

	// MedicationStatementStatusEnumOnHold is Actions implied by the statement have been temporarily halted, but are expected to continue later.
	MedicationStatementStatusEnumOnHold MedicationStatementStatusEnum = "on-hold"

	// MedicationStatementStatusEnumUnknown is The state of the medication use is not currently known.
	MedicationStatementStatusEnumUnknown MedicationStatementStatusEnum = "unknown"

	// MedicationStatementStatusEnumNotTaken is The medication was not consumed by the patient
	MedicationStatementStatusEnumNotTaken MedicationStatementStatusEnum = "not-taken"
)

type MedicationStatementInput struct {
	Status *MedicationStatementStatusEnum `json:"status,omitempty"`

	MedicationCode string `json:"medicationCode,omitempty"`

	PatientID string `json:"patientID,omitempty"`
}

type MedicationStatement struct {
	ID     string                         `json:"id,omitempty"`
	Status *MedicationStatementStatusEnum `json:"status,omitempty"`

	MedicationCode string `json:"medicationCode,omitempty"`

	PatientID string `json:"patientID,omitempty"`
}
