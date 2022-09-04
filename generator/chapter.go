package generator

import (
	"fmt"
	"os"
	"text/template"
)

type Chapter struct {
	ID        int32
	Generator string
	Title     string
	Content   string
}

func (c *Chapter) Write() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	tplFilename := fmt.Sprintf("%s\\..\\template\\epub3\\OEBPS\\text\\chapter.html", dir)

	temp, err := template.New("chapter.html").ParseFiles(tplFilename)
	if err != nil {
		return err
	}
	filename := fmt.Sprintf("chapter%d.html", c.ID)
	fd, err := os.Create(filename)
	if err != nil {
		return err
	}
	return temp.Execute(fd, c)
}
