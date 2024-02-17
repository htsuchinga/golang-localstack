package datetime

import "time"

var TimeNow = time.Now

var locationsJST *time.Location

func GetLocationJST() *time.Location {
	if locationsJST == nil {
		locationsJST = time.FixedZone("JST", 9*60*60)
	}
	return locationsJST
}

func NowInJST() time.Time {
	return TimeNow().In(GetLocationJST())
}

func ParseInJST(layout string, value string) (time.Time, error) {
	return time.ParseInLocation(layout, value, GetLocationJST())
}
