package logfx

import (
	"context"
)

type logAttrKey struct{}

func (i *Impl) Use(ctx context.Context, key string, value interface{}) context.Context {
	logattr, ok := ctx.Value(logAttrKey{}).(map[string]interface{})
	if !ok {
		logattr = map[string]interface{}{}
	}
	logattr[key] = value

	return context.WithValue(ctx, logAttrKey{}, logattr)
}
