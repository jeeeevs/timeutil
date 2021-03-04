package timeutil

import "time"

//UpTime for which the service was up
func UpTime() func() int64 {
	var startTime int64 = time.Now().Unix()
	return func() int64 {
		return time.Now().Unix() - startTime
	}
}
