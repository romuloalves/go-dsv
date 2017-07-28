package dsv

import "errors"

// StructToDSV returns a dsv string from the struct
func StructToDSV(i interface{}, separator string) (string, error) {
	if i == nil {
		return "", errors.New("The struct in the parameters can not be null")
	}
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
