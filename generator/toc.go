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
	NavMap    []NavPoint
}

type NavPoint struct {
	ID        string
	PlayOrder int32
	Label     string
	Content   string
}

func WriteToc(document *NCXDocument) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	filename := fmt.Sprintf("%s\\..\\template\\epub3\\OEBPS\\toc.ncx", dir)
	temp, err := template.New("toc.ncx").ParseFiles(filename)
	if err != nil {
		return err
	}
	temp.Execute(os.Stdout, document)
	return nil
}
