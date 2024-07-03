package pkg

import "time"

func GenerateCurrentTimeStamp() string {
	return time.Now().Format(time.RFC3339Nano)
}

func FormatTime(data time.Time, format string) string {
	return data.Format(format)
}

func ParseTime(data string, format string) (time.Time, error) {
	return time.Parse(format, data)
}

func ParseAndFormatTime(data string, fromFormat string, toFormat string) (string, error) {
	parsedTime, err := time.Parse(fromFormat, data)
	if err != nil {
		return "", err
	}

	return parsedTime.Format(toFormat), nil
}
