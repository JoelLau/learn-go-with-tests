package main

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	val, ok := d[key]
	if !ok {
		return "", ErrSearchKeyNotFound
	}

	return val, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrSearchKeyNotFound:
		d[key] = value
	case nil:
		return ErrAddKeyCollision
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, newValue string) error {
	_, err := d.Search(key)
	switch err {
	case nil:
		d[key] = newValue
	case ErrSearchKeyNotFound:
		return ErrUpdateKeyNotFound
	default:
		return err
	}

	return nil
}
