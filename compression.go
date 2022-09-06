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
		".\\uid-12345\\META-INF\\container.xml",
		".\\uid-12345\\OEBPS\\images\\1_togolife.jpg",
		".\\uid-12345\\OEBPS\\styles\\style.css",
		".\\uid-12345\\OEBPS\\text\\chapter0.html",
		".\\uid-12345\\OEBPS\\text\\chapter1.html",
		".\\uid-12345\\OEBPS\\text\\chapter2.html",
		".\\uid-12345\\OEBPS\\text\\cover.html",
		".\\uid-12345\\OEBPS\\content.opf",
		".\\uid-12345\\OEBPS\\toc.ncx",
		".\\uid-12345\\mimetype",
	}

	output := "test123.epub"
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

	header.Name = strings.TrimPrefix(filename, ".\\uid-12345\\")
	header.Method = zip.Deflate
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	return err
}
