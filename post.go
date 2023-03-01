package blogposts

import (
	"bufio"
	"io"
	"strings"
)

const (
	titlePrefix = "Title: "
	descriptionPrefix = "Description: "
	tagsPrefix = "Tags: "
)

type Post struct {
  Title string
	Description string
	Tags []string
}

func newPost(blogFile io.Reader) Post {
	scanner := bufio.NewScanner(blogFile)

	readLine := func(prefix string) string {
		scanner.Scan() 
		return strings.TrimPrefix(scanner.Text(), prefix)
  }

	title := readLine(titlePrefix)
	description := readLine(descriptionPrefix)
	tags := strings.Split(readLine(tagsPrefix), ", ")

	post := Post{
		Title:       title,
		Description: description,
		Tags:        tags,
	}
	return post
}