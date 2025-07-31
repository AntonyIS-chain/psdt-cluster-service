package logging

import (
    "github.com/sirupsen/logrus"
)

type LogrusLogger struct {
    logger *logrus.Logger
}

func NewLogrusLogger() *LogrusLogger {
    logger := logrus.New()
    logger.SetFormatter(&logrus.JSONFormatter{})
    return &LogrusLogger{logger: logger}
}

func (l *LogrusLogger) Info(msg string, fields map[string]interface{}) {
    l.logger.WithFields(fields).Info(msg)
}

func (l *LogrusLogger) Error(msg string, fields map[string]interface{}) {
    l.logger.WithFields(fields).Error(msg)
}

func (l *LogrusLogger) Debug(msg string, fields map[string]interface{}) {
    l.logger.WithFields(fields).Debug(msg)
}

func (l *LogrusLogger) Warn(msg string, fields map[string]interface{}) {
    l.logger.WithFields(fields).Warn(msg)
}
