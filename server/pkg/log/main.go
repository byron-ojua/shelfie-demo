package log

import "go.uber.org/zap"

type LoggerOptions struct {
	Level      zap.AtomicLevel
	OutputFile string
}

func New(params LoggerOptions) (*zap.SugaredLogger, error) {
	cfg := zap.NewDevelopmentConfig()

	// Set the log level
	cfg.Level = params.Level

	// Set the output file
	if params.OutputFile != "" {
		cfg.OutputPaths = append(cfg.OutputPaths, params.OutputFile)
	}

	// Build the logger
	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	// Use the logger
	sugar := logger.Sugar()

	return sugar, nil
}
