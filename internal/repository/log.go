package repository

import (
	"fmt"
	"log"
)

func init() {
	log.SetFlags(0)
}

// TODO
// - https://qiita.com/Imamotty/items/3fbe8ce6da4f1a653fae
// - https://medium.com/arigatobank-tech-blog/go-1-21-slog-5f0d9804204f

type LogRepositoryInterface interface {
	Printf(format string, a ...any)
	Fatalf(format string, a ...any)
	GetAll() string
}

type LogRepository struct {
	out string
}

func (repo *LogRepository) Printf(format string, a ...any) {
	text := fmt.Sprintf(format, a...)
	repo.out += text
	fmt.Print(text)
}

func (repo *LogRepository) Fatalf(format string, a ...any) {
	text := fmt.Sprintf(format, a...)
	repo.out += text
	// todo print out all logs here.
	log.Fatal(text)
}

func (repo *LogRepository) GetAll() string {
	return repo.out
}
