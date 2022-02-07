package logging

type LogLevel string

const (
	Trace LogLevel = "Trace"
	Debug LogLevel = "Debug"
	Info LogLevel = "Info"
	Warn LogLevel = "Warn"
	Error LogLevel = "Error"
	Fatal LogLevel = "Fatal"
	Panic LogLevel = "Panic"
)

type BasicLogMessage struct {
	LogLevel LogLevel
	Message string
}