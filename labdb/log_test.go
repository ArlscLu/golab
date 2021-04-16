package labdb

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestFormat(t *testing.T) {
	first := TaskConfig{
		GroupID: 1,
	}
	_ = gormDb.First(&first)
	jsonFormat := &logrus.JSONFormatter{
		DisableTimestamp:  true,
		DisableHTMLEscape: true,
		PrettyPrint:       true,
	}
	_ = &logrus.TextFormatter{
		ForceColors:               true,
		DisableColors:             false,
		ForceQuote:                false,
		DisableQuote:              true,
		EnvironmentOverrideColors: false,
		DisableTimestamp:          false,
		FullTimestamp:             false,
		TimestampFormat:           "",
		DisableSorting:            true,
		DisableLevelTruncation:    false,
		PadLevelText:              false,
		QuoteEmptyFields:          false,
	}

	logrus.SetFormatter(jsonFormat)
	logrus.Print(first)
}
