// Package text check if given path is text file
package text // import "github.com/shamsher31/goistext"

import (
	"github.com/shamsher31/gotextext"
	"path"
	"strings"
)

// Call Get from textext package
var extensions = textext.Get()

// Extensions is the extensions for different text file
var Extensions map[string]bool

func init() {
	Extensions = map[string]bool{}
	for _, ext := range extensions {
		Extensions[ext] = true
	}
}

// Is checks if a path is a text file
func Is(p string) bool {
	ext := strings.TrimLeft(path.Ext(p), ".")
	return Extensions[strings.ToLower(ext)]
}
