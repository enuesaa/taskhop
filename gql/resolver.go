package gql

import (
	"github.com/enuesaa/taskhop/internal/repository"
	"github.com/enuesaa/taskhop/gql/mutation"
	"github.com/enuesaa/taskhop/gql/query"
	"github.com/enuesaa/taskhop/gql/subscription"
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
