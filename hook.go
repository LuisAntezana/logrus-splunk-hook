package splunk

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

// Hook is a logrus hook for splunk
type Hook struct {
	Client    *Client
	levels    []logrus.Level
	formatter *logrus.JSONFormatter
	async     bool
	retries   int
}

// NewHook creates new hook
// client: splunk client instance (use NewClient),
// level: log level,
// async: logs are send to splunk in async method in a different gorutine,
// retries: max retries to send logs to splunk if fails
func NewHook(client *Client, levels []logrus.Level, async bool, retries int) *Hook {
	if retries <= 0 {
		retries = 0
	}
	return &Hook{client, levels, &logrus.JSONFormatter{}, async, retries}
}

// Fire triggers a splunk event
func (h *Hook) Fire(entry *logrus.Entry) error {
	line, err := h.formatter.Format(entry)
	preparedEntry := string(line)
	if err != nil {
		return err
	}
	if h.async {
		go func() {
			err = h.Client.Log(preparedEntry, h.retries)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to fire hook: %v\n", err)
			}
		}()
	} else {
		err = h.Client.Log(preparedEntry, h.retries)
	}
	return err
}

// Levels Required for logrus hook implementation
func (h *Hook) Levels() []logrus.Level {
	return h.levels
}
