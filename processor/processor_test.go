package processor

import "github.com/jeekishy/cli-tool/repository"

// using GO's "Example" test package to validate expected CLI output
// @url https://go.dev/blog/examples
// @url https://pkg.go.dev/testing#hdr-Examples
func ExampleProcessResponse() {
	testCases := []*repository.Message{
		{
			UserID:    1,
			ID:        2,
			Title:     "Mock title",
			Completed: false,
		},
		nil,
	}

	for _, tc := range testCases {
		ProcessResponse(tc)
	}
	// Output: Mock title - false
	// empty message
}
