package dsv

import (
	"errors"
	"fmt"
	"reflect"
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

// getFields returns all fields of a struct
func getFields(st interface{}) ([]field, error) {
	fields := make([]field, 0)

	values := reflect.ValueOf(st).Elem()

	for index := 0; index < values.NumField(); index++ {
		valueField := values.Field(index)
		typeField := values.Type().Field(index)

		currentField := field{
			value: fmt.Sprintf("%v", valueField.Interface()),
		}

		// Get the index
		indexFromTag, err := getIntegerTag(typeField.Tag, "index", index)
		if err != nil {
			return make([]field, 0), err
		}
		currentField.index = indexFromTag

		// Get length
		lengthFromTag, err := getIntegerTag(typeField.Tag, "length", -1)
		if err != nil {
			return make([]field, 0), err
		}
		currentField.length = lengthFromTag

		// Get the padding character
		paddingCharFromTag, err := getStringTag(typeField.Tag, "paddingChar")
		if err != nil {
			return make([]field, 0), err
		}
		currentField.paddingChar = paddingCharFromTag

		// Get the padding right tag
		paddingRightFromTag, err := getBooleanTag(typeField.Tag, "paddingRight", false)
		if err != nil {
			return make([]field, 0), err
		}
		currentField.paddingRight = paddingRightFromTag

		fields = append(fields, currentField)
	}

	return fields, nil
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
