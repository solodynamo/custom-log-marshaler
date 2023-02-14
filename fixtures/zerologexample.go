package fixtures

type User2 struct {
	Name    string `json:"name"`
	Email   string `notloggable`
	Address string `json:"address", notloggable`
}

type UserDetailsResponse2 struct {
	User
	RequestID string `json:"rid"`
	FromCache bool   `json:"fromCache"`
}
