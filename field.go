package dsv

// field represents a field with data in a struct
type field struct {
	index       int
	length      int
	value       string
	paddingChar string
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
