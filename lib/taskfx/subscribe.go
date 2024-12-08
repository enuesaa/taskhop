package taskfx

func (i *Impl) Subscribe() <-chan Status {
	return i.ch
}
