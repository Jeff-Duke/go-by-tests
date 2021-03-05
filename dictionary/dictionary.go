package dictionary

import "errors"

type Dictionary map[string]string

var (
	ErrWordNotFound = errors.New("could not find the word you were looking for")
	ErrWordExists   = errors.New("word already exists")
)

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrWordNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrWordNotFound
	}
	return definition, nil
}
