package generator

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
}

func (book *Book) convertToNCX() *NCXDocument {
	var navMap []*NavPoint
	for idx, chapter := range book.Chapters {
		navPoint := &NavPoint{
			ID:         chapter.NavPointID,
			PlayOrder:  int32(idx),
			Label:      chapter.Title,
			ContentSrc: chapter.NavPointSrc,
		}
		navMap = append(navMap, navPoint)
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

func (book *Book) convertToOPF() *PackageDocument {
	opf := &PackageDocument{
		UID:      book.UID,
		BookName: book.Name,
		Author:   book.Author,
		Date:     book.Date,
		Rights:   book.Rights,
		Language: book.Language,
	}
	var (
		manifests []*Manifest
		spines    []*Spine
		guides    []*Guide
	)
	for idx, chapter := range book.Chapters {
		if idx < 3 {
			guide := &Guide{
				Href:  chapter.NavPointSrc,
				Type:  TOCGuideType,
				Title: TOCGuideTitle,
			}
			guides = append(guides, guide)
		}
		manifest := &Manifest{
			ID:        chapter.NavPointID,
			Href:      chapter.NavPointSrc,
			MediaType: chapter.MediaType,
		}
		spine := &Spine{
			IDRef:  chapter.NavPointID,
			Linear: YESLinear,
		}
		manifests = append(manifests, manifest)
		spines = append(spines, spine)
	}
	opf.Manifests = manifests
	opf.Spines = spines
	opf.Guides = guides
	return opf
}

func (book *Book) Write(savePath string) error {
	// 生成 toc.ncx 文件
	ncx := book.convertToNCX()
	if err := ncx.Write(savePath); err != nil {
		return err
	}
	// 生成 content.opf 文件
	opf := book.convertToOPF()
	if err := opf.Write(savePath); err != nil {
		return err
	}
	// 下载封面图
	cover := book.Cover
	if err := cover.Download(savePath + "/images"); err != nil {
		return err
	}
	// 生成封面页
	if err := cover.Write(savePath); err != nil {
		return err
	}
	// 生成章节
	for _, chapter := range book.Chapters {
		if err := chapter.Write(savePath + "/text"); err != nil {
			return err
		}
	}
	// TODO 拷贝 css
	return nil
}
