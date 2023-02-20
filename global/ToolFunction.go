package global

import "time"

func NeverStop() {
	for {
		time.Sleep(time.Hour)
	}
}
