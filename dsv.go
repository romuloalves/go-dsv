package dsv

import (
	"errors"
	"strings"
)

// StructToDSV returns a dsv string from the struct
func StructToDSV(i interface{}, separator string) (string, error) {
	// Validating inputs
	if i == nil {
		return "", errors.New("The struct in the parameters can not be null")
	}
	if len(separator) == 0 {
		return "", errors.New("Separator can not be empty")
	}

	// Processing data
	fields, err := getFields(i)
	if err != nil {
		return "", err
	}

	sortedFields, err := sortFields(fields)
	if err != nil {
		return "", err
	}

	var responseLine string
	for index := 0; index < len(sortedFields); index++ {
		finalValue, err := padValue(sortedFields[index])
		if err != nil {
			return "", err
		}

		if index > 0 {
			responseLine += separator
		}
		responseLine += finalValue
	}

	return responseLine, nil
}

// ToStruct returns an interface based in the dsv data, the struct interface and the separator
func ToStruct(dsv string, structInterface interface{}, separator string) (interface{}, error) {
	// Validating inputs
	if len(dsv) == 0 {
		return nil, errors.New("DSV data can not be empty")
	}
	if structInterface == nil {
		return nil, errors.New("Interface of struct can not be null")
	}
	if len(separator) == 0 {
		return nil, errors.New("Separator can not be empty")
	}

	// Processing data
	fields, err := getFields(structInterface)
	if err != nil {
		return nil, err
	}

	sortedFields, err := sortFields(fields)
	if err != nil {
		return nil, err
	}

	splittedData := strings.Split(dsv, separator)

	// Validating the fields and the quantity of data in the dsv
	if len(splittedData) != len(sortedFields) {
		return nil, errors.New("Struct fields and dsv data have different lengths")
	}

	// Put data in the fields
	for index := 0; index < len(sortedFields); index++ {
		sortedFields[index].value = strings.Trim(splittedData[index], " ")
	}
}
