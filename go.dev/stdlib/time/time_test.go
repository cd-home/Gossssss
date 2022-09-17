package time_test

import (
	"fmt"
	"testing"
	"time"
)

const TimeLayOut = "2006-01-02 15:04:05"

// TestTimeNow
func TestTimeNow(t *testing.T) {
	t.Log(time.Now())
}

// TestTimeDate
func TestTimeDate(t *testing.T) {
	date := time.Date(2022, time.September, 10, 10, 30, 0, 0, time.Local)
	t.Log(date)
}

// TestTimeFormat
func TestTimeFormat(t *testing.T) {
	str := time.Now().Format("20060102150405")
	t.Log(str)
}

// TestParseStringTime
// 解析字符串时间为Time对象
func TestParseStringTime(t *testing.T) {
	timeStr := "2022-01-02 10:03:21"

	// Default UTC
	tm, _ := time.Parse(TimeLayOut, timeStr)
	fmt.Println(tm)

	// 正确指定时区解析
	// 获取时区
	loc, _ := time.LoadLocation("Asia/Shanghai")
	tm2, _ := time.ParseInLocation(TimeLayOut, timeStr, loc)
	fmt.Println(tm2)
	fmt.Println(tm2.In(loc))
}

// TestTimeParseDuration
// 解析字符串形式 "Duration" 为 Duration
func TestTimeParseDuration(t *testing.T) {
	duration, err := time.ParseDuration("1h30m10s100")
	if err != nil {
		t.Error(err)
	}
	hours := duration.Hours()
	t.Log(hours)

	after := time.Now().Add(duration)
	t.Log(after)
}

// TestTimeAfter
// 注意内存泄漏问题, 推荐使用 NewTicker
func TestTimeAfter(t *testing.T) {
	select {
	case now := <-time.After(2 * time.Second):
		t.Log(now)
	}
}

// TestAfterFunc
func TestAfterFunc(t *testing.T) {
	timer := time.AfterFunc(time.Second, func() {
		t.Log("AfterFunc")
	})
	defer timer.Stop()
	time.Sleep(time.Second * 2)
}

// TestTimeTick
// 对 ticker 的包装  要注意内存泄露
func TestTimeTick(t *testing.T) {
	count := 0
	tick := time.Tick(time.Second)
	for next := range tick {
		t.Log(next)
		count++
		if count == 5 {
			break
		}
	}
}

// TestTimeNewTicker
func TestTimeNewTicker(t *testing.T) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan struct{})
	go func() {
		time.Sleep(5 * time.Second)
		done <- struct{}{}
	}()
	for {
		select {
		case <-done:
			return
		case now := <-ticker.C:
			t.Log(now)
		}
	}
}
