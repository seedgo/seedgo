package seedgo

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.SugaredLogger

func InitLogger() {
	loggerList := viper.GetStringMap("logger")

	if len(loggerList) == 0 {
		// register default console logger
		l, _ := zap.NewDevelopment()
		Logger = l.Sugar()
		return
	}

	// parse loggerList
	var cList []zapcore.Core
	for loggerName, _ := range loggerList {
		c, err := parseLogger(loggerName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "parse logger %s failed, err: %s", loggerName, err)
			continue
		}

		cList = append(cList, c)
	}

	core := zapcore.NewTee(cList...)

	// Build the logger
	Logger = zap.New(core).WithOptions(zap.WithCaller(true), zap.AddCallerSkip(1)).Sugar()
}

func concatTid(ctx *gin.Context, template string) string {
	if ctx != nil {
		tid := ctx.GetString("tid")
		if len(tid) > 0 {
			template = " tid: " + tid + " " + template
		}
	}

	return template
}

func Infof(ctx *gin.Context, template string, arg ...interface{}) {
	Logger.Infof(concatTid(ctx, template), arg...)
}

func Errorf(ctx *gin.Context, template string, arg ...interface{}) {
	Logger.Errorf(concatTid(ctx, template), arg...)
}

func Debugf(ctx *gin.Context, template string, arg ...interface{}) {
	Logger.Debugf(concatTid(ctx, template), arg...)
}

func Warnf(ctx *gin.Context, template string, arg ...interface{}) {
	Logger.Warnf(concatTid(ctx, template), arg...)
}

func parseLogger(name string) (zapcore.Core, error) {
	cnf := viper.Sub("logger." + name)
	driverName := cnf.GetString("driver")
	switch driverName {
	case "console":
		return parseConsoleConf(cnf), nil
	case "file":
		return parseFileConf(cnf), nil
	default:
		return parseConsoleConf(cnf), nil
	}
}

func parseFileConf(cnf *viper.Viper) zapcore.Core {
	// Create a new lumberjack logger for file rotation
	lumberjackLogger := &lumberjack.Logger{
		Filename:   cnf.GetString("filename"),
		MaxSize:    cnf.GetInt("maxsize"),    // Max megabytes before rotation
		MaxBackups: cnf.GetInt("maxbackups"), // Max number of old log files to keep
		MaxAge:     cnf.GetInt("maxage"),     // Max days to retain old log files
		Compress:   cnf.GetBool("compress"),
	}

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	// Create a new zap core that writes to both stdout and the lumberjack logger
	fileEncoder := zapcore.NewJSONEncoder(config)
	return zapcore.NewCore(fileEncoder, zapcore.AddSync(lumberjackLogger), getZapLevel(cnf.GetString("level")))
}

func parseConsoleConf(cnf *viper.Viper) zapcore.Core {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(config)

	return zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), getZapLevel(cnf.GetString("level")))
}

func getZapLevel(levelName string) zapcore.Level {
	switch levelName {
	case "info":
		return zap.InfoLevel
	default:
		return zap.DebugLevel
	}

}
