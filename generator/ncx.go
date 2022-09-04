package generator

import (
	"fmt"
	"os"
	"text/template"
)

type NCXDocument struct {
	UID       string
	Generator string
	BookName  string
	Author    string
	NavMap    []*NavPoint
}

type NavPoint struct {
	ID        string
	PlayOrder int32
	Label     string
	Content   string
}

func (doc *NCXDocument) Write() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	tplFilename := fmt.Sprintf("%s\\..\\template\\epub3\\OEBPS\\toc.ncx", dir)
	temp, err := template.New("toc.ncx").ParseFiles(tplFilename)
	if err != nil {
		return err
	}

	fd, err := os.Create("toc.ncx")
	if err != nil {
		return err
	}
	return temp.Execute(fd, doc)
}
