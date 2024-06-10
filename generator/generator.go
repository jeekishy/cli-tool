package generator

func GenerateEvenNumbers(limit int) <-chan int {
	// have a channel to which we can push even number generated
	eventNumChan := make(chan int)

	go func() {
		// last event number to generate should be provided limit * 2
		lastEvenNum := limit * 2

		// generate event numbers to the limit
		for i := 2; i <= lastEvenNum; i += 2 {
			eventNumChan <- i
		}

		// close channel when done
		close(eventNumChan)
	}()

	return eventNumChan
}
