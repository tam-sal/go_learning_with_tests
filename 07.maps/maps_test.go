package maps

import (
	"fmt"
	"testing"
)

var dict1, dict2, dict3 Dictionary

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func init() {
	dict1 = Dictionary{"test1": "value of test 1"}
	dict2 = Dictionary{"test2": "value of test 2"}
	dict3 = make(map[string]string)
	ErrNotFound = DictionaryErr("key not found")
	ErrExistingKey = DictionaryErr("key already exists")
}

func TestAdd(t *testing.T) {
	t.Run("new key", func(t *testing.T) {
		key := "testOne"
		value := "value of testOne"
		dict1.Add(key, value)
		assertValues(t, dict1, key, value)
	})

	t.Run("existing key", func(t *testing.T) {
		key := "test1"
		value := "overwritting test1 value"
		err := dict1.Add(key, value)
		assertError(t, err, ErrExistingKey)
		originalValue := "value of test 1"
		assertValues(t, dict1, key, originalValue)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete not found key", func(t *testing.T) {
		key := "random"
		err := dict1.Delete(key)
		assertError(t, err, ErrNotFound)
	})

	t.Run("delete existing key", func(t *testing.T) {
		key := "test3"
		value := "value of test 3"
		dict3.Add(key, value)
		assertValues(t, dict3, key, value)
	})
}

func TestUpdate(t *testing.T) {
	value := "this is an updated value"
	t.Run("Update unfound", func(t *testing.T) {
		key := "rand"
		err := dict2.Update(key, value)
		assertError(t, err, ErrNotFound)
	})
	t.Run("update existing", func(t *testing.T) {
		key := "test2"
		dict2.Update(key, value)
		assertValues(t, dict2, key, value)
	})
}

func assertValues(t testing.TB, dictionary Dictionary, key, value string) {
	t.Helper()

	got, err := dictionary.Search(key)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, value)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got: %q Want: %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func ExampleDictionary_Add() {
	dict := Dictionary{}
	err := dict.Add("test", "value of test")

	if err != nil {
		fmt.Println("Add Error:", err)
	}

	// Output:
	// Add Error: <nil>
}

func ExampleDictionary_Search() {
	dict := Dictionary{"test": "value of test"}
	value, err := dict.Search("test")

	if err != nil {
		fmt.Println("Search Error:", err)
	} else {
		fmt.Println("Found:", value)
	}

	// Output:
	// Found: value of test
}

func ExampleDictionary_Update() {
	dict := Dictionary{"test": "value of test"}
	err := dict.Update("test", "new value")

	if err != nil {
		fmt.Println("Update Error:", err)
	} else {
		fmt.Println("Updated value:", dict["test"])
	}

	// Output:
	// Updated value: new value
}

func ExampleDictionary_Delete() {
	dict := Dictionary{"test": "value of test"}
	err := dict.Delete("test")

	if err != nil {
		fmt.Println("Delete Error:", err)
	} else {
		_, err := dict.Search("test")
		fmt.Println("Search after delete:", err)
	}

	// Output:
	// Search after delete: key not found
}
