package blogposts

import (
	"io"
	"io/fs"
	"strings"
)

type Post struct {
  Title string
}

func PostFromFS(filesystem fs.FS) []Post {
  dir, _ := fs.ReadDir(filesystem, ".")

  var posts []Post
  for _, f := range dir {
    post := makePostFromFile(filesystem, f)
    posts = append(posts, post)
  }
  return posts
}

func makePostFromFile(filesystem fs.FS, f fs.DirEntry) Post {
	blogFile, _    := filesystem.Open(f.Name()) 
  fileContent, _ := io.ReadAll(blogFile) 
  title := strings.TrimPrefix(string(fileContent), "Title: ")
  return Post{
  	Title: title,
  }
}