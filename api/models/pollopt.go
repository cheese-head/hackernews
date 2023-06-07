package models

type Pollopt struct {
	By    string `json:"by,omitempty"`
	ID    int    `json:"id,omitempty"`
	Poll  int    `json:"poll,omitempty"`
	Score int    `json:"score,omitempty"`
	Text  string `json:"text,omitempty"`
	Time  int    `json:"time,omitempty"`
	Type  string `json:"type,omitempty"`
}
