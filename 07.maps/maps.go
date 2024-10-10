package maps

import (
	"errors"
)

type Dictionary map[string]string

var (
	ErrNotFound    = errors.New("key not found")
	ErrExistingKey = errors.New("key already exists")
)

func (dict Dictionary) Search(key string) (string, error) {
	value, ok := dict[key]
	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}

func (dict Dictionary) Add(key, value string) error {
	_, err := dict.Search(key)
	switch err {
	case ErrNotFound:
		dict[key] = value
	case nil:
		return ErrExistingKey
	default:
		return err
	}
	return nil
}

func (dict Dictionary) Update(key, value string) error {
	_, err := dict.Search(key)

	switch err {
	case ErrNotFound:
		return ErrNotFound
	case nil:
		dict[key] = value
	default:
		return err
	}
	return nil
}

func (dict Dictionary) Delete(key string) error {
	_, err := dict.Search(key)
	if err != nil {
		return err
	}
	delete(dict, key)
	return nil
}
