package models

type Story struct {
	By          string `json:"by,omitempty"`
	Descendants int    `json:"descendants,omitempty"`
	ID          int    `json:"id,omitempty"`
	Kids        []int  `json:"kids,omitempty"`
	Score       int    `json:"score,omitempty"`
	Time        int    `json:"time,omitempty"`
	Title       string `json:"title,omitempty"`
	Type        string `json:"type,omitempty"`
	URL         string `json:"url,omitempty"`
}
