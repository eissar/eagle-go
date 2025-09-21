package eagle_test

import (
	"testing"

	"github.com/eissar/eagle-go"
)

const BASE_URL = "http://127.0.0.1:41595"

func TestTagList(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		baseUrl string
		want    []*eagle.TagItem
		wantErr bool
	}{
		{
			name:    "test",
			baseUrl: "http://127.0.0.1:41595",
			want:    []*eagle.TagItem{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotErr := eagle.TagList(tt.baseUrl)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("TagList() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("TagList() succeeded unexpectedly")
			}

			// TODO: update the condition below to compare got with tt.want.
			// if true {
			// 	t.Errorf("TagList() = %v, want %v", got, tt.want)
			// }
		})
	}
}
