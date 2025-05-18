package mutation

import "github.com/enuesaa/taskhop/lib"

func New(li *lib.Lib) *MutationResolver {
	return &MutationResolver{
		li: li,
	}
}

type MutationResolver struct {
	li *lib.Lib
}
