package time_test

import (
	"testing"
	"time"
)

func TestTimeStamp(t *testing.T) {
	ti, _ := time.Parse("2006-01-02 15:04:05", "2021-08-31 23:59:43")
	t.Log(ti.Unix())
}
