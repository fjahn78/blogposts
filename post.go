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
  Title, Description, Body string
	Tags                     []string
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
		_, err := fmt.Fprintln(&buffer, scanner.Text())
		if err != nil {
			panic(err)
		}
	}
	return strings.TrimSuffix(buffer.String(), "\n")
}