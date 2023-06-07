package models

type Comment struct {
	By     string `json:"by,omitempty"`
	ID     int    `json:"id,omitempty"`
	Kids   []int  `json:"kids,omitempty"`
	Parent int    `json:"parent,omitempty"`
	Text   string `json:"text,omitempty"`
	Time   int    `json:"time,omitempty"`
	Type   string `json:"type,omitempty"`
}
