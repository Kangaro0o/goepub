package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"strings"
)

func main2() {
	files := []string{
		"./epub-examples/example2/EPUB/css/images/iconmonstr-info-8-icon-20x20.png",
		"./epub-examples/example2/EPUB/css/images/iconmonstr-info-8-icon.png",
		"./epub-examples/example2/EPUB/css/images/iconmonstr-window-new-icon-20x20.png",
		"./epub-examples/example2/EPUB/css/images/iconmonstr-window-new-icon.png",

		"./epub-examples/example2/EPUB/css/theme/default.css",
		"./epub-examples/example2/EPUB/css/theme/night.css",
		"./epub-examples/example2/EPUB/css/theme/sepia.css",
		"./epub-examples/example2/EPUB/css/main.css",

		"./epub-examples/example2/EPUB/images/AAHQPRT0.jpg",
		"./epub-examples/example2/EPUB/images/ch27_f0018-01.jpg",
		"./epub-examples/example2/EPUB/images/cover.jpg",
		"./epub-examples/example2/EPUB/images/f00xi-01.jpg",
		"./epub-examples/example2/EPUB/images/f00xi-02.jpg",
		"./epub-examples/example2/EPUB/images/f0007-01.jpg",
		"./epub-examples/example2/EPUB/images/f0244-01.png",
		"./epub-examples/example2/EPUB/images/f0245-01.png",
		"./epub-examples/example2/EPUB/images/f0246-01.png",
		"./epub-examples/example2/EPUB/images/fxvii-01.png",
		"./epub-examples/example2/EPUB/images/pub1.jpg",
		"./epub-examples/example2/EPUB/images/study.jpg",

		"./epub-examples/example2/EPUB/xhtml/afterword.xhtml",
		"./epub-examples/example2/EPUB/xhtml/appendix-a.xhtml",
		"./epub-examples/example2/EPUB/xhtml/asides.xhtml",
		"./epub-examples/example2/EPUB/xhtml/bibliography.xhtml",
		"./epub-examples/example2/EPUB/xhtml/body-content.xhtml",
		"./epub-examples/example2/EPUB/xhtml/colophon.xhtml",
		"./epub-examples/example2/EPUB/xhtml/copyright.xhtml",
		"./epub-examples/example2/EPUB/xhtml/cover.xhtml",
		"./epub-examples/example2/EPUB/xhtml/credits.xhtml",
		"./epub-examples/example2/EPUB/xhtml/dedication.xhtml",
		"./epub-examples/example2/EPUB/xhtml/endnotes.xhtml",
		"./epub-examples/example2/EPUB/xhtml/figures-tables.xhtml",
		"./epub-examples/example2/EPUB/xhtml/foreword.xhtml",
		"./epub-examples/example2/EPUB/xhtml/glossary.xhtml",
		"./epub-examples/example2/EPUB/xhtml/halftitle.xhtml",
		"./epub-examples/example2/EPUB/xhtml/links.xhtml",
		"./epub-examples/example2/EPUB/xhtml/name-index.xhtml",
		"./epub-examples/example2/EPUB/xhtml/nav.xhtml",
		"./epub-examples/example2/EPUB/xhtml/part01.xhtml",
		"./epub-examples/example2/EPUB/xhtml/preface.xhtml",
		"./epub-examples/example2/EPUB/xhtml/seriespage.xhtml",
		"./epub-examples/example2/EPUB/xhtml/subject-index.xhtml",
		"./epub-examples/example2/EPUB/xhtml/titlepage.xhtml",
		"./epub-examples/example2/EPUB/xhtml/vol01.xhtml",

		"./epub-examples/example2/EPUB/package.opf",
		"./epub-examples/example2/EPUB/toc.ncx",

		"./epub-examples/example2/META-INF/container.xml",

		"./epub-examples/example2/mimetype",
	}

	output := "test.epub"
	if err := ZipFiles(output, files); err != nil {
		panic(err)
	}
	fmt.Println("Zipped Files: ", output)
}

func ZipFiles(filename string, files []string) error {
	newZipFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	for _, file := range files {
		if err = AddFileToZip(zipWriter, file); err != nil {
			return err
		}
	}
	return nil
}

func AddFileToZip(zipWriter *zip.Writer, filename string) error {
	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Name = strings.TrimPrefix(filename, "./epub-examples/example2/")
	header.Method = zip.Deflate
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	return err
}
