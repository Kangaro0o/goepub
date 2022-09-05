package generator

import "testing"

func TestNCXDocument_Write(t *testing.T) {
	type fields struct {
		UID       string
		Generator string
		BookName  string
		Author    string
		NavMap    []*NavPoint
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "case1",
			fields: fields{
				UID:       "test-id",
				Generator: "test-gen",
				BookName:  "test-book",
				Author:    "test-author",
				NavMap: []*NavPoint{
					{
						ID:         "test-id",
						PlayOrder:  1,
						Label:      "第一章",
						ContentSrc: "chapter1.html",
					},
					{
						ID:         "test-id2",
						PlayOrder:  2,
						Label:      "第二章",
						ContentSrc: "chapter2.html",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc := &NCXDocument{
				UID:       tt.fields.UID,
				Generator: tt.fields.Generator,
				BookName:  tt.fields.BookName,
				Author:    tt.fields.Author,
				NavMap:    tt.fields.NavMap,
			}
			if err := doc.Write("D:\\Workspace\\GoProjects\\goepub\\books"); (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
