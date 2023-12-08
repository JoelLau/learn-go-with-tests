package main

type Dictionary map[string]string

func (dictionary Dictionary) Search(term string) (string, error) {
	result, ok := dictionary[term]
	if !ok {
		return "", NotFoundError
	}

	return result, nil
}

func (dictionary Dictionary) Add(term, definition string) (string, error) {
	existing, ok := dictionary[term]
	if ok {
		return existing, TermExistsError
	}

	dictionary[term] = definition
	return definition, nil
}
