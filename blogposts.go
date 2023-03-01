package blogposts

import (
	"io/fs"
)

func PostFromFS(filesystem fs.FS) []Post {
  dir, _ := fs.ReadDir(filesystem, ".")

  var posts []Post
  for _, f := range dir {
    post := makePostFromFile(filesystem, f.Name())
    posts = append(posts, post)
  }
  return posts
}

func makePostFromFile(filesystem fs.FS, fileName string) Post {
	blogFile, _    := filesystem.Open(fileName) 
  return newPost(blogFile)
}

