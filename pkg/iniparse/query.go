package iniparse

import (
	"errors"
	"strings"
)

type Section map[string]string

var (
	ErrSectionNotFound = errors.New("section not found")
	ErrItemNotFound    = errors.New("item not found")
	ErrInvalidQuery    = errors.New("invalid query")
)

func QuerySection(result INIResult, query string) (Section, error) {
	value, ok := result[query]
	if ok {
		return value, nil
	}
	return nil, ErrSectionNotFound
}

func QueryItem(result INIResult, query string) (string, error) {
	queryExpression := strings.SplitN(strings.ReplaceAll(query, " ", ""), ".", 2)
	if len(queryExpression) == 2 {
		section := queryExpression[0]
		item := queryExpression[1]
		value, ok := result[section][item]
		if !ok {
			return "", ErrItemNotFound
		}
		return value, nil
	}
	return "", ErrInvalidQuery
}
