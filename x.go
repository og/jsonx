package xjson

import (
	"time"
)

var emptyObjectAndArrayNotNull = true
func emptyObjectString () string {
	if emptyObjectAndArrayNotNull {
		return "{}"
	} else {
		return "null"
	}
}
func emptyArrayString () string {
	if emptyObjectAndArrayNotNull {
		return "[]"
	} else {
		return "null"
	}
}
var autoNumberConvertString = true


var chinaLoc = time.FixedZone("CST", 8*3600)
const secondTimeLayout = "2006-01-02 15:04:05"
type ChinaTime struct {
	time.Time
}
func NewChinaTime(time time.Time) ChinaTime {
	return ChinaTime{Time: time.In(chinaLoc)}
}
func (t ChinaTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.In(chinaLoc).Format(secondTimeLayout) + `"`), nil
}
func (t *ChinaTime) UnmarshalJSON(b []byte) error {
	v, err := time.ParseInLocation(`"` + secondTimeLayout + `"`, string(b), chinaLoc)
	t.Time = v
	return err
}