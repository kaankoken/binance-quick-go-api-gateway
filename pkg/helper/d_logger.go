package helper

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

// DebugModule -> Dependency Injection for Debug logger
var DebugModule = fx.Options(fx.Provide(SetLoggerFormat))

// DLogger -> Dependency Injection Data Model for Debug logger
type DLogger struct {
	Logger *log.Logger
}

const (
	tag string = "Binance-Quick-Go-Api-Gateway -> "
)

/*
SetLoggerFormat -> Debug logger formatter initialization

[return] -> DLogger with initialized logger
*/
func SetLoggerFormat() *DLogger {
	log.SetFormatter(&log.TextFormatter{})

	return &DLogger{Logger: &log.Logger{}}
}

/*
Error -> Debug logger error logger without callback

	-> Checks whether error {nil} or {not}

[err] -> take parameter as error

[return] -> returns tag with {error} or {nil} if error does not {exist}
*/
func (logger DLogger) Error(err error) error {
	if err != nil {
		logger.Logger.Error(tag, err.Error())

		return fmt.Errorf("%s", tag+err.Error())
	}

	return nil
}

/*
ErrorWithCallback -> Debug logger error logger with callback

	-> Checks whether error {nil} or {not}

[err] -> take parameter as error
[f] -> callback method needs to be called if error {exist}

[return] -> returns tag with {error} or {nil} if error does not {exist}
*/
func (logger DLogger) ErrorWithCallback(err error, f func()) error {
	if err != nil {
		f()
		logger.Logger.Error(tag, err.Error())

		return fmt.Errorf("%s", tag+err.Error())
	}

	return nil
}

/*
Info -> Debug logger info logger without callback

[msg] -> take string message as parameter

[return] -> returns tag with {msg}
*/
func (logger DLogger) Info(msg string) string {
	logger.Logger.Infoln(tag, msg)

	return fmt.Sprintf(tag + msg)
}

/*
InfoWithCallback -> Debug logger info logger without callback

[msg] -> take string message as parameter
[f] -> callback method needs to be called

[return] -> returns tag with {msg}
*/
func (logger DLogger) InfoWithCallback(msg string, f func()) string {
	f()
	logger.Logger.Infoln(tag, msg)

	return fmt.Sprintf(tag + msg)
}
