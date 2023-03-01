package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

const (
	titlePrefix = "Title: "
	descriptionPrefix = "Description: "
	tagsPrefix = "Tags: "
)

type Post struct {
  Title       string
	Description string
	Tags        []string
	Body        string
}

func newPost(blogFile io.Reader) Post {
	scanner := bufio.NewScanner(blogFile)

	readLine := func(prefix string) string {
		scanner.Scan() 
		return strings.TrimPrefix(scanner.Text(), prefix)
  }

	title       := readLine(titlePrefix)
	description := readLine(descriptionPrefix)
	tags        := strings.Split(readLine(tagsPrefix), ", ")
	body        := readBody(scanner)

	post := Post{
		Title:       title,
		Description: description,
		Tags:        tags,
		Body:        body,
	}
	return post
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan()
	var buffer bytes.Buffer
	for scanner.Scan() {
		fmt.Fprintln(&buffer, scanner.Text())
	}
	return strings.TrimSuffix(buffer.String(), "\n")
}