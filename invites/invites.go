package invites

// roles
// name
// permissions
// description

// terms and conditions (served from API)

// defer: list of important articles to show during onboarding

// group
// identity
// name
// can be listed
// topic
// type
// ...more shall be defined

// security questions
// randomized into sets of 3
// save both question and response
// verify security question response by user
// question has: stem, description, response type e.g text
// ? what is the overlap between this and questionnaires?
// metrics: e.g measure popularity of questions

// questionnaires
// questions
// types

// healthcare worker invites client
// consent - accept terms and conditions
// consent - HCW indicate they have counselled the patient
// contacts
// opt in - by channel
// language(s)
// client type e.g PMTCT, OVC
// generate and send invite link
// invite code
// ? can the invite code also be rendered as a QR code
// first time PIN, to be changed after first use
// PIN hashing
// link an invite to a phone number and CCC number
// invite link is a deep link
// send SMS after invite
// re-send invite link
// security questions
// accept invite
// ? should the invite expire
// welcome survey / questionnaire
// make sure push tokens are registered after sign up

// verify PIN
// PIN hashing
// max of n attempts
// exponential backoff
// lockout after 30 min (suspend)
// metrics: record number of attempts

// reset PIN
// OTP
// re-send OTPs...different channel
// secret questions
// set new PIN
// ? can this be done by HCWs
// ? notify after reset PIN
// metrics: record resets
// anticipate additional rules e.g max resets, PIN reuse etc
// PIN hashing
// if HCW sets PIN, it must be changed
