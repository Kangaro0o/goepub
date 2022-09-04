package generator

import "testing"

func TestPackageDocument_Write(t *testing.T) {
	type fields struct {
		UID       string
		BookName  string
		Author    string
		Date      string
		Rights    string
		Language  string
		Manifests []*Manifest
		Spines    []*Spine
		Guides    []*Guide
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "case1",
			fields: fields{
				UID:      "test-uid",
				BookName: "test-book-name",
				Author:   "test-author",
				Date:     "2021",
				Rights:   "Created by Kelvin",
				Language: "zh-CN",
				Manifests: []*Manifest{
					{
						ID:        "ncxtoc",
						Href:      "toc.ncx",
						MediaType: NCXMediaType,
					},
					{
						ID:        "htmltoc",
						Href:      "book-toc.html",
						MediaType: HTMLMediaType,
					},
				},
				Spines: []*Spine{
					{
						IDRef:  "cover111",
						Linear: YESLinear,
					},
					{
						IDRef:  "htmltoc222",
						Linear: NOLinear,
					},
				},
				Guides: []*Guide{
					{
						Href:  "cover.html",
						Type:  CoverGuideType,
						Title: CoverGuideTitle,
					},
					{
						Href:  "toc.html",
						Type:  TOCGuideType,
						Title: TOCGuideTitle,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc := &PackageDocument{
				UID:       tt.fields.UID,
				BookName:  tt.fields.BookName,
				Author:    tt.fields.Author,
				Date:      tt.fields.Date,
				Rights:    tt.fields.Rights,
				Language:  tt.fields.Language,
				Manifests: tt.fields.Manifests,
				Spines:    tt.fields.Spines,
				Guides:    tt.fields.Guides,
			}
			if err := doc.Write(); (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
