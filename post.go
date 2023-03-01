package blogposts

import (
	"bufio"
	"io"
	"strings"
)

type Post struct {
  Title string
	Description string
}

func newPost(blogFile io.Reader) Post {
	// fileContent, _ := io.ReadAll(blogFile)
	// title := strings.TrimPrefix(string(fileContent), "Title: ")
	// return Post{
	// 	Title: title,
	// }
	scanner := bufio.NewScanner(blogFile)

	readLine := func(prefix string) string {
		scanner.Scan() 
		return strings.TrimPrefix(scanner.Text(), prefix)
  }

	title := readLine("Title: ")
	description := readLine("Description: ")

	post := Post{
		Title:       title,
		Description: description,
	}
	return post
}