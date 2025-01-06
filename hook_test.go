package belajargolanglogging

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

type sampleHook struct {
}

func (s *sampleHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}
}

func (s *sampleHook) Fire(entry *logrus.Entry) error {
	fmt.Println("Sample Hook", entry.Level)
	return nil
}

func TestHook(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.AddHook(&sampleHook{})

	// logger.Info("Hello")
	logger.Warn("HAII")

}
