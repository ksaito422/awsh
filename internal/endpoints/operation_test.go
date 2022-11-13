package endpoints_test

import (
	"awsh/internal/endpoints"
	mock_endpoints "awsh/mock/endpoints"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestOperation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cases := []struct {
		name string
		out  string
	}{
		{
			name: "success",
			out:  endpoints.ListBuckets.String(),
		},
	}

	for i, tt := range cases {
		t.Run("OK", func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				m := mock_endpoints.NewMockAppController(ctrl)
				m.EXPECT().
					Operation().
					Return(tt.out)

				resp := m.Operation()
				if resp != tt.out {
					t.Errorf("#%d: got: %#v want: %#v", i, resp, tt.out)
				}
			})

		})
	}

	t.Run("戻り値がstringであること", func(t *testing.T) {
		m := mock_endpoints.NewMockAppController(ctrl)
		m.EXPECT().
			Operation()

		resp := m.Operation()
		want := "string"
		got := reflect.TypeOf(resp).String()
		if want != got {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
