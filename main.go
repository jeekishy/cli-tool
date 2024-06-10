package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/jeekishy/cli-tool/generator"
	"github.com/jeekishy/cli-tool/processor"
	"github.com/jeekishy/cli-tool/repository"
)

func init() {
	// validate we have basic configurations available before starting
	if len(os.Args) < 3 {
		fmt.Println("failed to start application")
		os.Exit(0)
	}
}

func main() {
	// get base URL path and even number limit from cmd arguments
	basePath := os.Args[1]
	limit, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("failed to get application even number limit error, %s\n", err)
		return
	}

	// set some need dependencies
	httpClient := &http.Client{}
	r := repository.New(httpClient, basePath)
	wg := new(sync.WaitGroup)

	// execute main application
	for n := range generator.GenerateEvenNumbers(limit) {
		wg.Add(1)

		// ideally we should limit the number of goroutines we spin up as this can cause OOM kills
		// for this exercise I have left this out, but would use a buffered channel as a control mechanism
		go func() {
			// get data from endpoint
			message, err := r.GetMessage(n)
			if err != nil {
				fmt.Printf("failed to get data from endpoint, error %s\n", err.Error())
				return
			}

			// process message we received from endpoint
			processor.ProcessResponse(message)
			wg.Done()
		}()
	}

	wg.Wait()
}
