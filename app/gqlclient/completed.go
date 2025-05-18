package gqlclient

import "context"

func (u *UseCase) Completed(ctx context.Context) error {
	return u.adap.Completed(ctx)
}
