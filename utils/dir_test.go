package utils

import "testing"

func TestCopyDir(t *testing.T) {
	type args struct {
		srcPath  string
		destPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				srcPath:  "D:\\Workspace\\GoProjects\\goepub\\template\\epub3\\OEBPS\\styles",
				destPath: "D:\\Workspace\\GoProjects\\goepub\\epub3-book\\OEBPS\\styles",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CopyDir(tt.args.srcPath, tt.args.destPath); (err != nil) != tt.wantErr {
				t.Errorf("CopyDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
