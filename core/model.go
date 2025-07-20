package core

import (
	"fmt"
	"strings"
)

// Snippet represents a code snippet extracted from a source file
type Snippet struct {
	ID       int      `json:"id"`
	Title    string   `json:"title"`
	Tags     []string `json:"tags"`
	Language string   `json:"language"`
	Content  string   `json:"content"`
	Source   string   `json:"source"` // e.g. ./main.go:23
}

// NewSnippet creates a new Snippet from parsed values
func NewSnippet(id int, title, lang, content, source string, tags []string) Snippet {
	return Snippet{
		ID:       id,
		Title:    title,
		Tags:     tags,
		Language: lang,
		Content:  content,
		Source:   source,
	}
}

// ParseTags splits a tag string like "go,web,auth" into []string{"go", "web", "auth"}
func ParseTags(tagLine string) []string {
	tags := strings.Split(tagLine, ",")
	var clean []string
	for _, tag := range tags {
		tag = strings.TrimSpace(tag)
		if tag != "" {
			clean = append(clean, tag)
		}
	}
	return clean
}

// Display prints a short formatted summary (optional)
func (s Snippet) Display() string {
	return fmt.Sprintf("[%d] %s (%s)  [%s]", s.ID, s.Title, s.Language, strings.Join(s.Tags, ", "))
}
