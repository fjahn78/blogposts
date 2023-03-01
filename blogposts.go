package blogposts

import (
	"io/fs"
)

func PostFromFS(filesystem fs.FS) ([]Post, error) {
  dir, err := fs.ReadDir(filesystem, ".")

  if err != nil {
    return nil, err
  }

  var posts []Post
  for _, f := range dir {
    post := makePostFromFile(filesystem, f.Name())
    posts = append(posts, post)
  }
  return posts, nil
}

func makePostFromFile(filesystem fs.FS, fileName string) Post {
	blogFile, _    := filesystem.Open(fileName) 
  return newPost(blogFile)
}

