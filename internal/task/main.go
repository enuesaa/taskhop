package task

import "github.com/enuesaa/taskhop/internal/repository"

func GetReadme(repos repository.Repos) (string, error) {
	content, err := repos.Fs.Read("README.md")
	if err != nil {
		return "", err
	}
	return string(content), nil
}
