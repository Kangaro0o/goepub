package generator

import (
	"fmt"
	"github.com/Kangrao0o/goepub/utils"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
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
	tplFilename := fmt.Sprintf("%s\\..\\template\\epub3\\OEBPS\\text\\cover.html", dir)
	temp, err := template.New("cover.html").ParseFiles(tplFilename)
	if err != nil {
		return err
	}

	// 创建临时目录
	if err := utils.CreateDir(savePath); err != nil {
		log.Errorf("cover write err: %v when create tmp dir", err)
		return err
	}
	filename := fmt.Sprintf("%s/cover.html", savePath)
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
	filename := fmt.Sprintf("%s/%s", savePath, c.URL[idx+1:])
	return ioutil.WriteFile(filename, bytes, 0666)
}
