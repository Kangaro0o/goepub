package resource

import (
	"path/filepath"
)

var (
	ContainerEpub3Path = filepath.Join("template", "epub3", "META-INF", "container.xml")
	OPFEpub3Path       = filepath.Join("template", "epub3", "OEBPS", "content.opf")
	NCXEpub3Path       = filepath.Join("template", "epub3", "OEBPS", "toc.ncx")
	CoverEpub3Path     = filepath.Join("template", "epub3", "OEBPS", "text", "cover.html")
	StyleEpub3Path     = filepath.Join("template", "epub3", "OEBPS", "styles")
	MetaInfEpub3Path   = filepath.Join("template", "epub3", "META-INF")
	ChapterEpub3Path   = filepath.Join("template", "epub3", "OEBPS", "text", "chapter.html")
	MimeTypeEpub3Path  = filepath.Join("template", "epub3", "mimetype")
)
