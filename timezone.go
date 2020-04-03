package timezone

import "time"

// 缺省的时区
var timezone = time.Local

func GetTimezone() *time.Location {
	return timezone
}

func SetTimezone(name string) error {
	loc, err := time.LoadLocation(name)
	if err == nil {
		timezone = loc
	}
	return err
}

func SetUtc() {
	timezone = time.UTC
}

func SetLocal() {
	timezone = time.Local
}

func SetShanghai() error {
	return SetTimezone(Shanghai)
}

func SetNewYork() error {
	return SetTimezone(NewYork)
}
