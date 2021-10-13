package invites

import "time"

// TODO: Model notifications behavior; consider source document checklist
// TODO: Consider surveys here too (that leaves only communities + health diary + timeline)

// e.g reschedule, add to calendar etc
type Action struct {
	ID string

	Name    string
	Payload map[string]string
	Icon    string // TODO link
}

type Notification struct {
	ID string

	Title            string
	Body             string   // TODO: might be **formatted** e.g MD
	Link             string   // TODO: not all notifications need this but for some (e.g surveys) this is the main thing
	Icon             string   // TODO: link
	Badge            []string // TODO: e.g to mark missed appointments
	Status           string   // TODO: enum e.g resolved, pending
	NotificationType string   // TODO: enum e.g appointment, survey, article; use to create "blind" display
	Channels         []string // TODO: enums e.g PUSH, SMS
	Timestamp        time.Time
	Actions          []*Action
	Flavour          string // TODO enum
}

type ISendSingleNotification interface {
	// TODO: validation e.g ID is set after sending
	SendSingleNotification(
		userID string,
		notification *Notification,
	) (*Notification, error)
}

type ISendGroupNotification interface {
	// TODO: validation e.g ID is set after sending
	SendGroupNotification(
		groupID string, // TODO revisit after discussing communities
		notification *Notification,
	) (*Notification, error)
}

type ISendSurveyNotifications interface {
	// TODO: the link should be trackable e.g take this link and replace with trackable links
	// we'll only track interactions on our platform e.g clicks
	SendSurvey(
		userIDs []string, // TODO: these could come from e.g group members, random selection, criteria etc
		title string,
		link string,
		channels []string,
		flavour string,
	) (bool, error)
}

type NotificationUsecases interface {
	ISendSingleNotification
	ISendGroupNotification
	ISendSurveyNotifications
}

// TODO: CRUD
// 	filters e.g by status, flavour
