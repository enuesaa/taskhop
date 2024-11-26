package runfx

func (i *Impl) Forget() {
	i.Status = StatusWaiting
}
