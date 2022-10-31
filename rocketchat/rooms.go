package rocketchat

import "time"

// https://developer.rocket.chat/reference/api/schema-definition

type Room struct {
	ID           string       `json:"_id"`
	T            string       `json:"t"`
	Ts           time.Time    `json:"ts"`
	Name         string       `json:"name"`
	Lm           time.Time    `json:"lm"`
	Msgs         int          `json:"msgs"`
	Cl           bool         `json:"cl"`
	Ro           bool         `json:"ro"`
	Usernames    []string     `json:"usernames"`
	U            User         `json:"u"`
	CustomFields CustomFields `json:"customFields"`
}

type UserObject3 struct {
	A string `json:"a"`
	B string `json:"b"`
}

type CustomFields struct {
	UserDefinedField  string `json:"userDefinedField"`
	UserDefinedField2 bool   `json:"userDefinedField2"`
}
