package repository

import (
	"fmt"
	"log"
)

type LogMockRepository struct {
	out string
}

func (repo *LogMockRepository) Printf(format string, a ...any) {
	text := fmt.Sprintf(format, a...)
	repo.out += text
}

func (repo *LogMockRepository) Fatalf(format string, a ...any) {
	text := fmt.Sprintf(format, a...)
	repo.out += text
	log.Fatal(text)
}

func (repo *LogMockRepository) GetAll() string {
	return repo.out
}
