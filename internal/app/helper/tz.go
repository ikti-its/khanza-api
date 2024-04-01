package helper

import (
	"github.com/fathoor/simkes-api/internal/app/exception"
	"time"
)

func FormatTime(t time.Time, layout string) string {
	return t.In(time.FixedZone("WIB", 7*60*60)).Format(layout)
}

func ParseTime(s, layout string) time.Time {
	t, err := time.ParseInLocation(layout, s, time.FixedZone("WIB", 7*60*60))
	if err != nil {
		panic(&exception.BadRequestError{
			Message: "Invalid date format",
		})
	}

	return t
}

func ParseNow() time.Time {
	return time.Now().In(time.FixedZone("WIB", 7*60*60))
}
