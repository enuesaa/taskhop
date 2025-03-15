package app

import (
	"log"

	"go.uber.org/fx/fxevent"
)

func NewLogger() Logger {
	return Logger{
		Debug: false,
	}
}

type Logger struct {
	Debug bool
}

func (l *Logger) LogEvent(event fxevent.Event) {
	if l.Debug {
		l.logEventDebug(event)
	} else {
		l.logEventNormal(event)
	}
}

func (l *Logger) logEventDebug(event fxevent.Event) {
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

func (l *Logger) logEventNormal(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.Stopped:
		if e.Err != nil {
			log.Printf("Error: %s\n", e.Err.Error())
		}
	default:
		// do not log
	}
}
