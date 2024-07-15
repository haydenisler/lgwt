package gmaps

import "errors"

type Dictionary map[string]string

var (
	ErrWordNotFound      = errors.New("could not find the word you were looking for")
	ErrWordAlreadyExists = errors.New("could not add word because it already exists")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrWordNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, ok := d[word]

	if ok {
		return ErrWordAlreadyExists
	}

	d[word] = definition
	return nil
}
