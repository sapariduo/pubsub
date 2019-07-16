package consumer

import "testing"

func Test_parseToDB(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"aloha"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parseToDB()
		})
	}
}
