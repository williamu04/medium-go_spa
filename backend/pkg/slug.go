package pkg

import "github.com/gosimple/slug"

type Sluger struct{}

func NewSluger() *Sluger {
	return &Sluger{}
}

func (s *Sluger) Slug(title string) string {
	return slug.Make(title)
}
