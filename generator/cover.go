package generator

import (
	"fmt"
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

func (c *Cover) Write() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	filename := fmt.Sprintf("%s\\..\\template\\epub3\\OEBPS\\text\\cover.html", dir)
	temp, err := template.New("cover.html").ParseFiles(filename)
	if err != nil {
		return err
	}
	fd, err := os.Create("cover.html")
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
	idx := strings.LastIndex(c.URL, "/")
	savePath = fmt.Sprintf("%s/%s", savePath, c.URL[idx+1:])
	return ioutil.WriteFile(savePath, bytes, 0666)
}
