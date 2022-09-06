package generator

import (
	"fmt"
	"github.com/Kangrao0o/goepub/resource"
	"github.com/Kangrao0o/goepub/utils"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Chapter struct {
	ID        string
	Generator string
	Title     string
	Content   string
	Src       string
	PlayOrder int32
}

func (c *Chapter) Write(savePath string, index int32) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	dir = strings.TrimSuffix(dir, "generator")
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

	filename := filepath.Join(savePath, fmt.Sprintf("chapter%d.html", index))
	fd, err := os.Create(filename)
	defer fd.Close()
	if err != nil {
		return err
	}
	return temp.Execute(fd, c)
}
