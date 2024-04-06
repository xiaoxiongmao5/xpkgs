package xtime

import "time"

func ParseTimeString(input string) (time.Time, error) {
	formats := []string{
		time.DateOnly, //"2006-01-02",
		time.DateTime, //"2006-01-02 15:04:05",
		time.RFC3339,  //"2006-01-02T15:04:05Z07:00"
		"2006-01-02T15:04:05+08:00",
		"2006-01-02T15:04:05Z",
		"2006-1-2",
		"2006-1-2 15:04:05",
		"2006-1-2T15:04:05+08:00",
		"2006-1-2T15:04:05Z",
	}

	var parsedTime time.Time
	var err error

	for _, format := range formats {
		parsedTime, err = time.Parse(format, input)
		if err == nil {
			break
		}
	}

	if err != nil {
		return parsedTime, err
	}

	return parsedTime, nil
}
