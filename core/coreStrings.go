package inmem

import (
	"strings"
	"fmt"
	"io"
)

// Set sets a key to the given value.
func (im *Inmem) Set(key string, value string, writer io.Writer) error {
	im.setType(key, tString)
	im.strings[key] = trimQuotes(value)
	writer.Write([]byte(resOK))
	return nil
}

// Get returns a string from the given key.
func (im *Inmem) Get(key string, writer io.Writer) error {
	if exists, err := im.validateType(key, tString, writer); true {
		if !exists {
			writer.Write([]byte(resNil))
			return nil
		}

		if err != nil {
			return err
		}
	}

	value, _ := im.strings[key]
	writer.Write([]byte(fmt.Sprintf("\"%s\"", value)))
	return nil
}

// Append adds more content to a string.
func (im *Inmem) Append(key string, addendum string, writer io.Writer) error {
	if _, err := im.validateType(key, tString, writer); err != nil {
		return err
	}

	var builder strings.Builder
	if oldValue, exists := im.strings[key]; exists {
		builder.WriteString(oldValue)
	}

	builder.WriteString(addendum)

	im.strings[key] = builder.String()
	im.setType(key, tString)

	writer.Write([]byte(fmt.Sprintf(resInt, builder.Len())))
	return nil
}
