package cmdsfile

func New() I {
	return &Impl{}
}

type I interface {
	Read(filename string) (CmdsFile, error)
	Validate(cmdsfile CmdsFile) error
}

type Impl struct {}
