package generator

import (
	"fmt"
	"github.com/Kangaro0o/goepub/resource"
	"github.com/Kangaro0o/goepub/utils"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
)

type PackageDocument struct {
	UID       string
	BookName  string
	Author    string
	Date      string
	Rights    string
	Language  string
	CoverID   string
	Manifests []*Manifest
	Spines    []*Spine
	Guides    []*Guide
}

// Manifest 目录清单
type Manifest struct {
	ID        string
	Src       string
	MediaType MediaType
}

type MediaType string

// 文件媒体类型定义
const (
	NCXMediaType  MediaType = "application/x-dtbncx+xml"
	HTMLMediaType MediaType = "application/xhtml+xml"
	CSSMediaType  MediaType = "text/css"
	JPGMediaType  MediaType = "images/jpeg"
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
	Src   string
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

// Write 写入 content.opf
func (doc *PackageDocument) Write(savePath string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	dir = strings.TrimSuffix(dir, "generator")
	tplFilename := filepath.Join(dir, resource.OPFEpub3Path)
	temp, err := template.New("content.opf").ParseFiles(tplFilename)
	if err != nil {
		return err
	}

	// 创建临时目录
	if err := utils.CreateDir(savePath); err != nil {
		log.Errorf("opf write err: %v when create tmp dir", err)
		return err
	}
	filename := filepath.Join(savePath, "content.opf")
	fd, err := os.Create(filename)
	defer fd.Close()
	if err != nil {
		return err
	}
	return temp.Execute(fd, doc)
}

// GetManifests 获取 manifest 列表
func GetManifests(savePath string) ([]*Manifest, error) {
	destDir := filepath.Join(savePath, "OEBPS")
	filenames, err := utils.GetAllFile(destDir)
	if err != nil {
		return nil, err
	}

	sort.Strings(filenames)
	var manifests []*Manifest
	for _, f := range filenames {
		if strings.HasSuffix(f, ".opf") {
			continue
		}
		// 拿到相对路径
		f1 := strings.TrimPrefix(f, destDir+string(filepath.Separator))
		m := &Manifest{
			ID:        getShortName(f),
			Src:       strings.ReplaceAll(f1, "\\", "/"),
			MediaType: getMediaType(f),
		}
		manifests = append(manifests, m)
	}
	return manifests, nil
}

// GetSpines 获取 spine 列表
func GetSpines(savePath string) ([]*Spine, error) {
	destDir := filepath.Join(savePath, "OEBPS", "text")
	filenames, err := utils.GetAllFile(destDir)
	if err != nil {
		return nil, err
	}
	var spines []*Spine
	for _, f := range filenames {
		s := &Spine{
			IDRef:  getShortName(f),
			Linear: YESLinear,
		}
		if strings.Contains(f, "cover.html") {
			s.Linear = NOLinear
		}
		spines = append(spines, s)
	}
	return spines, nil
}

func GetGuides(savePath string) ([]*Guide, error) {
	destDir := filepath.Join(savePath, "OEBPS", "text")
	filenames, err := utils.GetAllFile(destDir)
	if err != nil {
		return nil, err
	}

	prefixPath := filepath.Join(savePath, "OEBPS")
	var guides []*Guide
	for _, f := range filenames {
		f1 := strings.TrimPrefix(f, prefixPath+string(filepath.Separator))
		if strings.Contains(f, "cover.html") {
			guides = append(guides, &Guide{
				Src:   strings.ReplaceAll(f1, "\\", "/"),
				Type:  CoverGuideType,
				Title: CoverGuideTitle,
			})
		}
		if strings.Contains(f, "book-toc.html") {
			guides = append(guides, &Guide{
				Src:   strings.ReplaceAll(f1, "\\", "/"),
				Type:  TOCGuideType,
				Title: TOCGuideTitle,
			})
		}
		if strings.Contains(f, "chapter0.html") {
			guides = append(guides, &Guide{
				Src:   strings.ReplaceAll(f1, "\\", "/"),
				Type:  TextGuideType,
				Title: TextGuideTitle,
			})
		}
	}
	return guides, nil
}

func getShortName(filename string) string {
	start := strings.LastIndex(filename, string(filepath.Separator))
	end := strings.LastIndex(filename, ".")
	if start == -1 {
		start = strings.LastIndex(filename, "/")
	}
	return filename[start+1 : end]
}

func getMediaType(filename string) MediaType {
	ext := strings.ToLower(filepath.Ext(filename))
	fmt.Println("ext: ", ext)
	switch ext {
	case ".css":
		return CSSMediaType
	case ".ncx":
		return NCXMediaType
	case ".jpg":
		return JPGMediaType
	default:
		return HTMLMediaType
	}
}
