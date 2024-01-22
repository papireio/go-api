package models

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Uuid     string `json:"uuid"`
	Verified bool   `json:"verified"`
}
