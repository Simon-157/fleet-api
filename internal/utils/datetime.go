package utils

import (
    "time"
)

func IsFutureDatetime(datetime time.Time) bool {
    return datetime.After(time.Now())
}
