package main

import (
	"fmt"
	"github.com/Kangrao0o/goepub/generator"
	"os"
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
			Title:     "Cover",
			Src:       "1_togolife.jpg",
			Alt:       "test-book-name",
			URL:       "https://avatar.csdnimg.cn/2/9/0/1_togolife.jpg",
		},
		Style:   &generator.Style{},
		MetaInf: &generator.MetaInf{},
		Chapters: []*generator.Chapter{
			{
				ID:        1,
				Generator: "Created by Kelvin",
				Title:     "第一章",
				Content:   "<p>哈哈哈哈哈</p>",
			},
			{
				ID:        2,
				Generator: "Created by Kelvin2",
				Title:     "第二章",
				Content:   "<p>哈哈哈哈哈哦哦哦哦哦</p>",
			},
		},
		MimeType: &generator.MimeType{},
	}
	dir, _ := os.Getwd()
	err := book.Write(dir + "/book-uid")
	fmt.Println("err: ", err)
	fmt.Println("book write done")
}
