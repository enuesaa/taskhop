package cmdsfile

func New() I {
	return &Impl{}
}

type I interface {
	Read(filename string) (CmdsFile, error)
}

type Impl struct {}
