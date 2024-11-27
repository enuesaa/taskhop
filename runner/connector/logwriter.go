package connector

import "errors"

var ErrLogWriterHasNoCallback = errors.New("log writer has no callback")

type LogWriter struct {
	Callback func(data string) error
}

func (w *LogWriter) Write(b []byte) (int, error) {
	if w.Callback == nil {
		return 0, ErrLogWriterHasNoCallback
	}
	if err := w.Callback(string(b)); err != nil {
		return 0, err
	}

	return len(b), nil
}
