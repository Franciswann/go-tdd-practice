package main

type Dictionary map[string]string

func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}

// DictionaryErr is a custom error type for dictionary operations.
// It implements the error interface by providing an Error() method.
const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot perform operation on word because it does not exist")
)

// Dictionary represents errors that can occur during dictionary operations.
type DictionaryErr string

// Error returns the string representation of the error.
func (e DictionaryErr) Error() string {
	return string(e)
}

// Search returns the definition of the given word
// It returns an error if the word does not exist in dictionary
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add adds a new word with its definition to the dictionary
// If the word already exists, it returns ErrWordExists and does not
// modify the existing definition.
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

// Update modifies the definition of an existing word in the dictionary.
// Returns an error if the word does not exist.
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
