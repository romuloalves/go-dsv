package dsv

import (
	"errors"
	"testing"

	"github.com/tj/assert"
)

func TestSortFields(t *testing.T) {
	fieldsAssert := []field{
		field{
			index: -1,
		}, field{
			index: 0,
		}, field{
			index: 0,
		}, field{
			index: 1,
		}, field{
			index: 2,
		}, field{
			index: 3,
		}, field{
			index: 4,
		}, field{
			index: 5,
		}, field{
			index: 7,
		}, field{
			index: 7,
		}, field{
			index: 12,
		},
	}

	fields := []field{
		field{
			index: 7,
		}, field{
			index: 2,
		}, field{
			index: 0,
		}, field{
			index: 0,
		}, field{
			index: 12,
		}, field{
			index: 5,
		}, field{
			index: 7,
		}, field{
			index: 3,
		}, field{
			index: 4,
		}, field{
			index: 1,
		}, field{
			index: -1,
		},
	}

	sortedFields, err := sortFields(fields)
	if err != nil {
		t.Fatal("Error should be nil")
	}

	assert.EqualValues(t, fieldsAssert, sortedFields, "Arrays should be equals")
}

func TestPadValueWithFiftyEmpty(t *testing.T) {
	assertFieldValue := "                                         123456789"
	fieldTest := field{
		index:       1,
		length:      50,
		paddingChar: " ",
		value:       "123456789",
	}

	fieldWithPad, err := padValue(fieldTest)
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, assertFieldValue, fieldWithPad, "Value should be with padding of 41 empty spaces")
}

func TestPadValueWithFourZeros(t *testing.T) {
	assertFieldValue := "0000123456789"
	fieldTest := field{
		index:       1,
		length:      13,
		paddingChar: "0",
		value:       "123456789",
	}

	fieldWithPad, err := padValue(fieldTest)
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, assertFieldValue, fieldWithPad, "Value should be with padding of 4 zeros")
}

func TestPadValueWithRightLength(t *testing.T) {
	assertFieldValue := "123456789"
	fieldTest := field{
		index:       1,
		length:      9,
		paddingChar: "0",
		value:       "123456789",
	}

	fieldWithPad, err := padValue(fieldTest)
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, assertFieldValue, fieldWithPad, "Value should keep unchanged")
}

func TestPadValueShouldNotAcceptPaddingCharWithMoreThanOneCharacter(t *testing.T) {
	paddingCharErrorAssert := errors.New("The padding char should be one char length")
	fieldTest := field{
		index:       1,
		length:      10,
		paddingChar: "00",
		value:       "123456789",
	}

	_, err := padValue(fieldTest)
	if err == nil {
		t.Fatalf("padValue should return an error about padding char length")
	}

	assert.EqualValues(t, paddingCharErrorAssert, err, "Error should be about padding char length")
}
