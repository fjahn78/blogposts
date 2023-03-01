package blogposts_test

import (
	"blogposts"
	"testing"
	"testing/fstest"
)

func TestBlogPosts(t *testing.T) {
  // Given
  fs := fstest.MapFS{
    "hello-world.md": {Data: []byte("hello, world!")},
    "hello-world2.md": {Data: []byte("hello, world2!")},
  }

  // When
  posts := blogposts.PostFromFS(fs)

  // Then
  if len(posts)!= len(fs) {
    t.Errorf("Expected %d posts, got %d", len(fs), len(posts))
  }
}