package generator

import "testing"

func TestWriteToc(t *testing.T) {
	type args struct {
		document *NCXDocument
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				document: &NCXDocument{
					UID:       "test-id",
					Generator: "test-gen",
					BookName:  "test-book",
					Author:    "test-author",
					NavMap: []NavPoint{
						{
							ID:        "test-id",
							PlayOrder: 1,
							Label:     "第一章",
							Content:   "chapter1.html",
						},
						{
							ID:        "test-id2",
							PlayOrder: 2,
							Label:     "第二章",
							Content:   "chapter2.html",
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteToc(tt.args.document); (err != nil) != tt.wantErr {
				t.Errorf("WriteToc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
