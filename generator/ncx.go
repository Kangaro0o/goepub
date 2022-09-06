package generator

import (
	"github.com/Kangrao0o/goepub/resource"
	"github.com/Kangrao0o/goepub/utils"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
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
	ID         string
	PlayOrder  int32
	Label      string
	ContentSrc string
}

func (doc *NCXDocument) Write(savePath string) error {
	dir, err := os.Getwd()
	if err != nil {
		log.Errorf("ncx write err: %v when os.Getwd", err)
		return err
	}

	tplFilename := filepath.Join(dir, resource.NCXEpub3Path)
	temp, err := template.New("toc.ncx").ParseFiles(tplFilename)
	if err != nil {
		log.Errorf("ncx write err: %v when template parse files", err)
		return err
	}

	if err := utils.CreateDir(savePath); err != nil {
		return err
	}
	filename := filepath.Join(savePath, "toc.ncx")
	fd, err := os.Create(filename)
	if err != nil {
		log.Errorf("ncx write err: %v when os.Create", err)
		return err
	}
	return temp.Execute(fd, doc)
}
