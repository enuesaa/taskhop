package gqlclient

func (u *UseCase) Log(text string) {
	u.LogWriter.Write([]byte(text))
}
