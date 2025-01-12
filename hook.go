package splunk

import (
	"github.com/sirupsen/logrus"
)

// Hook is a logrus hook for splunk
type Hook struct {
	Client    *Client
	levels    []logrus.Level
	formatter *logrus.JSONFormatter
}

// NewHook creates new hook
// client - splunk client instance (use NewClient)
// level - log level
func NewHook(client *Client, levels []logrus.Level) *Hook {

	return &Hook{client, levels, &logrus.JSONFormatter{}}
}

// Fire triggers a splunk event
func (h *Hook) Fire(entry *logrus.Entry) error {
	line, err := h.formatter.Format(entry)
	preparedEntry := string(line)

	if err != nil {
		return err
	}

	err = h.Client.Log(
		preparedEntry,
	)
	return err
}

// Levels Required for logrus hook implementation
func (h *Hook) Levels() []logrus.Level {
	return h.levels

}
