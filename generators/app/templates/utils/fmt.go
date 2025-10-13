package utils

import (
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"
)

// 转换access数据库时间到当前时间
func TimeFormat(day float64) string {
	seconds := day * 86400
	baseTime := time.Date(1899, time.December, 30, 0, 0, 0, 0, time.UTC)
	actualTime := baseTime.Add(time.Second * time.Duration(seconds))
	return actualTime.Format(time.DateTime)
}

func StampStrToFloat(stampStr string) (float64, error) {
	return strconv.ParseFloat(stampStr, 64)
}

func StampFloatToStr(stampFloat float64) string {
	return strconv.FormatFloat(stampFloat, 'f', 10, 64)
}

func DateTimeToStamp(dateStr string) float64 {
	theTime, err := time.Parse(time.DateTime, dateStr)
	if err != nil {
		log.Println("解析时间失败", err.Error())
		return 0
	}

	accessBaseTime := time.Date(1899, time.December, 30, 0, 0, 0, 0, time.UTC)
	durationDiff := theTime.Sub(accessBaseTime)
	dayDiff := int(durationDiff.Hours() / 24)
	secondDiff := durationDiff.Seconds() - float64(dayDiff)*86400.0
	return float64(dayDiff) + float64(secondDiff)/86400.0
}

func GenUuid() string {
	return uuid.NewString()
}

// Access时间转换函数
// Access 基准时间（1899-12-30 00:00:00 UTC）
var accessBaseTime = time.Date(1899, 12, 30, 0, 0, 0, 0, time.Local)

// AccessToTime 将 Access 时间戳（浮点数，单位为“天”）转换为 Go 的 time.Time
func AccessToTime(accessTs float64) time.Time {
	seconds := accessTs * 24 * 60 * 60
	return accessBaseTime.Add(time.Duration(seconds) * time.Second)
}

// TimeToAccess 将 Go 的 time.Time 转换为 Access 时间戳（浮点数，单位为“天”）
func TimeToAccess(t time.Time) float64 {
	duration := t.Sub(accessBaseTime)
	return duration.Hours() / 24
}
