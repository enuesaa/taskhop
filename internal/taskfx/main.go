package taskfx

func New() I {
	return &Impl{
		Workdir: ".",
	}
}

type I interface {
	Use(workdir string)
	Read() (Task, error)
	Validate(task Task) error
}

type Impl struct{
	Workdir string
}
