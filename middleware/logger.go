package middleware

import (
	"fmt"
	"kevinku/go-forum/config"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

func init() {
	var core = zapcore.NewTee(
		zapcore.NewCore(getEncoder(), getWriteSyncer(config.Cfg.LogFile.AccessLogPath), zap.DebugLevel),
		zapcore.NewCore(getEncoder(), getWriteSyncer(config.Cfg.LogFile.ErrorLogPath), zap.ErrorLevel),
	)
	Logger = zap.New(core, zap.AddCaller())
}

func getWriteSyncer(path string) (writeSyncer zapcore.WriteSyncer) {
	var logger = &lumberjack.Logger{
		Filename: path,
		MaxSize:  config.Cfg.LogFile.MaxSize,
	}
	return zapcore.AddSync(logger)
}

func getEncoder() (encoder zapcore.Encoder) {
	var encoderConfig = zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05Z0700")
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder

	return zapcore.NewConsoleEncoder(encoderConfig)
}

func GinLogger(logger *zap.Logger) (handler gin.HandlerFunc) {
	return func(ctx *gin.Context) {
		var loggerWithoutCaller *zap.Logger = zap.New(logger.Core())

		var start = time.Now()
		ctx.Next()
		var cost = time.Since(start)

		var output = fmt.Sprintf("| %3d | %13s | %15s | %-8s %s",
			ctx.Writer.Status(),
			cost,
			ctx.ClientIP(),
			ctx.Request.Method,
			ctx.Request.URL.Path)

		loggerWithoutCaller.Info(
			output,
			zap.String("query", ctx.Request.URL.RawQuery),
			zap.String("user-agent", ctx.Request.UserAgent()),
			zap.String("errors", ctx.Errors.ByType(gin.ErrorTypePrivate).String()),
		)
	}
}

func GinRecovery(logger *zap.Logger, stack bool) (handler gin.HandlerFunc) {
	return func(ctx *gin.Context) {
		var loggerWithoutCaller *zap.Logger = zap.New(logger.Core())

		defer func() {
			if err := recover(); err != nil {
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "brokenpipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				var request []byte
				request, _ = httputil.DumpRequest(ctx.Request, false)
				if brokenPipe {
					logger.Sugar().Error(
						"error", err,
						"request", string(request),
					)

					ctx.Error(err.(error))
					ctx.Abort()
					return
				}

				if stack {
					loggerWithoutCaller.Error(
						"[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(request)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					loggerWithoutCaller.Error(
						"[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(request)),
					)
				}
				ctx.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		ctx.Next()
	}
}
