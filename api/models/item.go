package models

type Item struct {
	By          string `json:"by,omitempty"`
	Dead        bool   `json:"dead,omitempty"`
	Deleted     bool   `json:"deleted,omitempty"`
	Descendants int    `json:"descendants,omitempty"`
	ID          int    `json:"id,omitempty"`
	Kids        []int  `json:"kids,omitempty"`
	Parent      int    `json:"parent,omitempty"`
	Parts       []int  `json:"parts,omitempty"`
	Poll        int    `json:"poll,omitempty"`
	Score       int    `json:"score,omitempty"`
	Text        string `json:"text,omitempty"`
	Time        int    `json:"time,omitempty"`
	Title       string `json:"title,omitempty"`
	Type        string `json:"type,omitempty"`
	URL         string `json:"url,omitempty"`
}
