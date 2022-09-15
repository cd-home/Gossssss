package time_test

import (
	"testing"
	"time"
)

func TestTimeStamp(t *testing.T) {
	ti, _ := time.Parse("2006-01-02 15:04:05", "2021-08-31 23:59:43")
	t.Log(ti.Unix())
}

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

func TestTimeSub(t *testing.T) {
	t1 := time.Now()
	time.Sleep(2 * time.Second)
	t2 := time.Now()
	t.Log(t2.Sub(t1))
}

func TestTimeAfter(t *testing.T) {
	for {
		select {
		case now := <-time.After(2 * time.Second):
			t.Log(now)
		}
	}
}

func TestTimeTick(t *testing.T) {
	tick := time.Tick(2 * time.Second)
	for next := range tick {
		t.Log(next)
	}
}

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
