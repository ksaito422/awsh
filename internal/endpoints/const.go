package endpoints

type Resource string

const (
	S3  = Resource("S3")
	ECS = Resource("ECS")
)

type Operation int

// NOTE: 操作を追加する場合、const,varに命名を追加すること
const (
	Undefined Operation = iota
	ListBuckets
	ListObjects
	DownloadObject
	StartECS
	ECS_EXEC
	StopECSTask
)

var operation = [7]string{
	"Undefined",
	"ListBuckets",
	"ListObjects",
	"DownloadObject",
	"StartECS",
	"ecs-exec",
	"StopECSTask",
}

// Operation定数からoperation配列の文字列を返す
func (o Operation) String() string {
	return operation[o]
}

// 引数の文字列からoperation配列のindexを返す
func Index(f string) Operation {
	for i, v := range operation {
		if v == f {
			return Operation(i)
		}
	}

	return Undefined
}
