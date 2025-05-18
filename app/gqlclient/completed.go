package gqlclient

func (u *UseCase) Completed() error {
	return u.adap.Completed()
}
