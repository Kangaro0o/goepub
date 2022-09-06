package generator

import (
	"fmt"
	"github.com/Kangrao0o/goepub/resource"
	"github.com/Kangrao0o/goepub/utils"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
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

	tplFilename := filepath.Join(dir, resource.ChapterEpub3Path)
	temp, err := template.New("chapter.html").ParseFiles(tplFilename)
	if err != nil {
		return err
	}

	// 创建临时目录
	if err := utils.CreateDir(savePath); err != nil {
		log.Errorf("chapter write err: %v when create tmp dir", err)
		return err
	}

	filename := filepath.Join(savePath, fmt.Sprintf("chapter%d.html", c.ID))
	fd, err := os.Create(filename)
	if err != nil {
		return err
	}
	return temp.Execute(fd, c)
}
