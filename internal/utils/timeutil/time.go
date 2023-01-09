package timeutil

import (
	"time"

	"github.com/ryo-funaba/example_echo/internal/utils/errorutil"
)

const (
	jst            = "Asia/Tokyo"
	layout         = "20060102150405"
	layoutDateTime = "2006/01/02 15:04:05"
	offset         = 9 * 60 * 60
)

// GetJSTLocation returns a JST location.
func GetJSTLocation() *time.Location {
	return time.FixedZone(jst, offset)
}

// GetTruncatedNowJSTTime returns the value of the current time in JST with seconds truncated.
// In MySQL 5.6 or later, the decimal point in DATETIME is rounded off, so truncate the seconds in advance.
func GetTruncatedNowJSTTime() time.Time {
	jst := GetJSTLocation()
	return time.Now().In(jst).Truncate(time.Second)
}

// FormatToDateTimeString converts t to a textual representation of the time value formatted according to "2006/01/02 15:04:05".
func FormatToDateTimeString(t time.Time) string {
	return t.Format(layoutDateTime)
}

// ParseInJST converts value to time.Time as JST.
func ParseInJST(value string) (*time.Time, error) {
	t, err := time.ParseInLocation(layout, value, GetJSTLocation())
	if err != nil {
		return nil, errorutil.Errorf(errorutil.Unknown, err.Error())
	}

	return &t, nil
}
