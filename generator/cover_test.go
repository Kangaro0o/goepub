package generator

import "testing"

func TestCover_Write(t *testing.T) {
	type fields struct {
		Generator string
		Title     string
		Src       string
		Alt       string
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
				Title:     "Cover123",
				Src:       "image/cover.jpg",
				Alt:       "test-book-name",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cover{
				Generator: tt.fields.Generator,
				Title:     tt.fields.Title,
				Src:       tt.fields.Src,
				Alt:       tt.fields.Alt,
			}
			if err := c.Write(); (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCover_Download(t *testing.T) {
	type fields struct {
		Generator string
		Title     string
		Src       string
		Alt       string
	}
	type args struct {
		savePath string
		url      string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "case1",
			fields: fields{},
			args: args{
				savePath: "D:\\Workspace\\GoProjects\\goepub\\generator",
				url:      "https://avatar.csdnimg.cn/2/9/0/1_togolife.jpg",
			},
			wantErr: false,
		},
		{
			name: "case2",
			args: args{
				savePath: "D:\\Workspace\\GoProjects\\goepub\\generator",
				url:      "https://hbimg.huaban.com/32f065b3afb3fb36b75a5cbc90051b1050e1e6b6e199-Ml6q9F_fw320",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cover{
				Generator: tt.fields.Generator,
				Title:     tt.fields.Title,
				Src:       tt.fields.Src,
				Alt:       tt.fields.Alt,
			}
			if err := c.Download(tt.args.savePath, tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("Download() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
