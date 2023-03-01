package blogposts_test

import (
	"blogposts"
	"testing"
	"testing/fstest"
)

func TestBlogPosts(t *testing.T) {
  // Given
  fs := fstest.MapFS{
    "hello-world.md": {Data: []byte("Title: hello, world!")},
    "hello-world2.md": {Data: []byte("Title: Hola, mundo!")},
  }

  // When
  posts := blogposts.PostFromFS(fs)

  // Then
  if len(posts)!= len(fs) {
    t.Errorf("Expected %d posts, got %d", len(fs), len(posts))
  }

  want := blogposts.Post{Title:"hello, world!"}

  if posts[0] != want {
    t.Errorf("Expected %#v, got %#v", want, posts[0])
  }
}