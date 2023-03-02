package main

import (
	"blogposts"
	"fmt"
	"log"
	"os"
)

func main() {
  posts, err := blogposts.PostFromFS(os.DirFS("posts"))
  if err != nil {
    log.Fatal(err) 
  }
  fmt.Println(posts)
}