package mutation

import (
	"github.com/enuesaa/taskhop/internal/repository"
)

type MutationResolver struct {
	Repos repository.Repos
}
