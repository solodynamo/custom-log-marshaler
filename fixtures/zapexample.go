package fixtures

type User struct {
	Name    *string `json:"name"`
	Email   string  `notloggable`
	Address string  `json:"address", notloggable`
}

type UserDetailsResponse struct {
	User
	RequestID string   `json:"rid"`
	FromCache bool     `json:"fromCache"`
	Metadata  []string `json:"md"`
}
