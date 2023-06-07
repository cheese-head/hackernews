package models

type Job struct {
	By    string `json:"by,omitempty"`
	ID    int    `json:"id,omitempty"`
	Score int    `json:"score,omitempty"`
	Text  string `json:"text,omitempty"`
	Time  int    `json:"time,omitempty"`
	Title string `json:"title,omitempty"`
	Type  string `json:"type,omitempty"`
	URL   string `json:"url,omitempty"`
}
