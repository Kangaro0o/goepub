package generator

import (
	"github.com/Kangrao0o/goepub/utils"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

type Book struct {
	UID       string
	Generator string
	Name      string
	Author    string
	Date      string
	Rights    string
	Language  string
	Chapters  []*Chapter
	Cover     *Cover
	Toc       *BookTOC
	Style     *Style
	MetaInf   *MetaInf
	MimeType  *MimeType
	SavePath  string // 存储目录
}

func (book *Book) convertToNCX() *NCXDocument {
	var navMap []*NavPoint
	if book.Cover != nil {
		// 添加封面
		c := book.Cover
		navMap = append(navMap, c.ConvertToNavPoint())
	}
	if book.Toc != nil {
		// 添加目录
		toc := book.Toc
		navMap = append(navMap, toc.ConvertToNavPoint())
	}
	idx := int32(len(navMap))
	for _, chapter := range book.Chapters {
		navPoint := &NavPoint{
			ID:         chapter.ID,
			PlayOrder:  idx,
			Label:      chapter.Title,
			ContentSrc: chapter.Src,
		}
		navMap = append(navMap, navPoint)
		idx++
	}
	ncx := &NCXDocument{
		UID:       book.UID,
		Generator: book.Generator,
		BookName:  book.Name,
		Author:    book.Author,
		NavMap:    navMap,
	}
	return ncx
}

func (book *Book) convertToOPF() (*PackageDocument, error) {
	opf := &PackageDocument{
		UID:      book.UID,
		BookName: book.Name,
		Author:   book.Author,
		Date:     book.Date,
		Rights:   book.Rights,
		Language: book.Language,
		CoverID:  getShortName(book.Cover.ImgSrc),
	}
	// 合成 manifests
	manifests, err := GetManifests(book.SavePath)
	if err != nil {
		log.Errorf("convert to opf failed, err: %v", err)
		return nil, err
	}
	// 合成 spines
	spines, err := GetSpines(book.SavePath)
	if err != nil {
		log.Errorf("convert to opf failed, err: %v", err)
		return nil, err
	}
	// 合成 guides
	guides, err := GetGuides(book.SavePath)
	if err != nil {
		log.Errorf("convert to opf guide failed, err: %v", err)
		return nil, err
	}
	opf.Manifests = manifests
	opf.Spines = spines
	opf.Guides = guides
	return opf, nil
}

func (book *Book) Write(savePath string) error {
	book.SavePath = savePath
	// 生成 toc.ncx 文件
	ncx := book.convertToNCX()
	if err := ncx.Write(savePath + "/OEBPS"); err != nil {
		return err
	}

	// 下载封面图
	cover := book.Cover
	if err := cover.Download(savePath + "/OEBPS/images"); err != nil {
		return err
	}
	// 生成封面页
	if err := cover.Write(savePath + "/OEBPS/text"); err != nil {
		return err
	}
	// 生成章节
	for idx, chapter := range book.Chapters {
		if err := chapter.Write(savePath+"/OEBPS/text", int32(idx)); err != nil {
			return err
		}
	}
	// 拷贝 css
	style := book.Style
	if err := style.Write(savePath + "/OEBPS/styles"); err != nil {
		return err
	}
	// 拷贝 container.xml
	metaInf := book.MetaInf
	if err := metaInf.Write(savePath + "/META-INF"); err != nil {
		return err
	}
	// 拷贝 mimetype
	mimetype := book.MimeType
	if err := mimetype.Write(savePath); err != nil {
		return err
	}
	// 最后步骤 生成 content.opf 文件
	opf, err := book.convertToOPF()
	if err != nil {
		return err
	}
	if err := opf.Write(savePath + "/OEBPS"); err != nil {
		return err
	}
	return nil
}

func (book *Book) MakeEpub3() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	tmpSavePath := filepath.Join(dir, book.UID)
	if err := book.Write(tmpSavePath); err != nil {
		return err
	}

	filenames, err := utils.GetAllFile(book.SavePath)
	if err != nil {
		return err
	}
	output := "test.epub"
	if err := utils.ZipFiles(output, filenames, tmpSavePath); err != nil {
		return err
	}
	return nil
}
