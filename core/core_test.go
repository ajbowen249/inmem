package inmem

// These tests are general inmem-related tests

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCommandValidation(t *testing.T) {
	cases := []string{
		"got",
		"get it",
		"fleeb",
	}

	inmem := NewInmem()

	for _, command := range cases {
		outBuffer := new(bytes.Buffer)
		err := inmem.Execute(command, []string{}, outBuffer)
		if err == nil {
			t.Error("Expected non-nil error")
		}

		expectedResponse := fmt.Sprintf(errUnknownCommand, command)
		response := outBuffer.String()

		if response != expectedResponse {
			t.Errorf("Expected \"%s\", got \"%s\"", expectedResponse, response)
		}
	}
}
