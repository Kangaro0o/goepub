package generator

import (
	"fmt"
	"os"
	"text/template"
)

type BookTOC struct {
	Generator string
	Chapters  []*ChapterBrief
}

type ChapterBrief struct {
	Href string
	Name string
}

func (toc *BookTOC) Write() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	tplFilename := fmt.Sprintf("%s\\..\\template\\epub3\\OEBPS\\text\\book-toc.html", dir)
	temp, err := template.New("book-toc.html").ParseFiles(tplFilename)
	if err != nil {
		return err
	}

	fd, err := os.Create("book-toc.html")
	if err != nil {
		return err
	}
	return temp.Execute(fd, toc)
}
