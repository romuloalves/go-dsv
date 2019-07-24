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
func ToStruct(dsv string, i interface{}, separator string) (interface{}, error) {
	// Validating inputs
	if len(dsv) == 0 {
		return nil, errors.New("DSV data can not be empty")
	}
	if i == nil {
		return nil, errors.New("Interface of struct can not be null")
	}
	if len(separator) == 0 {
		return nil, errors.New("Separator can not be empty")
	}

	// Processing data
	fields, err := getFields(i)
	if err != nil {
		return nil, err
	}
	fieldsLength := len(fields)

	splittedData := strings.Split(dsv, separator)
	splittedDataLength := len(splittedData)

	// Put data in the fields
	for dataIndex := 0; dataIndex < splittedDataLength; dataIndex++ {
		for fieldIndex := 0; fieldIndex < fieldsLength; fieldIndex++ {
			if fields[fieldIndex].index != dataIndex {
				continue
			}
			fields[fieldIndex].value = strings.Trim(splittedData[dataIndex], " ")
		}
	}

	interfaceWithData, err := setFieldsToStruct(fields, i)
	if err != nil {
		return nil, err
	}

	return interfaceWithData, nil
}
