package usecase

func (u *UseCase) Use(workdir string) {
	u.Workdir = workdir
}
