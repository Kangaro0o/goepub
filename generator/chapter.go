package generator

import (
	"fmt"
	"github.com/Kangrao0o/goepub/utils"
	log "github.com/sirupsen/logrus"
	"os"
	"text/template"
)

type Chapter struct {
	ID          int32
	Generator   string
	NavPointID  string
	NavPointSrc string
	Title       string
	Content     string
	MediaType   MediaType
}

func (c *Chapter) Write(savePath string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	tplFilename := fmt.Sprintf("%s\\..\\template\\epub3\\OEBPS\\text\\chapter.html", dir)

	temp, err := template.New("chapter.html").ParseFiles(tplFilename)
	if err != nil {
		return err
	}

	// 创建临时目录
	if err := utils.CreateDir(savePath); err != nil {
		log.Errorf("chapter write err: %v when create tmp dir", err)
		return err
	}

	filename := fmt.Sprintf("%s/chapter%d.html", savePath, c.ID)
	fd, err := os.Create(filename)
	if err != nil {
		return err
	}
	return temp.Execute(fd, c)
}
