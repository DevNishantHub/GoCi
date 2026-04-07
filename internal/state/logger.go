package state

import "time"

type LogEntry struct {
	RunIds    string
	Stage     string
	Job       string
	Level     string
	Message   string
	TimeStamp time.Time
}
