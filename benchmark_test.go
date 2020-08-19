package ltsvlog_test

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"

	"github.com/hnakamur/ltsvlog"
	ltsv "github.com/hnakamur/zap-ltsv"
	"github.com/kataras/golog"
	"github.com/rs/zerolog"
	"go.uber.org/zap"
)

func BenchmarkLTSVLog(b *testing.B) {
	tmpfile, err := ioutil.TempFile("", "benchmark")
	if err != nil {
		b.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	logger := ltsvlog.NewLTSVLogger(tmpfile, false)
	for i := 0; i < b.N; i++ {
		logger.Info().String("msg", "hello").String("key1", "value1").String("key2", "value2").Log()
	}
}

// NOTE: This does not produce a proper LTSV log since a log record does not have the "time: label.
// This is used just for benchmark comparison.
func BenchmarkStandardLog(b *testing.B) {
	tmpfile, err := ioutil.TempFile("", "benchmark")
	if err != nil {
		b.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	logger := log.New(tmpfile, "", log.LstdFlags|log.Lmicroseconds)
	for i := 0; i < b.N; i++ {
		logger.Printf("level:Info\tmsg:sample log message\tkey1:%s\tkey2:%s", "value1", "value2")
	}
}

func init() {
	err := ltsv.RegisterLTSVEncoder()
	if err != nil {
		panic(err)
	}
}

func BenchmarkZapJSONProductionLog(b *testing.B) {
	tmpfile, err := ioutil.TempFile("", "benchmark")
	if err != nil {
		b.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{tmpfile.Name()}
	logger, err := cfg.Build()
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		logger.Info("sample log message", zap.String("key1", "value1"), zap.String("key2", "value2"))
	}
}

func BenchmarkZapJSONProductionLogSugar(b *testing.B) {
	tmpfile, err := ioutil.TempFile("", "benchmark")
	if err != nil {
		b.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{tmpfile.Name()}
	logger, err := cfg.Build()

	sugarLogger := logger.Sugar()

	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		sugarLogger.Infof("Success! statusCode = %s for URL %s", 200, "https://www.google.com")
	}
}

func BenchmarkZapJSONDevelopmentLog(b *testing.B) {
	tmpfile, err := ioutil.TempFile("", "benchmark")
	if err != nil {
		b.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	cfg := zap.NewDevelopmentConfig()
	cfg.OutputPaths = []string{tmpfile.Name()}
	logger, err := cfg.Build()
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		logger.Info("sample log message", zap.String("key1", "value1"), zap.String("key2", "value2"))
	}
}

func BenchmarkZapJSONDevelopmentLogSugar(b *testing.B) {
	tmpfile, err := ioutil.TempFile("", "benchmark")
	if err != nil {
		b.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	cfg := zap.NewDevelopmentConfig()
	cfg.OutputPaths = []string{tmpfile.Name()}
	logger, err := cfg.Build()

	sugarLogger := logger.Sugar()

	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		sugarLogger.Infof("Success! statusCode = %s for URL %s", 200, "https://www.google.com")
	}
}

func BenchmarkZapLTSVProductionLog(b *testing.B) {
	tmpfile, err := ioutil.TempFile("", "benchmark")
	if err != nil {
		b.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	cfg := ltsv.NewProductionConfig()
	cfg.OutputPaths = []string{tmpfile.Name()}
	logger, err := cfg.Build()
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		logger.Info("sample log message", zap.String("key1", "value1"), zap.String("key2", "value2"))
	}
}

func BenchmarkZapLTSVDevelopmentLog(b *testing.B) {
	tmpfile, err := ioutil.TempFile("", "benchmark")
	if err != nil {
		b.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	cfg := ltsv.NewDevelopmentConfig()
	cfg.OutputPaths = []string{tmpfile.Name()}
	logger, err := cfg.Build()
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		logger.Info("sample log message", zap.String("key1", "value1"), zap.String("key2", "value2"))
	}
}

func BenchmarkZerologTimestamp(b *testing.B) {
	tmpfile, err := ioutil.TempFile("", "benchmark")
	if err != nil {
		b.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	zerolog.TimeFieldFormat = ""
	logger := zerolog.New(tmpfile).With().Timestamp().Logger()
	for i := 0; i < b.N; i++ {
		logger.Info().
			Str("key1", "value1").
			Str("key2", "value2").
			Msg("sample log message")
	}
}

func BenchmarkZerologTimeSecondsFromEpoch(b *testing.B) {
	tmpfile, err := ioutil.TempFile("", "benchmark")
	if err != nil {
		b.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	zerolog.TimeFieldFormat = ""
	logger := zerolog.New(tmpfile).With().Time("time", time.Now()).Logger()
	for i := 0; i < b.N; i++ {
		logger.Info().
			Str("key1", "value1").
			Str("key2", "value2").
			Msg("sample log message")
	}
}

func BenchmarkZerologRFC3339Time(b *testing.B) {
	tmpfile, err := ioutil.TempFile("", "benchmark")
	if err != nil {
		b.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	zerolog.TimeFieldFormat = time.RFC3339
	logger := zerolog.New(tmpfile).With().Time("time", time.Now()).Logger()
	for i := 0; i < b.N; i++ {
		logger.Info().
			Str("key1", "value1").
			Str("key2", "value2").
			Msg("sample log message")
	}
}

func BenchmarkZerologRFC3339NanoTime(b *testing.B) {
	tmpfile, err := ioutil.TempFile("", "benchmark")
	if err != nil {
		b.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	zerolog.TimeFieldFormat = time.RFC3339Nano
	logger := zerolog.New(tmpfile).With().Time("time", time.Now()).Logger()
	for i := 0; i < b.N; i++ {
		logger.Info().
			Str("key1", "value1").
			Str("key2", "value2").
			Msg("sample log message")
	}
}

func BenchmarkIrisGolog(b *testing.B) {
	tmpfile, err := ioutil.TempFile("", "benchmark")
	if err != nil {
		b.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	golog.SetOutput(tmpfile)
	golog.SetLevel("debug")
	// disable time formatting because logrus and std doesn't print the time.
	// note that the time is being set-ed to time.Now() inside the golog's Log structure, same for logrus,
	// Therefore we set the time format to empty on golog test in order
	// to acomblish a fair comparison between golog and logrus.
	golog.SetTimeFormat("")

	for i := 0; i < b.N; i++ {
		golog.Infof("[%d] This is an info message", i)
	}
}

func jsonOutput(l *golog.Log) bool {
	enc := json.NewEncoder(l.Logger.Printer)
	enc.SetIndent("", "    ")
	err := enc.Encode(l)
	return err == nil
}

func BenchmarkIrisGologJSON(b *testing.B) {
	tmpfile, err := ioutil.TempFile("", "benchmark")
	if err != nil {
		b.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	golog.SetOutput(tmpfile)
	golog.Handle(jsonOutput)

	golog.SetLevel("debug")
	// disable time formatting because logrus and std doesn't print the time.
	// note that the time is being set-ed to time.Now() inside the golog's Log structure, same for logrus,
	// Therefore we set the time format to empty on golog test in order
	// to acomblish a fair comparison between golog and logrus.
	golog.SetTimeFormat("")

	for i := 0; i < b.N; i++ {
		//golog.Infof("[%d] This is an info message", i)
		golog.Debugf("This is a %s with data (debug prints the stacktrace too)", "message", golog.Fields{
			"username": "kataras",
		})
	}
}
