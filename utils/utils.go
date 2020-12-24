package utils

import (
	"strconv"
	"time"
)

// NowMillis 当前时间 毫秒
func NowMillis() int64 {
	nanos := time.Now().UnixNano()
	millis := nanos / int64(time.Millisecond)
	return millis
}

// NowMillisStr 当前时间字符串 毫秒
func NowMillisStr() string {
	millis := NowMillis()

	return strconv.FormatInt(millis, 10)
}
