package models

type Ask struct {
	By          string `json:"by,omitempty"`
	Descendants int    `json:"descendants,omitempty"`
	ID          int    `json:"id,omitempty"`
	Kids        []int  `json:"kids,omitempty"`
	Score       int    `json:"score,omitempty"`
	Text        string `json:"text,omitempty"`
	Time        int    `json:"time,omitempty"`
	Title       string `json:"title,omitempty"`
	Type        string `json:"type,omitempty"`
}
