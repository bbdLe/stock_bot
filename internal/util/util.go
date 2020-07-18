package util

import (
	"fmt"
	"os"
	"strconv"
	"syscall"
	"time"
)

const (
	DateFormat  = "2006-01-02"
	DateTimeFormat  = "2006-01-02 15:04:05"
)

func ConvVal2MarkDown(str string) string {
	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return ""
	}

	if val > 0 {
		return fmt.Sprintf(`<font color="#FF0000">%.2f</font>`, val)
	} else {
		return fmt.Sprintf(`<font color="info">%.2f</font>`, val)
	}
}

func ConvPercent2MarkDown(str string) string {
	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return ""
	}

	if val > 0 {
		return fmt.Sprintf(`<font color="#FF0000">%.2f%%</font>`, val)
	} else {
		return fmt.Sprintf(`<font color="info">%.2f%%</font>`, val)
	}
}

// 当天零时
func ZeroTime(t time.Time) time.Time {
	timeStr := t.Format(DateFormat)
	tt, _ := time.ParseInLocation(DateFormat, timeStr, time.Local)
	return tt
}

func IsMarkTime() bool {
	now := time.Now().Local()
	zeroTime := ZeroTime(now)
	openTime := zeroTime.Add(time.Hour * 9 + time.Minute * 30 - time.Second)
	closeTime := zeroTime.Add(time.Hour * 15 + time.Second)

	if now.After(openTime) && now.Before(closeTime) {
		return true
	} else {
		return false
	}
}

func WritePid(filePath string) error {
	file, err := os.OpenFile(filePath, os.O_CREATE | os.O_TRUNC | os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(fmt.Sprintf("%d", syscall.Getpid())); err != nil {
		return err
	}

	return nil
}