package taskfx

func New() I {
	return &Impl{}
}

type I interface {
	Read(filename string) (Task, error)
	Validate(task Task) error
}

type Impl struct{}
