package middleware

import (
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
	var writeSyncer = getWriteSyncer()
	var encoder = getEncoder()
	var level = new(zapcore.Level)

	var core = zapcore.NewCore(encoder, writeSyncer, level)
	Logger = zap.New(core, zap.AddCaller())
}

func getWriteSyncer() (writeSyncer zapcore.WriteSyncer) {
	var logger = &lumberjack.Logger{
		Filename: config.Cfg.LogFile.Path,
		MaxSize:  config.Cfg.LogFile.MaxSize,
	}
	return zapcore.AddSync(logger)
}

func getEncoder() (encoder zapcore.Encoder) {
	var encoderConfig = zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder

	return zapcore.NewJSONEncoder(encoderConfig)
}

func GinLogger(logger *zap.Logger) (handler gin.HandlerFunc) {
	return func(ctx *gin.Context) {
		var start = time.Now()
		var path = ctx.Request.URL.Path
		var query = ctx.Request.URL.RawQuery
		ctx.Next()

		var cost = time.Since(start)
		logger.Sugar().Info(
			path,
			"status", ctx.Writer.Status(),
			"method", ctx.Request.Method,
			"path", path,
			"query", query,
			"ip", ctx.ClientIP(),
			"user-agent", ctx.Request.UserAgent(),
			"errors", ctx.Errors.ByType(gin.ErrorTypePrivate).String(),
			"cost", cost,
		)
	}
}

func GinRecovery(logger *zap.Logger, stack bool) (handler gin.HandlerFunc) {
	return func(ctx *gin.Context) {
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
					logger.Sugar().Error(
						"[Recovery from panic]",
						"error", err,
						"request", string(request),
						"stack", string(debug.Stack()),
					)
				} else {
					logger.Sugar().Error(
						"[Recovery from panic]",
						"error", err,
						"request", string(request),
					)
				}
				ctx.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		ctx.Next()
	}
}
