package time

import "time"

func PastNDay(n int) []string {
	now := time.Now()
	days := make([]string, n)
	for i := n; i > 0; i-- {
		days[n-i] = now.AddDate(0, 0, -i).Format("2006-01-02")
	}
	return days
}
