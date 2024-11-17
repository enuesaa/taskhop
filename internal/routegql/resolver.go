package routegql

import (
	"github.com/enuesaa/taskhop/internal/repository"
	"github.com/enuesaa/taskhop/internal/routegql/mutation"
	"github.com/enuesaa/taskhop/internal/routegql/query"
	"github.com/enuesaa/taskhop/internal/routegql/subscription"
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

func (r *Resolver) Mutation() MutationResolver {
	resolver := mutation.MutationResolver{
		Repos: r.Repos,
	}
	return &resolver
}

func (r *Resolver) Subscription() SubscriptionResolver {
	resolver := subscription.SubscriptionResolver{
		Repos: r.Repos,
	}
	return &resolver
}
