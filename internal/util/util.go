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

func ConvPrice2MarkDown(kpj string, xj string) string {
	kpjVal, err := strconv.ParseFloat(kpj, 64)
	if err != nil {
		return ""
	}

	xjVal, err := strconv.ParseFloat(xj, 64)
	if err != nil {
		return ""
	}

	if xjVal > kpjVal {
		return fmt.Sprintf(`<font color="#FF0000">%.2f</font>`, xjVal)
	} else {
		return fmt.Sprintf(`<font color="info">%.2f</font>`, xjVal)
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
	MorningOpenTime := zeroTime.Add(time.Hour * 9 + time.Minute * 30 + time.Minute)
	MorningCloseTime := MorningOpenTime.Add(time.Hour * 2 + time.Second)

	afternoonOpenTime := zeroTime.Add(time.Hour * 13 + time.Minute)
	afterNoonCloseTime := afternoonOpenTime.Add(time.Hour * 2 + time.Second)

	if (now.After(MorningOpenTime) && now.Before(MorningCloseTime)) || (now.After(afternoonOpenTime) && now.Before(afterNoonCloseTime)) {
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