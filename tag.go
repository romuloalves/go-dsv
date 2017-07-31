package dsv

import (
	"errors"
	"strconv"
	"strings"
)

const (
	// Indexes of each information inside the dsv tag
	indexTagPosition        = 0
	lengthTagPosition       = 1
	paddingCharTagPosition  = 2
	paddingRightTagPosition = 3

	paddingCharMaxLength = 1
)

func getIndexFromTag(tag []string) (int, error) {
	if len(tag) <= indexTagPosition {
		return -1, errors.New("Index not found in tag")
	}

	index := tag[indexTagPosition]

	data, err := strconv.Atoi(index)
	if err != nil {
		return -1, err
	}

	return data, nil
}

func getLengthFromTag(tag []string) (int, error) {
	if len(tag) <= lengthTagPosition {
		return -1, nil
	}

	length := tag[lengthTagPosition]

	data, err := strconv.Atoi(length)
	if err != nil {
		return -1, err
	}

	return data, nil
}

func getPaddingCharFromTag(tag []string, defaultValue string) (string, error) {
	if len(tag) <= paddingCharTagPosition {
		return defaultValue, nil
	}

	paddingChar := tag[paddingCharTagPosition]
	paddingCharLength := len(paddingChar)

	if paddingCharLength == 0 || paddingCharLength > paddingCharMaxLength {
		return "", errors.New("The padding char must contain one character")
	}

	if strings.Contains(paddingChar, "-") {
		return defaultValue, nil
	}

	return paddingChar, nil
}

func getPaddingRightFromTag(tag []string, defaultValue bool) (bool, error) {
	if len(tag) <= paddingRightTagPosition {
		return defaultValue, nil
	}

	paddingRight := tag[paddingRightTagPosition]

	if strings.Contains(paddingRight, "-") {
		return defaultValue, nil
	}

	data, err := strconv.ParseBool(paddingRight)
	if err != nil {
		return defaultValue, err
	}

	return data, nil
}
