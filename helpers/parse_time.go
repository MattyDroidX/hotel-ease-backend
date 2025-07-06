package helpers

import (
	"time"
)

// ParseHTMLDateTime tenta RFC 3339 e, se falhar, tenta “YYYY-MM-DDTHH:MM”.
func ParseHTMLDateTime(s string) (time.Time, error) {
	// 1) primeiro tenta o formato completo RFC 3339
	if t, err := time.Parse(time.RFC3339, s); err == nil {
		return t, nil
	}
	// 2) tenta o formato sem segundos / fuso
	const short = "2006-01-02T15:04"
	return time.ParseInLocation(short, s, time.Local)
}
