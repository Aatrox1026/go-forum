package model

type Registration struct {
	Name     string `json:"name" example:"Username"`
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"PaSsWoRd"`
	Confirm  string `json:"confirm" example:"PaSsWoRd"`
}

type Login struct {
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"PaSsWoRd"`
}
