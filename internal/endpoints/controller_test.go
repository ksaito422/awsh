package endpoints_test

import (
	"errors"
	"testing"

	"awsh/internal/endpoints"
	"awsh/mock/ecs/service"
	"awsh/mock/s3/service"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/golang/mock/gomock"
)

func TestController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("OK", func(t *testing.T) {
		mockS3 := mock_s3.NewMockS3Servicer(ctrl)
		mockEcs := mock_ecs.NewMockECSServicer(ctrl)

		cases := []struct {
			name string
			in   int
			f    *gomock.Call
		}{
			{
				name: "Undefined",
				in:   0,
			},
			{
				name: "ListBuckets",
				in:   1,
				f:    mockS3.EXPECT().ListBuckets(gomock.Any()).Return(nil),
			},
			{
				name: "ListObjects",
				in:   2,
				f:    mockS3.EXPECT().ListObjects(gomock.Any()).Return(nil),
			},
			{
				name: "DownloadObject",
				in:   3,
				f:    mockS3.EXPECT().DownloadObject(gomock.Any()).Return(nil),
			},
			{
				name: "StartECS",
				in:   4,
				f:    mockEcs.EXPECT().StartEcs(gomock.Any()).Return(nil),
			},
			{
				name: "ecs-exec",
				in:   5,
				f:    mockEcs.EXPECT().EcsExec(gomock.Any()).Return(nil),
			},
			{
				name: "StopECSTask",
				in:   6,
				f:    mockEcs.EXPECT().StopEcsTask(gomock.Any()).Return(nil),
			},
		}

		for i, tt := range cases {
			t.Run("endpoints/const.goの定数で条件分岐が正しく処理されること", func(t *testing.T) {
				cfg := aws.Config{}

				t.Run(tt.name, func(t *testing.T) {
					// switchの中で呼ぶ関数のmock代入
					r := &endpoints.Route{}
					r.S3 = mockS3
					r.ECS = mockEcs

					err := r.Controller(cfg, endpoints.Operation(tt.in))
					// 0の場合にcaseに一致した場合
					if endpoints.Undefined == endpoints.Operation(tt.in) && err == nil {
						t.Errorf("想定外のcaseに一致: #%d, want: %#v, got = %#v", i, endpoints.Undefined, endpoints.Operation(tt.in))
					}
					// 1以上の場合にerrがリターンされた場合
					if endpoints.Undefined != endpoints.Operation(tt.in) && err != nil {
						t.Errorf("#%d, want: %d, got: %d", i, err, tt.in)
					}
				})
			})

		}
	})

	t.Run("Error", func(t *testing.T) {
		mockS3 := mock_s3.NewMockS3Servicer(ctrl)
		mockEcs := mock_ecs.NewMockECSServicer(ctrl)

		cases := []struct {
			name string
			in   int
			f    *gomock.Call
		}{
			{
				name: "Undefined",
				in:   0,
			},
			{
				name: "ListBuckets",
				in:   1,
				f:    mockS3.EXPECT().ListBuckets(gomock.Any()).Return(errors.New("error")),
			},
			{
				name: "ListObjects",
				in:   2,
				f:    mockS3.EXPECT().ListObjects(gomock.Any()).Return(errors.New("error")),
			},
			{
				name: "DownloadObject",
				in:   3,
				f:    mockS3.EXPECT().DownloadObject(gomock.Any()).Return(errors.New("error")),
			},
			{
				name: "StartECS",
				in:   4,
				f:    mockEcs.EXPECT().StartEcs(gomock.Any()).Return(errors.New("error")),
			},
			{
				name: "ecs-exec",
				in:   5,
				f:    mockEcs.EXPECT().EcsExec(gomock.Any()).Return(errors.New("error")),
			},
			{
				name: "StopECSTask",
				in:   6,
				f:    mockEcs.EXPECT().StopEcsTask(gomock.Any()).Return(errors.New("error")),
			},
		}

		for i, tt := range cases {
			t.Run("Errorが返ること", func(t *testing.T) {
				cfg := aws.Config{}

				// switchの中で呼ぶ関数のmock代入
				r := &endpoints.Route{}
				r.S3 = mockS3
				r.ECS = mockEcs

				err := r.Controller(cfg, endpoints.Operation(tt.in))

				if err == nil {
					t.Errorf("想定外のnil: #%d", i)
				}
			})
		}
	})
}
