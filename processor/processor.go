package processor

import (
	"fmt"

	"github.com/jeekishy/cli-tool/repository"
)

func ProcessResponse(message *repository.Message) {
	if message == nil {
		fmt.Print("empty message\n")
		return
	}

	fmt.Printf("%s - %t\n", message.Title, message.Completed)
}
