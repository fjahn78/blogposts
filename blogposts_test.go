package blogposts_test

import (
	"blogposts"
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestPostsFromFS(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		// Given
		fileSystem := fstest.MapFS{
			"hello-world.md":  {Data: []byte("Title: hello, world!\nDescription: What's up, world!\nTags: tdd, go")},
			"hello-world2.md": {Data: []byte("Title: Hola, mundo!\nDescription: que pasa, mundo!\nTags: tdd, go")},
		}

		// When
		posts, err := blogposts.PostFromFS(fileSystem)

		// Then
		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fileSystem) {
			t.Errorf("Expected %d posts, got %d", len(fileSystem), len(posts))
		}

		want := blogposts.Post{
			Title:       "hello, world!",
			Description: "What's up, world!",
      Tags:        []string{"tdd", "go"},
		}

		assertPost(t, posts[0], want)
	})
	t.Run("failing fs", func(t *testing.T) {
		_, err := blogposts.PostFromFS(FailingFS{})
		if err == nil {
			t.Error("Expected an error, but didn't get one")
		}
	})
}

func assertPost(t *testing.T, got, want blogposts.Post) {
  t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %#v, got %#v", want, got)
	}
}

type FailingFS struct {}

func (f FailingFS) Open(name string) (fs.File, error) {
	var err = errors.New("failed to open")
	return nil, err
}
