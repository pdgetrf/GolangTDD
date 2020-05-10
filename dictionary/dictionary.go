package dictionary

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrDuplicatedWord   = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) add(word, definition string) error {
	_, err := d.search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrDuplicatedWord
	default:
		return err
	}
	return nil
}

func (d Dictionary) update(word, definition string) error {
	_, err := d.search(word)

	switch err {
	case ErrWordDoesNotExist:
		return err
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) delete(word string) {
	delete(d, word)
}
