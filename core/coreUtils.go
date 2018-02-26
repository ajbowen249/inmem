package inmem

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

func (im *Inmem) deleteKey(key string, kt keyType) {
	switch kt {
	case tString:
		delete(im.strings, key)
	case tHash:
		delete(im.hashes, key)
	case tList:
		delete(im.lists, key)
	case tSet:
		delete(im.sets, key)
	}
}

func (im *Inmem) setType(key string, newType keyType) {
	if oldType, exists := im.keys[key]; exists && oldType != newType {
		im.deleteKey(key, oldType)
	}

	im.keys[key] = newType
}

func (im *Inmem)validateType(key string, requiredType keyType, writer io.Writer) (bool, error) {
	kt, exists := im.keys[key]
	if !exists {
		return false, nil
	}

	if kt != requiredType {
		writer.Write([]byte(errWrongType))
		return true, errors.New(errWrongType)
	}

	return true, nil
}

func validate(command string, argCount int, writer io.Writer) error {
	if requiredCount, ok := argCounts[command]; ok {
		if requiredCount != argCount {
			argumentError := fmt.Sprintf(errArgumentCount, command)
			writer.Write([]byte(argumentError))
			return errors.New(argumentError)
		}

		return nil
	}

	return unknownError(command, writer)
}

func unknownError(command string, writer io.Writer) error {
	unknownError := fmt.Sprintf(errUnknownCommand, command)
	writer.Write([]byte(unknownError))
	return errors.New(unknownError)
}

func trimQuotes(value string) string {
	return strings.TrimSuffix(strings.TrimPrefix(value, "\""), "\"")
}
