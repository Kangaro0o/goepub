package generator

import (
	"reflect"
	"testing"
)

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
						Src:       "toc.ncx",
						MediaType: NCXMediaType,
					},
					{
						ID:        "htmltoc",
						Src:       "book-toc.html",
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
						Src:   "cover.html",
						Type:  CoverGuideType,
						Title: CoverGuideTitle,
					},
					{
						Src:   "toc.html",
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
			if err := doc.Write("D:\\Workspace\\GoProjects\\goepub\\books"); (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPackageDocument_GetManifests(t *testing.T) {
	type args struct {
		savePath string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				savePath: "D:\\Workspace\\GoProjects\\goepub\\template\\epub3",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetManifests(tt.args.savePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetManifests() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetManifests() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getShortName(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				filename: "styles\\style.css",
			},
			want: "style",
		},
		{
			name: "case2",
			args: args{
				filename: "content.opf",
			},
			want: "content",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getShortName(tt.args.filename); got != tt.want {
				t.Errorf("getShortName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMediaType(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want MediaType
	}{
		{
			name: "case1",
			args: args{
				filename: "styles\\style.css",
			},
			want: CSSMediaType,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMediaType(tt.args.filename); got != tt.want {
				t.Errorf("getMediaType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSpines(t *testing.T) {
	type args struct {
		savePath string
	}
	tests := []struct {
		name    string
		args    args
		want    []*Spine
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				savePath: "D:\\Workspace\\GoProjects\\goepub\\template\\epub3",
			},
			want: []*Spine{
				{
					IDRef:  "book-intro",
					Linear: YESLinear,
				},
				{
					IDRef:  "book-toc",
					Linear: YESLinear,
				},
				{
					IDRef:  "chapter",
					Linear: YESLinear,
				},
				{
					IDRef:  "cover",
					Linear: NOLinear,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSpines(tt.args.savePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSpines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSpines() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetGuides(t *testing.T) {
	type args struct {
		savePath string
	}
	tests := []struct {
		name    string
		args    args
		want    []*Guide
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				savePath: "D:\\Workspace\\GoProjects\\goepub\\template\\epub3",
			},
			want: []*Guide{
				{
					Src:   "text\\book-toc.html",
					Type:  TOCGuideType,
					Title: TOCGuideTitle,
				},
				{
					Src:   "text\\cover.html",
					Type:  CoverGuideType,
					Title: CoverGuideTitle,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetGuides(tt.args.savePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGuides() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGuides() got = %v, want %v", got, tt.want)
			}
		})
	}
}
