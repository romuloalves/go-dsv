package dsv

import (
	"reflect"
	"strconv"
)

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
