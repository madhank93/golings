package blog_test

import (
	"testing"

	"example.com/blog"
)

func TestPermalink(t *testing.T) {
	if got, want := blog.Permalink("Hello  Modules World"), "/posts/hello-modules-world"; got != want {
		t.Errorf("Permalink = %q, want %q", got, want)
	}
}
