package util

import "time"

var JST = time.FixedZone("Asia/Tokyo", 9*60*60)

func UnixToJST(unix int64, isMilli bool) time.Time {
	if isMilli {
		return time.UnixMilli(unix).In(JST)
	}
	return time.Unix(unix, 0).In(JST)
}
