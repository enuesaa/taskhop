package repository

func New() I {
	return &Impl{}
}

type I interface{}

type Impl struct{}
