package generator

import "testing"

func TestChapter_Write(t *testing.T) {
	type fields struct {
		ID        string
		Generator string
		Title     string
		Content   string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "case1",
			fields: fields{
				ID:        "chapter1",
				Generator: "Created by Kelvin",
				Title:     "第一章",
				Content:   "<p>哈哈哈哈哈</p>",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Chapter{
				ID:        tt.fields.ID,
				Generator: tt.fields.Generator,
				Title:     tt.fields.Title,
				Content:   tt.fields.Content,
			}
			if err := c.Write("D:\\Workspace\\GoProjects\\goepub\\book-chapter", 0); (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
