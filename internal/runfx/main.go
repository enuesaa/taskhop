package runfx

func New() I {
	return &Impl{
		Status: StatusWaiting,
		ch: make(chan Status, 1),
	}
}

type I interface {
	GetStatus() Status
	Register() error
	Completed() error
	Subscribe() <-chan Status
}
type Impl struct {
	Status Status
	ch chan Status
}
