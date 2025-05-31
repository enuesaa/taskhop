package gqlclient

import (
	"context"
	"errors"
	"time"
)

var ErrConnect = errors.New("failed to connect")

type sessionIDKey struct{}

func (u *UseCase) Connect(ctx context.Context) (context.Context, error) {
	if err := u.dialPolling(); err != nil {
		return nil, err
	}
	if err := u.adap.GetHealth(ctx); err != nil {
		return nil, err
	}
	sessionID, err := u.adap.Register(ctx)
	if err != nil {
		return nil, err
	}
	ctx = u.withSessionID(ctx, sessionID)

	return ctx, nil
}

func (u *UseCase) withSessionID(ctx context.Context, sessionID string) context.Context {
	return context.WithValue(ctx, sessionIDKey{}, sessionID)	
}

func (u *UseCase) getSessionID(ctx context.Context) (string, bool) {
	value, ok := ctx.Value(sessionIDKey{}).(string)
	if !ok {
		return "", false
	}
	return value, true
}

func (u *UseCase) dialPolling() error {
	for range 120 {
		if u.adap.Dial() {
			return nil
		}
		time.Sleep(5 * time.Second)
	}
	return ErrConnect
}
