package invites

import "time"

type SecurityQuestion struct {
	ID           string
	QuestionStem string
	Description  string // help text
	ResponseType string // TODO: Enum
	Flavour      string // TODO: Enum
	Active       bool
	Sequence     *int // for sorting
}

type SecurityQuestionResponse struct {
	ID string

	TimeStamp          time.Time
	UserID             string // foreign key to question
	SecurityQuestionID string // foreign key to question
	Response           string // TODO: ensure we can encode/decode different response types
	Flavour            string
}

type RecordSecurityQuestionResponse struct {
	SecurityQuestionID string
	IsCorrect          bool
}

type IGetSecurityQuestions interface {
	// TODO: try to give the same user a stable set of questions (order)
	//
	GetSecurityQuestions(userID string, flavour string, n int) ([]*SecurityQuestion, error)
}

type IRecordSecurityQuestionResponses interface {

	// TODO: Validate the responses...all fields in the struct are required
	// get userID from responses
	// infer flavour and question from responses
	// TODO Save responses for the user...for use in future verification
	// TODO Wire in metrics
	RecordSecurityQuestionResponses(
		responses []*SecurityQuestionResponse,
	) ([]*RecordSecurityQuestionResponse, error)
}

type IVerifySecurityQuestionResponses interface {

	// TODO: get user ID and question from responses
	// TODO: look up saved responses for the user and compare...all must match
	// TODO: wire in metrics
	VerifySecurityQuestionResponses(
		responses []*SecurityQuestionResponse,
	) (bool, error)
}

type SecurityQuestionUseCases interface {
	IGetSecurityQuestions
	IRecordSecurityQuestionResponses
	IVerifySecurityQuestionResponses
}

// TODO CRUD for questions, including inactivate/reactivate
