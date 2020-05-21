package logger

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogrus(t *testing.T) {
	logger := Logrus(
		logrus.NewEntry(
			logrus.StandardLogger(),
		),
	)
	if _, ok := logger.(*wrapLogrus); !ok {
		t.Errorf("Expect wrapped logrus")
	}
	if _, ok := logger.WithError(nil).(*wrapLogrus); !ok {
		t.Errorf("Expect WithError wraps logrus")
	}
	if _, ok := logger.WithField("foo", "bar").(*wrapLogrus); !ok {
		t.Errorf("Expect WithField logrus")
	}
}
