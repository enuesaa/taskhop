package gqlclient

import (
	"context"
	"io"

	"github.com/enuesaa/taskhop/app/gqlserver/model"
)

type LogWriter struct {
	io.Writer

	gql GQLClient
}

func (l *LogWriter) Write(b []byte) (int, error) {
	data := string(b)

	input := model.LogInput{
		Output: data,
	}
	if _, err := l.gql.Log(context.Background(), input); err != nil {
		return 0, err
	}
	return len(b), nil
}
