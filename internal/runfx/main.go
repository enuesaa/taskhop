package runfx

func New() I {
	return &Impl{
		Status: StatusWaiting,
	}
}

type I interface {
	GetStatus() Status
	Register() error
	UnRegister() error
}
type Impl struct {
	Status Status
}
