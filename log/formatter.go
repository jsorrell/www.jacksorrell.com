package log

/* Heavily copied from logrus/text_formatter */

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

const timestampFormat = "2006-01-02 15:04:05"

// TextFormatter formats logs into text
type TextFormatter struct{}

// Format renders a single log entry
// TODO format stacktrace fields specially
func (f *TextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	stacktrace, hasStack := entry.Data["stacktrace"]
	if hasStack {
		delete(entry.Data, "stacktrace")
	}

	// Sort keys
	keys := make([]string, 0, len(entry.Data))
	for k := range entry.Data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Create or use given buffer
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	// Level
	levelColor := getLevelColor(entry.Level)
	levelText := getLevelText(entry.Level)
	levelColor.Fprintf(b, "%-5s", levelText)
	b.WriteByte(' ')
	// Time
	color.New(color.FgHiMagenta).Fprint(b, entry.Time.Format(timestampFormat))
	b.WriteByte(' ')
	// Message
	// Remove a single newline if it already exists in the message to keep
	// the behavior of formatter the same as the stdlib log package
	entry.Message = strings.TrimSuffix(entry.Message, "\n")
	fmt.Fprintf(b, "%-44s", entry.Message)
	b.WriteByte(' ')

	for _, k := range keys {
		f.appendKeyValue(b, k, entry.Data[k], levelColor)
	}

	if hasStack {
		color.New(color.FgHiWhite).Fprintf(b, "\n\n%s", stacktrace)
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}

func getLevelColor(level logrus.Level) *color.Color {
	switch level {
	case logrus.DebugLevel:
		return color.New(color.FgWhite)
	case logrus.WarnLevel:
		return color.New(color.FgYellow)
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		return color.New(color.FgHiRed)
	default:
		return color.New(color.FgBlue)
	}
}

func getLevelText(level logrus.Level) string {
	switch level {
	case logrus.DebugLevel:
		return "DEBUG"
	case logrus.InfoLevel:
		return "INFO"
	case logrus.WarnLevel:
		return "WARN"
	case logrus.ErrorLevel:
		return "ERROR"
	case logrus.FatalLevel:
		return "FATAL"
	case logrus.PanicLevel:
		return "PANIC"
	default:
		return "?????"
	}
}

func (f *TextFormatter) appendKeyValue(b *bytes.Buffer, k string, v interface{}, c *color.Color) {
	b.WriteByte(' ')
	c.Fprint(b, k)
	b.WriteByte('=')
	f.appendValue(b, v)
}

func (f *TextFormatter) appendValue(b *bytes.Buffer, value interface{}) {
	switch value.(type) {
	case ColoredField:
		v := value.(ColoredField)
		str := fmt.Sprint(v.field)
		if !f.needsQuoting(str) {
			v.color.Fprint(b, str)
		} else {
			v.color.Fprintf(b, "%q", str)
		}
	default:
		str := fmt.Sprint(value)
		if !f.needsQuoting(str) {
			b.WriteString(str)
		} else {
			fmt.Fprintf(b, "%q", str)
		}
	}
}

func (f *TextFormatter) needsQuoting(text string) bool {
	if len(text) == 0 {
		return true
	}
	for _, ch := range text {
		if !((ch >= 'a' && ch <= 'z') ||
			(ch >= 'A' && ch <= 'Z') ||
			(ch >= '0' && ch <= '9') ||
			ch == '-' || ch == '.' || ch == '_' || ch == '/' || ch == '@' || ch == '^' || ch == '+') {
			return true
		}
	}
	return false
}

// ColoredField pass as field to colorize the output to the given color
type ColoredField struct {
	color *color.Color
	field interface{}
}
