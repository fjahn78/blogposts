package blogposts_test

import (
	"blogposts"
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"
)

func TestBlogPosts(t *testing.T) {
  t.Run("happy path", func(t *testing.T) { 
    // Given
    fs := fstest.MapFS{
      "hello-world.md": {Data: []byte("Title: hello, world!")},
      "hello-world2.md": {Data: []byte("Title: Hola, mundo!")},
    }
    
    // When
    posts, err := blogposts.PostFromFS(fs)
    
    // Then
    if err != nil {
      t.Fatal(err)
    }
    
    if len(posts)!= len(fs) {
      t.Errorf("Expected %d posts, got %d", len(fs), len(posts))
  }
  
  want := blogposts.Post{Title:"hello, world!"}

  if posts[0] != want {
    t.Errorf("Expected %#v, got %#v", want, posts[0])
  }
})
t.Run("failing fs", func(t *testing.T) {
  _, err := blogposts.PostFromFS(FailingFS{})
  if err == nil {
    t.Error("Expected an error, but didn't get one")
  }
})
}
type FailingFS struct {
}
func (f FailingFS) Open(name string) (fs.File, error) {
	var err = errors.New("failed to open")
	return nil, err
}