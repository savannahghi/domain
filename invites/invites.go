package invites

// facilities
// MFL codes
// names

// roles
// name
// permissions
// description

// metrics API
// data to collect
// behavior e.g add metric

// users
// nickname / username; patient can edit but we track
// ? explore twitter like username / handle + visible name, the latter can be edited
// names...? overlap with patients e.g if a patient is invited, can we prefill most of these
// type e.g HCW, client, admin etc
// user invite via email, SMS etc
// ? do we need admin role or we just work with roles and perms
// roles -> permissions
// add new user e.g super admin adding a HCW
// optional: staff number
// optional: link to facility(ies), for HCW
// email contacts
// phone contacts
// gender
// names...align | first, middle, other
// listing users
// search for users
// add role
// remove role
// could some users also require second factors e.g biometric or OTP?
// some users have expiring PINs e.g professionals

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

// get client from OpenMRS
// register client
// main identifier is CCC number (composite of MFL code and client code)
// other identifiers possible e.g phone number
// 3 names
// date of birth
// gender
// optional ID
// date of treatment enrollment (or estimate)
// calculate: length of treatment
// most important contacts: phone
// physical address
// county
// facility / clinic
// treatment buddy (optional)
// next of kin and relationship
// action: invite to client app
// action: add to group(s)
// search for a patient...search criteria
// inactivating and reactivating a client
// ? transfer out of our care ?

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
