package model

type Registration struct {
	Name     string `json:"name" example:"Username"`
	Password string `json:"password" example:"PaSsWoRd"`
	Confirm  string `json:"confirm" example:"PaSsWoRd"`
}
