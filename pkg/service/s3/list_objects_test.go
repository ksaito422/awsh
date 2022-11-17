package s3_test

import (
	"errors"
	"testing"

	"awsh/mock/s3/api"
	s3ser "awsh/pkg/service/s3"
	"awsh/testutil"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/golang/mock/gomock"
)

func TestListObjectsService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("OK", func(t *testing.T) {
		t.Skip() // TODO: プロンプト選択のSelectBucketNameをモックする
		cfg := aws.Config{}
		m := mock_s3.NewMockS3Api(ctrl)
		m.EXPECT().
			ListBuckets(gomock.Any()).
			Return(&s3.ListBucketsOutput{
				Buckets: []types.Bucket{{Name: testutil.Ptr("test")}},
			}, nil)

		m.EXPECT().
			ListObjects(gomock.Any(), gomock.Any()).
			Return(nil, nil)

		s := &s3ser.S3Service{}
		s.Api = m

		err := s.ListObjects(cfg)
		if err != nil {
			t.Errorf("想定外のError: %v", err)
		}
	})

	t.Run("Error", func(t *testing.T) {
		t.Run("aws sdk: listbucketsからエラーが返った場合", func(t *testing.T) {
			cfg := aws.Config{}
			m := mock_s3.NewMockS3Api(ctrl)
			m.EXPECT().
				ListBuckets(gomock.Any()).
				Return(&s3.ListBucketsOutput{}, errors.New("error"))

			// 上のmockを使用するように代入
			s := &s3ser.S3Service{}
			s.Api = m

			err := s.ListBuckets(cfg)
			if err == nil {
				t.Errorf("想定外のnil: %v", err)
			}
		})

		t.Run("aws sdk: listbucketsの返り値が0件の場合", func(t *testing.T) {
			cfg := aws.Config{}
			m := mock_s3.NewMockS3Api(ctrl)
			m.EXPECT().
				ListBuckets(gomock.Any()).
				Return(&s3.ListBucketsOutput{
					Buckets: []types.Bucket{},
				}, nil)

			// 上のmockを使用するように代入
			s := &s3ser.S3Service{}
			s.Api = m

			err := s.ListBuckets(cfg)
			if err == nil {
				t.Errorf("想定外のnil:")
			}
		})

		// TODO: プロンプトモックしたら、その先のテスト書く
	})
}
