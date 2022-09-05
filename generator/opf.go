package generator

import (
	"fmt"
	"github.com/Kangrao0o/goepub/utils"
	log "github.com/sirupsen/logrus"
	"os"
	"text/template"
)

type PackageDocument struct {
	UID       string
	BookName  string
	Author    string
	Date      string
	Rights    string
	Language  string
	Manifests []*Manifest
	Spines    []*Spine
	Guides    []*Guide
}

// Manifest 目录清单
type Manifest struct {
	ID        string
	Href      string
	MediaType MediaType
}

type MediaType string

// 文件媒体类型定义
const (
	NCXMediaType  MediaType = "application/x-dtbncx+xml"
	HTMLMediaType MediaType = "application/xhtml+xml"
	CSSMediaType  MediaType = "text/css"
	JPGMediaType  MediaType = "image/jpeg"
)

type Spine struct {
	IDRef  string
	Linear Linear
}

type Linear string

const (
	YESLinear Linear = "yes"
	NOLinear  Linear = "no"
)

type Guide struct {
	Href  string
	Type  GuideType
	Title GuideTitle
}

type GuideType string

const (
	CoverGuideType GuideType = "cover"
	TOCGuideType   GuideType = "toc"
	TextGuideType  GuideType = "text"
)

type GuideTitle string

const (
	CoverGuideTitle GuideTitle = "Cover"
	TOCGuideTitle   GuideTitle = "Table Of Contents"
	TextGuideTitle  GuideTitle = "Beginning"
)

func (doc *PackageDocument) Write(savePath string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	tmpFilename := fmt.Sprintf("%s\\..\\template\\epub3\\OEBPS\\content.opf", dir)
	temp, err := template.New("content.opf").ParseFiles(tmpFilename)
	if err != nil {
		return err
	}

	// 创建临时目录
	if err := utils.CreateDir(savePath); err != nil {
		log.Errorf("opf write err: %v when create tmp dir", err)
		return err
	}
	filename := fmt.Sprintf("%s/content.opf", savePath)
	fd, err := os.Create(filename)
	if err != nil {
		return err
	}
	return temp.Execute(fd, doc)
}
