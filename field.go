package dsv

import (
	"errors"
	"strings"
)

// field represents a field with data in a struct
type field struct {
	index        int
	length       int
	value        string
	paddingChar  string
	paddingRight bool
}

// softFields will sort the fields by index
func sortFields(fields []field) ([]field, error) {
	fieldsLen := len(fields)
	sortedFields := make([]field, fieldsLen)
	copy(sortedFields, fields)
	ordered := false
	for !ordered {
		for index := 0; index < fieldsLen; index++ {
			nextIndex := index + 2
			if nextIndex > fieldsLen {
				break
			}
			valuesToCompare := sortedFields[index:nextIndex]
			if valuesToCompare[0].index > valuesToCompare[1].index {
				tempValue := sortedFields[index]
				sortedFields[index] = sortedFields[index+1]
				sortedFields[index+1] = tempValue
			}
		}
		fieldsLen--
		if fieldsLen == 0 {
			ordered = true
		}
	}
	return sortedFields, nil
}

// padValue will return the value of the field with the right padding using paddingChar and length
func padValue(f field) (string, error) {
	if len(f.paddingChar) > 1 {
		return f.value, errors.New("The padding char should be one char length")
	}
	if len(f.value) >= f.length {
		return f.value, nil
	}
	lengthDiff := f.length - len(f.value)

	newValue := strings.Repeat(f.paddingChar, lengthDiff)
	if f.paddingRight {
		newValue = f.value + newValue
	} else {
		newValue += f.value
	}

	return newValue, nil
}
