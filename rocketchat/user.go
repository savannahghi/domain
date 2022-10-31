package rocketchat

type User struct {
	ID           string `json:"_id"`
	Name         string `json:"name"`
	UserName     string `json:"username"`
	Status       string `json:"status"`
	Token        string `json:"token"`
	TokenExpires int64  `json:"tokenExpires"`
}

type CreateUserRequest struct {
	Name                  string            `json:"name"`
	Email                 string            `json:"email"`
	Password              string            `json:"password"`
	Username              string            `json:"username"`
	Roles                 []string          `json:"roles,omitempty"`
	Active                bool              `json:"active"`
	JoinDefaultChannels   bool              `json:"joinDefaultChannels"`
	RequirePasswordChange bool              `json:"requirePasswordChange"`
	SendWelcomeEmail      bool              `json:"sendWelcomeEmail"`
	Verified              bool              `json:"verified"`
	CustomFields          map[string]string `json:"customFields,omitempty"`
}

type RegisterUserRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"pass"`
	Username  string `json:"username"`
	SecretURL string `json:"secretURL"`
}

type UpdateUserRequest struct {
	UserID string `json:"userId"`
	Data   struct {
		Name         string            `json:"name"`
		Email        string            `json:"email"`
		Password     string            `json:"password"`
		Username     string            `json:"username"`
		CustomFields map[string]string `json:"customFields,omitempty"`
	} `json:"data"`
}

type UsersList struct {
	ID        string `json:"_id"`
	Type      string `json:"type"`
	Status    string `json:"status"`
	Active    bool   `json:"active"`
	Name      string `json:"name"`
	Utcoffset bool   `json:"utcoffset"`
	Username  string `json:"username"`
}

type UserInfo struct {
	ID        string `json:"_id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Status    string `json:"status"`
	UtcOffset int    `json:"utcOffset"`
	Active    bool   `json:"active"`
	Type      string `json:"type"`
}

type Visitor struct {
	ID    string `json:"_id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Token string `json:"token"`
}
