package generator

import "testing"

func TestStyle_Write(t *testing.T) {
	type args struct {
		savePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				savePath: "D:\\Workspace\\GoProjects\\goepub\\epub3-book\\OEBPS\\styles",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Style{}
			if err := s.Write(tt.args.savePath); (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
