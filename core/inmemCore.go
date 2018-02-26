// Package inmem provides a strongly Redis-inspired and partially Redis-
// compatible library for in-memory data storage.
package inmem

import (
	"io"
	"strings"
)

// Inmem is the root inmem data structure and single access point for all
// data in the system. Its primary interface is the Execute method. Most
// other methods are imlementations of the set of Redis-compatible commands
// and Redis-inspired extensions.
type Inmem struct {
	strings map[string]string
	hashes  map[string]map[string]string
	lists   map[string][]string
	sets    map[string]StringSet

	keys map[string]keyType
}

// NewInmem creates a new Inmem.
func NewInmem() *Inmem {
	return &Inmem{
		make(map[string]string),
		make(map[string]map[string]string),
		make(map[string][]string),
		make(map[string]StringSet),
		make(map[string]keyType),
	}
}

// Execute executes the given command with the given arguments and outputs
// data via the provided writer.
func (im *Inmem) Execute(rawCommand string, args []string, writer io.Writer) error {
	command := strings.ToLower(rawCommand)
	if err := validate(command, len(args), writer); err != nil {
		return err
	}

	switch command {
	case cmdSet:
		return im.Set(args[0], args[1], writer)
	case cmdGet:
		return im.Get(args[0], writer)
	case cmdAppend:
		return im.Append(args[0], args[1], writer)
	default:
		return unknownError(command, writer)
	}
}
