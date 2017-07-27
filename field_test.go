package dsv

import (
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
