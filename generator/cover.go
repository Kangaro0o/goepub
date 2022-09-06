package generator

import (
	"github.com/Kangaro0o/goepub/resource"
	"github.com/Kangaro0o/goepub/utils"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// Cover 封面
type Cover struct {
	Generator string
	Title     string
	Desc      string
	ImgSrc    string
	HtmlSrc   string
	Alt       string
	URL       string
	PlayOrder int32
}

func (c *Cover) ConvertToNavPoint() *NavPoint {
	return &NavPoint{
		ID:         c.Title,
		PlayOrder:  c.PlayOrder,
		Label:      c.Desc,
		ContentSrc: c.HtmlSrc,
	}
}

// Write 写入 cover.html
func (c *Cover) Write(savePath string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	dir = strings.TrimSuffix(dir, "generator")
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
	defer fd.Close()
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
