package inmem

// These tests are specific to commands that deal with the string type

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
)

func TestGetUnsetString(t *testing.T) {
	inmem := NewInmem()

	getString(inmem, "testKey", resNil, t)
}

func TestSetString(t *testing.T) {
	cases := []struct{ key, value, expectedReturn string }{
		{"name", "Fred", "\"Fred\""},
		{"name", "\"Fred\"", "\"Fred\""},
		{"age", "35", "\"35\""},
		{"test", "\"\"", "\"\""},
	}

	inmem := NewInmem()

	for _, c := range cases {
		setString(inmem, c.key, c.value, t)
		getString(inmem, c.key, c.expectedReturn, t)
	}
}

func TestMultipleGetAndSetString(t *testing.T) {
	inmem := NewInmem()

	testKey := "name"
	testValue1 := "Fred"
	testValue2 := "Steve"

	setString(inmem, testKey, testValue1, t)
	getString(inmem, testKey, fmt.Sprintf("\"%s\"", testValue1), t)

	setString(inmem, testKey, testValue2, t)
	getString(inmem, testKey, fmt.Sprintf("\"%s\"", testValue2), t)
}

func TestArgValidation(t *testing.T) {
	cases := []struct {
		command string
		args    []string
	}{
		{"set", []string{}},
		{"set", []string{"justAKey"}},
		{"set", []string{"lots", "of", "args"}},
		{"get", []string{}},
		{"get", []string{"lots", "of", "args"}},
	}

	inmem := NewInmem()

	for _, c := range cases {
		outBuffer := new(bytes.Buffer)
		err := inmem.Execute(c.command, c.args, outBuffer)
		if err == nil {
			t.Error("Expected non-nil error")
		}

		expectedResponse := fmt.Sprintf(errArgumentCount, c.command)
		checkResponse(outBuffer.String(), expectedResponse, t)
	}
}

func TestAppend(t *testing.T) {
	cases := []struct {
		hasInitialValue  bool
		initialvalue, addendum, expectedResponse string
	}{
		{ true, "hello", " dere", "(integer) 10" },
		{ false, "", "whole new value", "(integer) 15" },
		{ true, "a", "b", "(integer) 2" },
		{ true, "a", "", "(integer) 1" },
		{ true, "", "a", "(integer) 1" },
	}

	testKey := "testKey"

	for _, c := range cases {
		inmem := NewInmem()
		if c.hasInitialValue {
			setString(inmem, testKey, c.initialvalue, t)
		}

		appendString(inmem, testKey, c.addendum, c.expectedResponse, t)
	}
}

func TestIncr(t *testing.T) {
	nonIntError := errors.New(errNotIntOrOutOfRange)

	cases := []struct {
		hasInitialValue bool
		initialvalue, expectedResponse string
		expectedError error
	}{
		{ true, "1", "(integer) 2", nil },
		{ true, "10", "(integer) 11", nil },
		{ true, "0", "(integer) 1", nil },
		{ false, "", "(integer) 1", nil },
		{ true, "not a number", errNotIntOrOutOfRange, nonIntError },
	}

	testKey := "testKey"

	for _, c := range cases {
		inmem := NewInmem()
		if c.hasInitialValue {
			setString(inmem, testKey, c.initialvalue, t)
		}

		outBuffer := new(bytes.Buffer)
		err := inmem.Execute("incr", []string{ testKey }, outBuffer);

		checkResponse(outBuffer.String(), c.expectedResponse, t)
		if c.expectedError != nil && err.Error() != c.expectedError.Error() {
			t.Errorf("Expected error %e, got %e", c.expectedError, err)
		}
	}
}

func setString(inmem *Inmem, key string, value string, t *testing.T) {
	outBuffer := new(bytes.Buffer)
	if err := inmem.Execute("set", []string{ key, value }, outBuffer); err != nil {
		t.Errorf("Unexpected error %e", err)
	}

	checkResponse(outBuffer.String(), resOK, t)
}

func getString(inmem *Inmem, key string, expectedResponse string, t *testing.T) {
	outBuffer := new(bytes.Buffer)
	if err := inmem.Execute("get", []string{ key }, outBuffer); err != nil {
		t.Errorf("Unexpected error %e", err)
	}

	checkResponse(outBuffer.String(), expectedResponse, t)
}

func appendString(inmem *Inmem, key string, addendum string, expectedResponse string, t *testing.T) {
	outBuffer := new(bytes.Buffer)
	if err := inmem.Execute("append", []string{ key, addendum }, outBuffer); err != nil {
		t.Errorf("Unexpected error %e", err)
	}

	checkResponse(outBuffer.String(), expectedResponse, t)
}

func checkResponse(actualResponse string, expectedResponse string, t *testing.T) {
	if actualResponse != expectedResponse {
		t.Errorf("Expected \"%s\", got \"%s\"", expectedResponse, actualResponse)
	}
}
