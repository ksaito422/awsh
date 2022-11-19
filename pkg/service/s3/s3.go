package s3

import "golang.org/x/xerrors"

var (
	noBucket = xerrors.New("Bucket does not exist.")
	noObject = xerrors.New("No object exists in the bucket.")
)

const (
	OBJECT_DOWNLOAD = "Object download succeeded!"
)
