package taskfx

func (i *TaskSrv) SubscribeEnd() <-chan bool {
	return i.endCh
}