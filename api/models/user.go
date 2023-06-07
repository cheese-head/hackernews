package models

type User struct {
	About     string `json:"about,omitempty"`
	Created   int    `json:"created,omitempty"`
	Delay     int    `json:"delay,omitempty"`
	ID        string `json:"id,omitempty"`
	Karma     int    `json:"karma,omitempty"`
	Submitted []int  `json:"submitted,omitempty"`
}
