package zeropad

import (
	"testing"
)

func TestZeroPadFileName(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
		want     string
		wantErr  bool
	}{
		{
			name:     "valid_number",
			fileName: "123.png",
			want:     "0123.png",
			wantErr:  false,
		},
		{
			name:     "valid_number_no_extension",
			fileName: "45",
			want:     "0045",
			wantErr:  false,
		},
		{
			name:     "valid_number_with_multiple_digits",
			fileName: "9999.jpeg",
			want:     "9999.jpeg",
			wantErr:  false,
		},
		{
			name:     "number_with_leading_zeros",
			fileName: "007.txt",
			want:     "0007.txt",
			wantErr:  false,
		},
		{
			name:     "invalid_non_numeric_file_name",
			fileName: "file.png",
			want:     "",
			wantErr:  true,
		},
		{
			name:     "invalid_empty_string",
			fileName: "",
			want:     "",
			wantErr:  true,
		},
		{
			name:     "invalid_file_with_only_extension",
			fileName: ".jpg",
			want:     "",
			wantErr:  true,
		},
		{
			name:     "valid_number_with_large_digits",
			fileName: "123456789.csv",
			want:     "123456789.csv",
			wantErr:  false,
		},
		{
			name:     "invalid_special_characters",
			fileName: "!@#.png",
			want:     "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := zeroPadFileName(tt.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("zeroPadFileName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("zeroPadFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}
