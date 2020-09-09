package app

import (
	"context"
	"log"
	"time"
)

func every(ctx context.Context, d time.Duration, work func(time.Time) error) {
	ticker := time.NewTicker(d)

	go func() {
		for {
			select {
			case t := <-ticker.C:
				if err := work(t); err != nil {
					log.Printf("error in job (skip it): %v\n", err)
				}
			case <-ctx.Done():
				return
			}
		}
	}()
}
