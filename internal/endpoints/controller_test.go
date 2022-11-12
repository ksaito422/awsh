package endpoints_test

import (
	"errors"
	"testing"

	"awsh/internal/endpoints"
	"awsh/mock/endpoints"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/golang/mock/gomock"
)

func TestController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cases := []struct {
		name string
		in   int
	}{
		{
			name: "Undefined",
			in:   0,
		},
		{
			name: "ListBuckets",
			in:   1,
		},
		{
			name: "ListObjects",
			in:   2,
		},
		{
			name: "DownloadObject",
			in:   3,
		},
		{
			name: "StartECS",
			in:   4,
		},
		{
			name: "ecs-exec",
			in:   5,
		},
		{
			name: "StopECSTask",
			in:   6,
		},
	}

	for i, tt := range cases {
		t.Run("OK: endpoints/const.goの定数で条件分岐が正しく処理されること", func(t *testing.T) {
			cfg := aws.Config{}

			t.Run(tt.name, func(t *testing.T) {
				m := mock_endpoints.NewMockAppController(ctrl)
				m.EXPECT().
					Controller(gomock.Any(), gomock.Any()).
					DoAndReturn(
						func(cfg aws.Config, in endpoints.Operation) error {
							switch in {
							case endpoints.ListBuckets:
								return nil
							case endpoints.ListObjects:
								return nil
							case endpoints.DownloadObject:
								return nil
							case endpoints.StartECS:
								return nil
							case endpoints.ECS_EXEC:
								return nil
							case endpoints.StopECSTask:
								return nil
							}
							return errors.New(tt.name)
						}).AnyTimes()

				err := m.Controller(cfg, endpoints.Operation(tt.in))
				// 0の場合にcaseに一致した場合
				if endpoints.Undefined == endpoints.Operation(tt.in) && err == nil {
					t.Errorf("想定外のcaseに一致: want = %v, but got = %v", endpoints.Undefined, endpoints.Operation(tt.in))
				}
				// 1以上の場合にerrがリターンされた場合
				if endpoints.Undefined != endpoints.Operation(tt.in) && err != nil {
					t.Errorf("operation = %v, want = %d, but got = %d", err, i, tt.in)
				}
			})
		})

	}

	t.Run("Error:", func(t *testing.T) {
		cfg := aws.Config{}

		t.Run("想定外のエラー", func(t *testing.T) {
			m := mock_endpoints.NewMockAppController(ctrl)
			m.EXPECT().
				Controller(gomock.Any(), gomock.Any()).
				Return(errors.New("呼び出しエラー"))

			if err := m.Controller(cfg, 0); err == nil {
				t.Errorf("想定外のnil: %v", err)
			}
		})
	})
}
