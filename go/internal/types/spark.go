package types

import "time"

type WebSpark struct {
	ID       string   `json:"id"`
	Title    string   `json:"title"`
	Markdown string   `json:"markdown,omitempty"`
	Slug     string   `json:"slug,omitempty"`
	UpdatedAt time.Time `json:"updatedAt"`
	Tags     []string `json:"tags,omitempty"`
}

type SparkUpdateRequest struct {
	Title    *string  `json:"title,omitempty"`
	Markdown *string  `json:"markdown,omitempty"`
	Slug     *string  `json:"slug,omitempty"`
	Tags     []string `json:"tags,omitempty"`
}
