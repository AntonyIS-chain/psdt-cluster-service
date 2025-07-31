package logging

import "github.com/AntonyIS-chain/psdt-cluster-service/internal/core/ports"

type CompositeLogger struct {
    loggers []ports.LoggingService
}

func NewCompositeLogger(loggers ...ports.LoggingService) *CompositeLogger {
    return &CompositeLogger{loggers: loggers}
}

func (c *CompositeLogger) Info(msg string, fields map[string]interface{}) {
    for _, logger := range c.loggers {
        logger.Info(msg, fields)
    }
}

func (c *CompositeLogger) Error(msg string, fields map[string]interface{}) {
    for _, logger := range c.loggers {
        logger.Error(msg, fields)
    }
}

func (c *CompositeLogger) Debug(msg string, fields map[string]interface{}) {
    for _, logger := range c.loggers {
        logger.Debug(msg, fields)
    }
}

func (c *CompositeLogger) Warn(msg string, fields map[string]interface{}) {
    for _, logger := range c.loggers {
        logger.Warn(msg, fields)
    }
}
