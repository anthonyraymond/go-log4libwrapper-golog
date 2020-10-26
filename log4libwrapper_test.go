package log4libwrapper

import (
	"github.com/kataras/golog"
	"io"
	"os"
	"strings"
	"testing"
)

func TestWrapLogrusLogger_MustLog(t *testing.T) {
	logReceiver := &strings.Builder{}

	goLogLogger := golog.New()
	goLogLogger.SetTimeFormat("")
	goLogLogger.SetLevel(golog.DebugLevel.String())
	goLogLogger.SetOutput(io.MultiWriter(logReceiver, os.Stdout))

	logger := WrapGologLogger(goLogLogger)
	logger.Debug("coucou")
	logger.Info("coucou")
	logger.Warn("coucou")
	logger.Error("coucou")
	logger.Debug("coucou", "joe", "la bidouille")
	logger.Info("coucou", "joe", "la bidouille")
	logger.Warn("coucou", "joe", "la bidouille")
	logger.Error("coucou", "joe", "la bidouille")
	logger.Debugf("coucou %s :)", "joe")
	logger.Infof("coucou %s :)", "joe")
	logger.Warnf("coucou %s :)", "joe")
	logger.Errorf("coucou %s :)", "joe")
	logger.Debugf("coucou %s %d :)", "joe", 12)
	logger.Infof("coucou %s %d :)", "joe", 12)
	logger.Warnf("coucou %s %d :)", "joe", 12)
	logger.Errorf("coucou %s %d :)", "joe", 12)

	expected := []string{
		"[DBUG] coucou",
		"[INFO] coucou",
		"[WARN] coucou",
		"[ERRO] coucou",
		"[DBUG] coucoujoela bidouille",
		"[INFO] coucoujoela bidouille",
		"[WARN] coucoujoela bidouille",
		"[ERRO] coucoujoela bidouille",
		"[DBUG] coucou joe :)",
		"[INFO] coucou joe :)",
		"[WARN] coucou joe :)",
		"[ERRO] coucou joe :)",
		"[DBUG] coucou joe 12 :)",
		"[INFO] coucou joe 12 :)",
		"[WARN] coucou joe 12 :)",
		"[ERRO] coucou joe 12 :)",
		"",
	}

	if !strings.EqualFold(strings.Join(expected, "\n"), logReceiver.String()) {
		t.Fatalf("log output is not correct, expected:\n%s\nactual:\n%s", strings.Join(expected, "\n"), logReceiver.String())
	}
}
