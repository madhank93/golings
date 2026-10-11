// modules11
// This module vendors its dependencies: vendor/ holds a copy of every module
// it needs, and the build reads from there instead of the module cache.
// Someone added the slug dependency to go.mod without refreshing vendor/, and
// the build now refuses to run. Bring vendor/ back in line with go.mod.

// I AM NOT DONE
package blog

import "example.com/slug"

// Permalink is the path a post is published under.
func Permalink(title string) string {
	return "/posts/" + slug.Make(title)
}
