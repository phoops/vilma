package log

import (
	"fmt"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(production bool, serviceName string) (*zap.SugaredLogger, error) {
	var zapLogger *zap.Logger
	var err error

	if serviceName == "" {
		return nil, fmt.Errorf("serviceName cannot be empty")
	}

	// Check for environment
	if production {
		conf := zap.NewProductionConfig()
		conf.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		zapLogger, err = conf.Build()
	} else {
		zapLogger, err = zap.NewDevelopment()
	}
	if err != nil {
		return nil, errors.Wrap(err, "error during log initialization")
	}
	return zapLogger.Sugar().With("service", serviceName), nil
}
