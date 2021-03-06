/*
Package dictionary creates a module for working with words and their definitions.
*/
package dictionary

// Dictionary contains words with definitions
type Dictionary map[string]string

const (
	// ErrWordNotFound indicates we tried to find a word but it's not yet in the dictionary
	ErrWordNotFound = Err("could not find the word you were looking for")

	// ErrWordExists indicates we tried to add a word that already exists
	ErrWordExists = Err("word already exists")
)

// An Err indicates an error on one of the dictionary operations
type Err string

func (e Err) Error() string {
	return string(e)
}

// Add sets a new word/definition in the dictionary, if it doesn't already exist
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

// Search is used to read words from the dictionary
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrWordNotFound
	}
	return definition, nil
}

// Update changes the definition of a word
func (d Dictionary) Update(word, definition string) {
	d[word] = definition
}
