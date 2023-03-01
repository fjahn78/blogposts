package blogposts

import (
	"io"
	"strings"
)

type Post struct {
  Title string
}

func newPost(blogFile io.Reader) Post {
	fileContent, _ := io.ReadAll(blogFile)
	title := strings.TrimPrefix(string(fileContent), "Title: ")
	return Post{
		Title: title,
	}
}