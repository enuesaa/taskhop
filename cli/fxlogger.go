package cli

import (
	"log"

	"go.uber.org/fx/fxevent"
)

func NewFxLogger(cl ICli) FxLogger {
	return FxLogger{
		cli: cl,
	}
}

type FxLogger struct {
	cli ICli
}

func (l *FxLogger) LogEvent(event fxevent.Event) {
	if l.cli.IsDebug() {
		l.logEventDebug(event)
	} else {
		l.logEventNormal(event)
	}
}

func (l *FxLogger) logEventDebug(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.Started:
		log.Printf("Started: %s\n", e)
	case *fxevent.Stopping:
		log.Printf("Stopping: %s\n", e)
	case *fxevent.Stopped:
		if e.Err != nil {
			log.Printf("Error: %s\n", e.Err.Error())
		}
	default:
		log.Printf("Event: %v\n", e)
	}
}

func (l *FxLogger) logEventNormal(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.Stopped:
		if e.Err != nil {
			log.Printf("Error: %s\n", e.Err.Error())
		}
	default:
		// do not log
	}
}
