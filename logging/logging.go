package logging

import "github.com/sirupsen/logrus"

func Init(logLevel LogLevel) {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(getLogrusLevel(logLevel))
  }

func Log(logMessage BasicLogMessage) {
	logEntryToConsole(*logrus.WithFields(logrus.Fields{}), logMessage.Message, logMessage.LogLevel)
}

func logEntryToConsole(entry logrus.Entry, message string, logLevel LogLevel) {
	loop:
	switch logLevel {
		case Trace:
			entry.Trace(message)
			break loop
		case Debug:
			entry.Debug(message)
			break loop
		case Info:
			entry.Info(message)
			break loop
		case Error:
			entry.Error(message)
			break loop
		case Fatal:
			entry.Fatal(message)
			break loop
		case Panic:
			entry.Panic(message)
			break loop
		default:
			entry.Warn(message)
	}
}

func getLogrusLevel(logLevel LogLevel) logrus.Level {
	switch logLevel {
	case Trace:
		return logrus.TraceLevel
	case Debug:
		return logrus.DebugLevel
	case Info:
		return logrus.InfoLevel
	case Error:
		return logrus.ErrorLevel
	case Fatal:
		return logrus.FatalLevel
	case Panic:
		return logrus.PanicLevel
	default:
		return logrus.WarnLevel
	}
}