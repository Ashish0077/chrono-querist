package workers

import "log"

type BackgroundJobFunc func()

func Go(fn BackgroundJobFunc) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recovered from panic in background job: %v", r)
			}
		}()

		// Executing the background job
		fn()
	}()
}
