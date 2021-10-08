package invites

import "time"

// TODO: define primary OTP channel (enum) in settings/env
// TODO: fallback channels in settings, in priority order
// TODO: setting/env to define max. number of OTP resends

type OTP struct {
	ID string // globally unique ID

	UserID      string // foreign key to user
	IsValid     bool   // set to false after it's used or invalidated for any reason
	GeneratedAt time.Time
	ValidUntil  time.Time
	Channel     string // TODO enum: channel the OTP was sent via
	Flavour     string // TODO enum
}

// TODO: behavior: verify...collect metrics
//	consider IsValid and ValidUntil
//	invalidate after verification i.e use once, invalidate all active OTPs after secure authentication

type IGenerateOTP interface {
	// TODO: ensure generated OTP is valid e.g valid until > generated at
	// metrics
	GenerateOTP(userID string, flavour string) (string, error)
}

type ISendOTP interface {
	// delegate to GenerateOTP
	// clients should call: SendOTP
	// send on the primary channel
	// metrics
	// the middle parameter is an error code e.g if rate limited
	SendOTP(userID string, flavour string) (bool, string, error)
}

type IResendOTP interface {
	// send via the next channel in priority
	// metrics
	// the middle parameter is an error code e.g if rate limited
	// TODO: consider limiting re-sends e.g
	ReSendOTP(userID string, flavour string) (bool, string, error)
}

type IVerifyOTP interface {
	// TODO: only accept valid OTPs (single use)
	// TODO: only accept unexpired OTPs
	// TODO: invalidate the currently active OTPs after successful verification (bulk)
	VerifyOTP(userID string, flavour string, otp string) (bool, error)
}

type OTPUseCases interface {
	IGenerateOTP
	ISendOTP
	IResendOTP
	IVerifyOTP
}
