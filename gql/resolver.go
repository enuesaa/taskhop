package gql

import (
	"github.com/enuesaa/taskhop/gql/mutation"
	"github.com/enuesaa/taskhop/gql/query"
	"github.com/enuesaa/taskhop/gql/subscription"
	"github.com/enuesaa/taskhop/internal/fs"
)

type Resolver struct {
	Fs fs.FsRepositoryInterface
}

func (r *Resolver) Query() QueryResolver {
	resolver := query.QueryResolver{
		Fs: r.Fs,
	}
	return &resolver
}

func (r *Resolver) Mutation() MutationResolver {
	resolver := mutation.MutationResolver{}
	return &resolver
}

func (r *Resolver) Subscription() SubscriptionResolver {
	resolver := subscription.SubscriptionResolver{}
	return &resolver
}
