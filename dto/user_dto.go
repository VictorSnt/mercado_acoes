package DTO

type CreateUser struct {
	Name    string
	Balance float64
}

type DisplayUser struct {
	ID      uint    `json:"user_id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

type UpdateUser struct {
	Name string `json:"name"`
}
