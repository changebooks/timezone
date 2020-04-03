package timezone

import "time"

func NowFormat(layout string) string {
	return Now().Format(layout)
}

func Now() time.Time {
	return time.Now().In(GetTimezone())
}

func Unix() int64 {
	return Now().Unix()
}

func UnixMilli() int64 {
	return Now().UnixNano() / 1e6
}

func UnixMicro() int64 {
	return Now().UnixNano() / 1e3
}

func UnixNano() int64 {
	return Now().UnixNano()
}
