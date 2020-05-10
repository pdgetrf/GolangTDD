package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}

	t.Run("key exists", func(t *testing.T) {
		expect := "this is a test"
		actual, err := dictionary.search("test")
		assertNil(t, err)
		assertEqual(t, expect, actual)
	})

	t.Run("key does not exist", func(t *testing.T) {
		expect := ""
		actual, err := dictionary.search("abc")
		assertNotNil(t, err)
		assertEqual(t, ErrNotFound, err)
		assertEqual(t, expect, actual)
	})
}

func TestAdd(t *testing.T) {
	t.Run("add without dup", func(t *testing.T) {
		dictionary := Dictionary{}
		expect := "this is a test"
		err := dictionary.add("test", expect)
		assertNil(t, err)

		actual, err := dictionary.search("test")
		assertNil(t, err)
		assertEqual(t, expect, actual)
	})
	t.Run("add with dup", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.add("test", "this is a test")

		err := dictionary.add("test", "this is another test")
		assertNotNil(t, err)
		assertEqual(t, err, ErrDuplicatedWord)

		actual, err := dictionary.search("test")
		assertNil(t, err)
		expect := "this is a test"
		assertEqual(t, expect, actual)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update without dup", func(t *testing.T) {
		dictionary := Dictionary{}
		expect := "this is a test"
		err := dictionary.update("test", expect)
		assertNotNil(t, err)
		assertEqual(t, ErrNotFound, err)
	})
	t.Run("update with dup", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.add("test", "this is a test")

		expect := "this is another test"
		err := dictionary.update("test", expect)
		assertNil(t, err)

		actual, err := dictionary.search("test")
		assertNil(t, err)
		assertEqual(t, expect, actual)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete an existing item", func(t *testing.T) {
		dictionary := Dictionary{}
		expect := "this is a test"
		dictionary.add("test", expect)

		dictionary.delete("test")

		_, err := dictionary.search("test")
		assertEqual(t, err, ErrNotFound)
	})

	t.Run("delete a non-existing item", func(t *testing.T) {
		dictionary := Dictionary{}

		dictionary.delete("test")

		_, err := dictionary.search("test")
		assertEqual(t, err, ErrNotFound)
	})
}

func assertEqual(t *testing.T, expect, actual interface{}) {
	t.Helper()
	switch expect.(type) {
	case error, string, nil:
		if expect != actual {
			t.Errorf("actual %q expect %q", actual, expect)
		}
	default:
		t.Errorf("unknown type for comparison")
	}
}

func assertNil(t *testing.T, actual interface{}) {
	t.Helper()
	assertEqual(t, nil, actual)
}

func assertNotNil(t *testing.T, actual interface{}) {
	t.Helper()
	if nil == actual {
		t.Errorf("actual %q expect nil", actual)
	}
}
