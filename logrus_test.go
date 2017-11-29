package logrus

import (
	"bytes"
	"regexp"
	"testing"

	"github.com/sirupsen/logrus"
	"gopkg.in/birkirb/loggers.v1"
)

func TestLogrusInterface(t *testing.T) {
	var _ loggers.Contextual = NewDefaultLogger()
	var _ loggers.Advanced = &logrus.Logger{}
}

func TestLogrusLevelOutput(t *testing.T) {
	l, b := newBufferedLogrusLog()
	l.Info("This is a test")

	expectedMatch := "(?i)info.*This is a test"
	actual := b.String()
	if ok, _ := regexp.Match(expectedMatch, []byte(actual)); !ok {
		t.Errorf("Log output mismatch %s (actual) != %s (expected)", actual, expectedMatch)
	}
}

func TestLogrusLevelfOutput(t *testing.T) {
	l, b := newBufferedLogrusLog()
	l.Errorf("This is %s test", "a")

	expectedMatch := "(?i)erro.*This is a test"
	actual := b.String()
	if ok, _ := regexp.Match(expectedMatch, []byte(actual)); !ok {
		t.Errorf("Log output mismatch %s (actual) != %s (expected)", actual, expectedMatch)
	}
}

func TestLogrusLevellnOutput(t *testing.T) {
	l, b := newBufferedLogrusLog()
	l.Debugln("This is a test.", "So is this.")

	expectedMatch := "(?i)debu.*This is a test. So is this."
	actual := b.String()
	if ok, _ := regexp.Match(expectedMatch, []byte(actual)); !ok {
		t.Errorf("Log output mismatch %s (actual) != %s (expected)", actual, expectedMatch)
	}
}

func TestLogrusWithFieldsOutput(t *testing.T) {
	l, b := newBufferedLogrusLog()
	l.WithFields("test", true).Warn("This is a message.")

	expectedMatch := "(?i)warn.*This is a message.*test.*=true"
	actual := b.String()
	if ok, _ := regexp.Match(expectedMatch, []byte(actual)); !ok {
		t.Errorf("Log output mismatch %s (actual) != %s (expected)", actual, expectedMatch)
	}
}

func TestLogrusWithFieldsfOutput(t *testing.T) {
	l, b := newBufferedLogrusLog()
	l.WithFields("test", true, "Error", "serious").Errorf("This is a %s.", "message")

	expectedMatch := "(?i)erro.*This is a message.*Error.*=serious.*test.*=true"
	actual := b.String()
	if ok, _ := regexp.Match(expectedMatch, []byte(actual)); !ok {
		t.Errorf("Log output mismatch %s (actual) != %s (expected)", actual, expectedMatch)
	}
}

func newBufferedLogrusLog() (loggers.Contextual, *bytes.Buffer) {
	var b []byte
	var bb = bytes.NewBuffer(b)

	l := logrus.New()
	l.Out = bb
	l.Level = logrus.DebugLevel
	return NewLogger(l), bb
}
