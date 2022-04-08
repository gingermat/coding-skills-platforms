package logs

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Message extracts the message from the provided log line.
func Message(line string) string {
	text := strings.SplitAfterN(line, ":", 2)
	return strings.TrimSpace(text[1])
}

// MessageLen counts the amount of characters (runes) in the message of the log line.
func MessageLen(line string) int {
	line = Message(line)

	if utf8.RuneCountInString(line) > 0 {
		return len([]rune(line))
	}

	return len(line)
}

// LogLevel extracts the log level string from the provided log line.
func LogLevel(line string) string {
	text := strings.SplitN(line, ":", 2)
	bytesText := []byte(text[0])

	return strings.ToLower(string(bytesText[1 : len(bytesText)-1]))
}

// Reformat reformats the log line in the format `message (logLevel)`.
func Reformat(line string) string {
	return fmt.Sprintf("%s (%s)", Message(line), LogLevel(line))
}
