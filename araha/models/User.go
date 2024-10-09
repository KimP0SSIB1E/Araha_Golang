package models

type User struct {
	ID        string `json:"ID"`
	UserName  string `json:"UserName"`
	Email     string `json:"Email"`
	Password  string `json:"Password"`
	UserToken string `json:"UserToken"`
	Balance   float64
}
