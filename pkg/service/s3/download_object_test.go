package s3_test

import (
	"errors"
	"testing"
	"time"

	"awsh/mock/s3/api"
	s3ser "awsh/pkg/service/s3"
	"awsh/testutil"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/golang/mock/gomock"
)

func TestDownloadObjectService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("OK", func(t *testing.T) {
		cfg := aws.Config{}
		m := mock_s3.NewMockS3Api(ctrl)
		m.EXPECT().
			ListBuckets(gomock.Any()).
			Return(&s3.ListBucketsOutput{
				Buckets: []types.Bucket{{Name: testutil.Ptr("test")}},
			}, nil)

		m.EXPECT().
			SelectBucketName(gomock.Any()).
			Return("test")

		m.EXPECT().
			ListObjects(gomock.Any(), gomock.Any()).
			Return(&s3.ListObjectsV2Output{
				Contents: []types.Object{{
					Key:          testutil.Ptr("key"),
					LastModified: testutil.Ptr(time.Now()),
					Size:         1,
					StorageClass: types.ObjectStorageClassStandard,
				}},
			}, nil)

		m.EXPECT().
			DownloadObject(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil)

		s := &s3ser.S3Service{}
		s.Api = m

		err := s.DownloadObject(cfg)
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
				Return(&s3.ListBucketsOutput{}, errors.New("aws sdkからエラーが返った"))

			// 上のmockを使用するように代入
			s := &s3ser.S3Service{}
			s.Api = m

			err := s.DownloadObject(cfg)
			if err == nil {
				t.Errorf("想定外のnil")
			}
		})

		t.Run("aws sdk: listbucketsの返り値が0件の場合", func(t *testing.T) {
			cfg := aws.Config{}
			m := mock_s3.NewMockS3Api(ctrl)
			m.EXPECT().
				ListBuckets(gomock.Any()).
				Return(&s3.ListBucketsOutput{
					Buckets: nil,
				}, nil)

			// 上のmockを使用するように代入
			s := &s3ser.S3Service{}
			s.Api = m

			err := s.DownloadObject(cfg)
			if err == nil {
				t.Errorf("想定外のnil:")
			}
		})

		t.Run("aws sdk: listobjectsからエラーが返った場合", func(t *testing.T) {
			cfg := aws.Config{}
			m := mock_s3.NewMockS3Api(ctrl)
			m.EXPECT().
				ListBuckets(gomock.Any()).
				Return(&s3.ListBucketsOutput{
					Buckets: []types.Bucket{{Name: testutil.Ptr("test")}},
				}, nil)

			m.EXPECT().
				SelectBucketName(gomock.Any()).
				Return("test")

			m.EXPECT().
				ListObjects(gomock.Any(), gomock.Any()).
				Return(nil, errors.New("aws sdkからエラーが返った"))

			// 上のmockを使用するように代入
			s := &s3ser.S3Service{}
			s.Api = m

			err := s.DownloadObject(cfg)
			if err == nil {
				t.Errorf("想定外のnil:")
			}
		})

		t.Run("aws sdk: listobjectsの返り値が0件の場合", func(t *testing.T) {
			cfg := aws.Config{}
			m := mock_s3.NewMockS3Api(ctrl)
			m.EXPECT().
				ListBuckets(gomock.Any()).
				Return(&s3.ListBucketsOutput{
					Buckets: []types.Bucket{{Name: testutil.Ptr("test")}},
				}, nil)

			m.EXPECT().
				SelectBucketName(gomock.Any()).
				Return("test")

			m.EXPECT().
				ListObjects(gomock.Any(), gomock.Any()).
				Return(&s3.ListObjectsV2Output{
					Contents: nil,
				}, nil)

			// 上のmockを使用するように代入
			s := &s3ser.S3Service{}
			s.Api = m

			err := s.DownloadObject(cfg)
			if err == nil {
				t.Errorf("想定外のnil:")
			}
		})

		t.Run("aws sdk: download objectsからエラーが返った場合", func(t *testing.T) {
			cfg := aws.Config{}
			m := mock_s3.NewMockS3Api(ctrl)
			m.EXPECT().
				ListBuckets(gomock.Any()).
				Return(&s3.ListBucketsOutput{
					Buckets: []types.Bucket{{Name: testutil.Ptr("test")}},
				}, nil)

			m.EXPECT().
				SelectBucketName(gomock.Any()).
				Return("test")

			m.EXPECT().
				ListObjects(gomock.Any(), gomock.Any()).
				Return(&s3.ListObjectsV2Output{
					Contents: []types.Object{{
						Key:          testutil.Ptr("key"),
						LastModified: testutil.Ptr(time.Now()),
						Size:         1,
						StorageClass: types.ObjectStorageClassStandard,
					}},
				}, nil)

			m.EXPECT().
				DownloadObject(gomock.Any(), gomock.Any(), gomock.Any()).
				Return(errors.New("aws sdkからエラーが返った"))

			// 上のmockを使用するように代入
			s := &s3ser.S3Service{}
			s.Api = m

			err := s.DownloadObject(cfg)
			if err == nil {
				t.Errorf("想定外のnil:")
			}
		})
	})
}
