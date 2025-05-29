package internal

type Metadata struct {
	URL string `json:"url"`
}

type Document struct {
	Content  string   `json:"content"`
	Metadata Metadata `json:"metadata"`
}