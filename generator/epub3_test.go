package generator

import "testing"

func TestBook_Write(t *testing.T) {
	type fields struct {
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
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "case1",
			fields: fields{
				UID:       "uid-12345",
				Generator: "Created by Kelvin",
				Name:      "test-book",
				Author:    "test-author",
				Date:      "2021",
				Rights:    "it owns kelvin",
				Language:  "zh-CN",
				Cover: &Cover{
					Generator: "Created by Kelvin",
					Title:     "Cover",
					ImgSrc:    "1_togolife.jpg",
					Alt:       "test-book-name",
					URL:       "https://avatar.csdnimg.cn/2/9/0/1_togolife.jpg",
				},
				Chapters: []*Chapter{
					{
						ID:        "chapter0",
						Generator: "Created by Kelvin",
						Src:       "chapter0.html",
						Title:     "内容简介",
						Content:   "<p>这是一段内容简介</p>",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			book := &Book{
				UID:       tt.fields.UID,
				Generator: tt.fields.Generator,
				Name:      tt.fields.Name,
				Author:    tt.fields.Author,
				Date:      tt.fields.Date,
				Rights:    tt.fields.Rights,
				Language:  tt.fields.Language,
				Chapters:  tt.fields.Chapters,
				Cover:     tt.fields.Cover,
			}
			if err := book.Write("D:\\Workspace\\GoProjects\\goepub\\epub3-book"); (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
