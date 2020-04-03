package timezone

import "time"

func Format(layout string, second int64) string {
	return time.Unix(second, 0).In(GetTimezone()).Format(layout)
}

func Parse(layout string, value string) (time.Time, error) {
	return time.ParseInLocation(layout, value, GetTimezone())
}
