package generator

func GenerateEvenNumbers(limit int) <-chan int {
	// have a channel to which we can push even number generated
	evenNumChan := make(chan int)

	go func() {
		// last even number to generate should be the provided limit * 2
		lastEvenNum := limit * 2

		// generate even numbers up to the limit
		for i := 2; i <= lastEvenNum; i += 2 {
			evenNumChan <- i
		}

		// close channel when done
		close(evenNumChan)
	}()

	return evenNumChan
}
