package initialize

import (
	"fmt"
	"gin_api_02/global"
	"gin_api_02/utils"
	"io"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Zap() (logger *zap.Logger) {
	ok, _ := utils.PathExists(global.SYS_CONFIG.Zap.Director)
	if !ok {
		fmt.Printf("create %v directory\n", global.SYS_CONFIG.Zap.Director)
		_ = os.Mkdir(global.SYS_CONFIG.Zap.Director, os.ModePerm)
	}
	config := GetEncoderConfig()
	logLevel := transportLevel(global.SYS_CONFIG.Zap.Level)
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel && lvl >= logLevel
	})

	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel && lvl >= logLevel
	})

	infoWriter := getWriter(global.SYS_CONFIG.Zap.Director + "/info.log")
	warnWriter := getWriter(global.SYS_CONFIG.Zap.Director + "/error.log")

	logger = zap.New(zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(config), zapcore.AddSync(infoWriter), infoLevel),                         //将info及以下写入logPath，NewConsoleEncoder 是非结构化输出
		zapcore.NewCore(zapcore.NewConsoleEncoder(config), zapcore.AddSync(warnWriter), warnLevel),                         //warn及以上写入errPath
		zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), logLevel), //同时将日志输出到控制台，NewJSONEncoder 是结构化输出
	))

	if global.SYS_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

func GetEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:    "message",
		LevelKey:      "level",
		TimeKey:       "time",
		NameKey:       "logger",
		CallerKey:     "caller",
		StacktraceKey: global.SYS_CONFIG.Zap.StacktraceKey,
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
}

func getWriter(filename string) io.Writer {
	return &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    10,    //最大M数，超过则切割
		MaxBackups: 5,     //最大文件保留数，超过就删除最老的日志文件
		MaxAge:     30,    //保存30天
		Compress:   false, //是否压缩
	}
}

func transportLevel(level string) zapcore.Level {

	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.WarnLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
