package blogposts

import (
	"io/fs"
)

type Post struct {
  
}

func PostFromFS(filesystem fs.FS) []Post {
  dir, _ := fs.ReadDir(filesystem, ".")

  var posts []Post
  for range dir {
    posts = append(posts, Post{})
  }
  return posts
}