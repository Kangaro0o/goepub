package generator

import (
	"fmt"
	"os"
	"text/template"
)

type BookIntro struct {
	Generator string
	Title     string
	Content   string
}

func (intro *BookIntro) Write() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	tplFilename := fmt.Sprintf("%s\\..\\template\\epub3\\OEBPS\\text\\book-intro.html", dir)
	temp, err := template.New("book-intro.html").ParseFiles(tplFilename)
	if err != nil {
		return err
	}

	fd, err := os.Create("book-intro.html")
	if err != nil {
		return err
	}
	return temp.Execute(fd, intro)
}
