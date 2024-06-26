package dsv

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const (
	tagName      = "dsv"
	tagSeparator = ","

	// Default values in the tags
	tagDefaultPaddingChar  = " "
	tagDefaultPaddingRight = false
)

// field represents a field with data in a struct
type field struct {
	fieldName    string
	index        int
	length       int
	value        string
	paddingChar  string
	paddingRight bool
}

// getFields returns all fields of a struct
func getFields(st interface{}) ([]field, error) {
	var fields []field

	values := reflect.ValueOf(st).Elem()

	for index := 0; index < values.NumField(); index++ {
		valueField := values.Field(index)
		typeField := values.Type().Field(index)

		currentField := field{
			fieldName: typeField.Name,
			value:     fmt.Sprintf("%v", valueField.Interface()),
		}

		// Get dsv tag
		tagData := typeField.Tag.Get(tagName)
		if len(tagData) == 0 {
			// Field will not be included
			continue
		}

		splittedTag := strings.Split(tagData, tagSeparator)

		// Get the index
		indexFromTag, err := getIndexFromTag(splittedTag)
		if err != nil {
			return []field{}, err
		}
		currentField.index = indexFromTag

		// Get length
		lengthFromTag, err := getLengthFromTag(splittedTag)
		if err != nil {
			return []field{}, err
		}
		currentField.length = lengthFromTag

		// Get the padding character
		paddingCharFromTag, err := getPaddingCharFromTag(splittedTag, tagDefaultPaddingChar)
		if err != nil {
			return []field{}, err
		}
		currentField.paddingChar = paddingCharFromTag

		// Get the padding right tag
		paddingRightFromTag, err := getPaddingRightFromTag(splittedTag, tagDefaultPaddingRight)
		if err != nil {
			return []field{}, err
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

// setFieldsToStruct will set the values in the interface with the fields data
func setFieldsToStruct(fields []field, i interface{}) (interface{}, error) {
	values := reflect.ValueOf(i).Elem()

	fieldsLength := len(fields)

	for index := 0; index < fieldsLength; index++ {
		fieldValue := fields[index].value
		fieldName := fields[index].fieldName

		element := values.FieldByName(fieldName)
		if !element.CanSet() {
			continue
		}
		switch element.Kind() {
		case reflect.String:
			element.SetString(fieldValue)
			break
		case reflect.Int:
			integer, err := strconv.ParseInt(fieldValue, 10, 0)
			if err != nil {
				return nil, err
			}
			element.SetInt(integer)
			break
		case reflect.Float32, reflect.Float64:
			float, err := strconv.ParseFloat(fieldValue, 0)
			if err != nil {
				return nil, err
			}
			element.SetFloat(float)
			break
		case reflect.Bool:
			boolean, err := strconv.ParseBool(fieldValue)
			if err != nil {
				return nil, err
			}
			element.SetBool(boolean)
			break
		}
	}

	return i, nil
}
