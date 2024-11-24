package usecase

func (u *Usecase) GetReadme() (string, error) {
	content, err := u.fs.Read("README.md")
	if err != nil {
		return "", err
	}
	return string(content), nil
}
