package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

const (
	LOG_OUTPUT = "LOG_OUTPUT"
	LOG_LEVEL  = "LOG_LEVEL"
)

func init() {
	logConfig := zap.Config{
		// Define o formato que o log sera exibido.
		Encoding: "json",
		// Define onde o log será exibido.
		OutputPaths: []string{getOutputLogs()},
		// Define o level para filtrar oque será exibido de log.
		Level: zap.NewAtomicLevelAt(getLevelLogs()),
		// Configurações do formato que será exibido o log.
		EncoderConfig: zapcore.EncoderConfig{
			// Campo data e hora.
			TimeKey: "time",
			// Campo do level setado.
			LevelKey: "level",
			// Campo da mensagem setada.
			MessageKey: "message",
			// Campo que exbibe de onde o log foi chamado.
			CallerKey: "caller",
			// Formato que a data e hora será exibida.
			EncodeTime: zapcore.ISO8601TimeEncoder,
			// Transforma a label do level toda em minuscula.
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	var err error
	log, err = logConfig.Build()
	if err != nil {
		panic("erro ao inicializar logger: " + err.Error())
	}
}

// Métodos separados por nível

func Debug(message string, tags ...zap.Field) {
	log.Debug(message, tags...)
	log.Sync()
}

func Info(message string, tags ...zap.Field) {
	log.Info(message, tags...)
	log.Sync()
}

func Warn(message string, tags ...zap.Field) {
	log.Warn(message, tags...)
	log.Sync()
}

func Error(message string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.Error(message, tags...)
	log.Sync()
}

func Fatal(message string, tags ...zap.Field) {
	log.Fatal(message, tags...)
}

// Helpers de configuração

func getOutputLogs() string {
	// Aqui recuperamos o output que está cadastrado nas variáveis de ambiente
	// para definição de onde será mostrado o log

	output := strings.ToLower(strings.TrimSpace(os.Getenv(LOG_OUTPUT)))
	if output == "" {
		return "stdout"
	}
	return output
}

func getLevelLogs() zapcore.Level {
	// Aqui recuperaremos o level que está cadastrado nas variaveis de ambiente
	// para controle do que irá aparecer para o usuario no console

	switch strings.ToLower(strings.TrimSpace(os.Getenv(LOG_LEVEL))) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}
