package query

import "github.com/enuesaa/taskhop/lib"

func New(li *lib.Lib) *QueryResolver {
	return &QueryResolver{
		li: li,
	}
}

type QueryResolver struct {
	li *lib.Lib
}
