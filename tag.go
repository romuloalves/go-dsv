package dsv

import (
	"reflect"
	"strconv"
)

// getFields returns all fields of a struct
func getFields(st interface{}) ([]field, error) {
	fields := make([]field, 0)

	values := reflect.ValueOf(st).Elem()

	for index := 0; index < values.NumField(); index++ {
		valueField := values.Field(index)
		typeField := values.Type().Field(index)

		currentField := field{
			value: valueField.String(),
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

func getStringTag(tag reflect.StructTag, fieldName string) (string, error) {
	fieldTag := tag.Get(fieldName)
	return fieldTag, nil
}

func getIntegerTag(tag reflect.StructTag, fieldName string, defaultValue int) (int, error) {
	responseInteger := defaultValue

	fieldTag := tag.Get(fieldName)
	if fieldTag == "" {
		return responseInteger, nil
	}

	fieldTagInteger, err := strconv.Atoi(fieldTag)
	if err != nil {
		return responseInteger, err
	}
	responseInteger = fieldTagInteger

	return responseInteger, nil
}

func getBooleanTag(tag reflect.StructTag, fieldName string, defaultValue bool) (bool, error) {
	responseBool := defaultValue

	fieldTag := tag.Get(fieldName)
	if fieldTag == "" {
		return responseBool, nil
	}

	fieldTagBool, err := strconv.ParseBool(fieldTag)
	if err != nil {
		return responseBool, err
	}
	responseBool = fieldTagBool

	return responseBool, nil
}
