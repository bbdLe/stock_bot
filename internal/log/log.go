package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var (
	Logger *zap.Logger // 普通日志
	StatLogger *zap.Logger	// 统计日志
)

func newCore(filePath string, level zapcore.Level, maxSize int, maxBackUps int, maxAge int, compress bool) zapcore.Core {
	hook := lumberjack.Logger{
		Filename : filePath,
		MaxSize : maxSize,
		MaxAge : maxAge,
		MaxBackups : maxBackUps,
		Compress : compress,
	}

	// 日志等级
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)

	// 编码器配置
	encoderConfig := zapcore.EncoderConfig{
		TimeKey: "time",
		LevelKey: "level",
		NameKey: "logger",
		CallerKey: "linenum",
		MessageKey: "msg",
		StacktraceKey: "stacktrace",
		LineEnding: zapcore.DefaultLineEnding,
		EncodeLevel: zapcore.LowercaseLevelEncoder,
		EncodeTime: zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeName: zapcore.FullNameEncoder,
	}
	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		atomicLevel,
	)
}

func NewLogger(filePath string, level zapcore.Level, maxSize int, maxBackUps int, maxAge int, compress bool) *zap.Logger {
	core := newCore(filePath, level, maxSize, maxBackUps, maxAge, compress)
	return zap.New(core, zap.AddCaller(), zap.Development())
}

func init() {
	Logger = NewLogger("../log/stock.log", zapcore.DebugLevel, 500, 5, 30, true)
	StatLogger = NewLogger("../log/stock.log", zapcore.DebugLevel, 500, 5, 30, true)
}
