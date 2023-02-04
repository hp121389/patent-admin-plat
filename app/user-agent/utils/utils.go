package utils

import "time"

func FormatCurrentTime() string {
	t := time.Now()
	time := t.Format("2006年1月2日")
	return time
}
