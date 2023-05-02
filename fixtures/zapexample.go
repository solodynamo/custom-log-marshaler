package fixtures

type User struct {
	Name    *string `json:"name"`
	Email   string  `notloggable`
	Address string  `json:"address", notloggable`
}

type Translation struct {
	Language    string `json:"language,omitempty"`
	Translation string `json:"translation"`
}

type customType map[string]string
type number int

type UserDetailsResponse struct {
	User
	FromCache    bool          `json:"fromCache"`
	Translations []Translation `json:"translations"`
	Metadata     customType    `json:"metadata"`
	No           number        `json:"no"`
}
