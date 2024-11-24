package task

import "github.com/enuesaa/taskhop/internal/fs"

func GetReadme(fsrepo fs.FsRepositoryInterface) (string, error) {
	content, err := fsrepo.Read("README.md")
	if err != nil {
		return "", err
	}
	return string(content), nil
}
