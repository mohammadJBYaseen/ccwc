package error

type (
	InvalidOptions struct {
	}

	FileOperationIssue struct {
		err error
	}
)

func newInvalidOptionsProvided() error {
	return &InvalidOptions{}
}

func NewFileOperationIssue(err error) error {
	return &FileOperationIssue{
		err: err,
	}
}

var (
	WRONG_OPTIONS = newInvalidOptionsProvided()
)

func (e *InvalidOptions) Error() string {
	return "Wrong options provided - please use one of the following:\nl,c,w,m"
}

func (f *FileOperationIssue) Error() string {
	return f.err.Error()
}
