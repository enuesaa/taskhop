package logwriter

import (
	"context"
	"io"

	"github.com/enuesaa/taskhop/app/gqlclient/adapter"
)

func New(adap *adapter.Adapter) io.Writer {
	return &LogWriter{
		adap: adap,
	}
}

type LogWriter struct {
	io.Writer

	adap *adapter.Adapter
}

func (l *LogWriter) Write(b []byte) (int, error) {
	ctx := context.Background()
	data := string(b)

	if err := l.adap.Log(ctx, data); err != nil {
		return 0, err
	}
	return len(b), nil
}
