package logfx

import "log"

func init() {
	log.SetFlags(0)
}

func NewRepository() IRepository {
	return &Repository{}
}

type IRepository interface {
	Print(text string)
}

type Repository struct {}

func (r *Repository) Print(text string) {
	log.Print(text)
}
