package main

import (
	"fmt"
	"github.com/Kangaro0o/goepub/generator"
)

//import (
//	"os"
//	"text/template"
//)

//var template1 = `func (m *{{.ModelName}}) HMapKey() string {
//	return fmt.Sprintf("{{.TableName}}:{{.EntityDBID}}:%v", m.{{.EntityID}})
//}`
//
//func main() {
//	data := map[string]interface{}{
//		"ModelName":  "A",
//		"TableName":  "t1",
//		"EntityDBID": "id",
//		"EntityID":   "ID",
//	}
//	temp, _ := template.New("test").Parse(template1)
//	temp.Execute(os.Stdout, data)
//}

func main() {
	book := &generator.Book{
		UID:       "uid-12345",
		Generator: "Created by Kelvin",
		Name:      "test-book",
		Author:    "test-author",
		Date:      "2021",
		Rights:    "it owns kelvin",
		Language:  "zh-CN",
		Cover: &generator.Cover{
			Generator: "Created by Kelvin",
			Title:     "cover",
			Desc:      "封面",
			PlayOrder: 0,
			ImgSrc:    "images/1_togolife.jpg",
			HtmlSrc:   "text/cover.html",
			Alt:       "test-book-name",
			URL:       "https://avatar.csdnimg.cn/2/9/0/1_togolife.jpg",
		},
		Toc: &generator.BookTOC{
			Generator:     "Created by Kelvin",
			NavPointID:    "book-toc",
			PlayOrder:     1,
			NavPointLabel: "目录",
			ContentSrc:    "text/book-toc.html",
		},
		Style:   &generator.Style{},
		MetaInf: &generator.MetaInf{},

		Chapters: []*generator.Chapter{
			{
				ID:        "chapter0",
				Generator: "Created by kelvin",
				Title:     "内容简介",
				Content:   "<p>this is introduction</p>",
				Src:       "text/chapter0.html",
			},
			{
				ID:        "chapter1",
				Generator: "Created by Kelvin",
				Title:     "第一章",
				Content:   "<p>哈哈哈哈哈</p>",
				Src:       "text/chapter1.html",
			},
			{
				ID:        "chapter2",
				Generator: "Created by Kelvin2",
				Title:     "第二章",
				Content:   "<p>哈哈哈哈哈哦哦哦哦哦</p>",
				Src:       "text/chapter2.html",
			},
		},
		MimeType: &generator.MimeType{},
	}
	//dir, _ := os.Getwd()
	//err := book.Write(dir + "/book-uid")
	err := book.MakeEpub3()
	fmt.Println("err: ", err)
	fmt.Println("book write done")
}
