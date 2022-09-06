package generator

import (
	"github.com/Kangrao0o/goepub/resource"
	"github.com/Kangrao0o/goepub/utils"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Cover struct {
	Generator string
	Title     string
	Src       string
	Alt       string
	URL       string
}

func (c *Cover) Write(savePath string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	tplFilename := filepath.Join(dir, resource.CoverEpub3Path)
	temp, err := template.New("cover.html").ParseFiles(tplFilename)
	if err != nil {
		return err
	}

	// 创建临时目录
	if err := utils.CreateDir(savePath); err != nil {
		log.Errorf("cover write err: %v when create tmp dir", err)
		return err
	}

	filename := filepath.Join(savePath, "cover.html")
	fd, err := os.Create(filename)
	if err != nil {
		return err
	}
	return temp.Execute(fd, c)
}

func (c *Cover) Download(savePath string) error {
	rsp, err := http.Get(c.URL)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()
	bytes, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err
	}

	savePath = strings.TrimSuffix(savePath, "/")
	if err := utils.CreateDir(savePath); err != nil {
		log.Errorf("download cover images err: %v when create tmp dir", err)
		return err
	}
	idx := strings.LastIndex(c.URL, "/")
	filename := filepath.Join(savePath, c.URL[idx+1:])
	return ioutil.WriteFile(filename, bytes, 0666)
}
