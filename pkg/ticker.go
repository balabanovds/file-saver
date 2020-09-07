package ticker

import "time"

func Every(d time.Duration, work func(time.Time) error) chan bool {
	ticker := time.NewTicker(d)
	done := make(chan bool, 1)

	go func() {
		for {
			select {
			case time := <-ticker.C:
				if err := work(time); err != nil {
					done <- true
				}
			case <-done:
				return
			}
		}
	}()

	return done
}
