package types

type WebSpark struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Markdown string `json:"markdown,omitempty"`
}

type SparkUpdateRequest struct {
	Title    *string `json:"title,omitempty"`
	Markdown *string `json:"markdown,omitempty"`
}
