package main

import (
	"os"
	"text/template"
)

var template1 = `func (m *{{.ModelName}}) HMapKey() string {
	return fmt.Sprintf("{{.TableName}}:{{.EntityDBID}}:%v", m.{{.EntityID}})
}`

func main() {
	data := map[string]interface{}{
		"ModelName":  "A",
		"TableName":  "t1",
		"EntityDBID": "id",
		"EntityID":   "ID",
	}
	temp, _ := template.New("test").Parse(template1)
	temp.Execute(os.Stdout, data)
}
