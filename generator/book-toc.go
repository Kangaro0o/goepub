package generator

import (
	"fmt"
	"os"
	"text/template"
)

// BookTOC 目录
type BookTOC struct {
	Generator     string
	NavPointID    string
	PlayOrder     int32
	NavPointLabel string
	ContentSrc    string
	Chapters      []*Chapter
}

type ChapterBrief struct {
	Href string
	Name string
}

func (toc *BookTOC) ConvertToNavPoint() *NavPoint {
	return &NavPoint{
		ID:         toc.NavPointID,
		PlayOrder:  toc.PlayOrder,
		Label:      toc.NavPointLabel,
		ContentSrc: toc.ContentSrc,
	}
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
	defer fd.Close()
	if err != nil {
		return err
	}
	return temp.Execute(fd, toc)
}
