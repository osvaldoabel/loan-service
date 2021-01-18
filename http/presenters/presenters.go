package presenters

// UserPayload -
type UserPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Birthday string `json:"birhday"`
	Status   string `json:"status"`
}

// UserCollection -
type UserCollection struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	Birthday  string `json:"birthday"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// DefaultQueryString -
type DefaultQueryString struct {
	Limit  string `json:"per_page"`
	Offset string `json:"page"`
}
