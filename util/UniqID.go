package util

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

func CreateUniqID() string {
	timeNow := time.Now().Format("060102150405.000")
	uuID := strings.Replace(uuid.New().String(), "-", "", -1)
	uniqID := strings.Replace(timeNow, ".", "", 1) + strings.ToUpper(uuID[:5])
	return uniqID
}
