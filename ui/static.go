package ui

import (
	"embed"
	"io/fs"
)

type StaticResource struct {
	FS embed.FS
}

//go:embed dist
var Dist embed.FS

func (s *StaticResource) Open(name string) (fs.File, error) {
	return s.FS.Open(name)
}
