package routegql

import (
	"github.com/enuesaa/taskhop/internal/repository"
	"github.com/enuesaa/taskhop/internal/routegql/query"
)

type Resolver struct {
	Repos repository.Repos
}

func (r *Resolver) Query() QueryResolver {
	resolver := query.QueryResolver{
		Repos: r.Repos,
	}
	return &resolver
}
