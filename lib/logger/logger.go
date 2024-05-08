package logger

import (
	"kevinku/go-forum/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

var cfg *config.Config

func Init() {
	cfg = config.Cfg

	var core = zapcore.NewTee(
		zapcore.NewCore(getEncoder(), getWriteSyncer(cfg.LogFile.AccessLogPath), zap.DebugLevel),
		zapcore.NewCore(getEncoder(), getWriteSyncer(cfg.LogFile.ErrorLogPath), zap.ErrorLevel),
	)
	Logger = zap.New(core, zap.AddCaller())
}

func getWriteSyncer(path string) (writeSyncer zapcore.WriteSyncer) {
	var logger = &lumberjack.Logger{
		Filename: path,
		MaxSize:  cfg.LogFile.MaxSize,
	}
	return zapcore.AddSync(logger)
}

func getEncoder() (encoder zapcore.Encoder) {
	var encoderConfig = zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05Z0700")
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder

	return zapcore.NewConsoleEncoder(encoderConfig)
}
