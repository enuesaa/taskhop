package runfx

func New() I {
	return &Impl{
		Has: false,
	}
}

type I interface {
	Register()
	UnRegister()
}
type Impl struct {
	Has bool
}
