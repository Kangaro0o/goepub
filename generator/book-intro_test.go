package generator

import "testing"

func TestBookIntro_Write(t *testing.T) {
	type fields struct {
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
				Generator: "Created by Kelvin",
				Title:     "内容简介",
				Content:   "<p>哈哈哈哈啊哈 hello world</p>",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			intro := &BookIntro{
				Generator: tt.fields.Generator,
				Title:     tt.fields.Title,
				Content:   tt.fields.Content,
			}
			if err := intro.Write(); (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
