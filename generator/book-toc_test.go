package generator

import "testing"

func TestBookTOC_Write(t *testing.T) {
	type fields struct {
		Generator string
		Chapters  []*Chapter
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "case1",
			fields: fields{
				Generator: "Created by Kelvin",
				Chapters: []*Chapter{
					{
						Src:   "chapter0.html",
						Title: "内容简介",
					},
					{
						Src:   "chapter1.html",
						Title: "第一章 test-chapter",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			toc := &BookTOC{
				Generator: tt.fields.Generator,
				Chapters:  tt.fields.Chapters,
			}
			if err := toc.Write(); (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
