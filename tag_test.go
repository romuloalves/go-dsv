package dsv

import (
	"testing"
)

func TestGetIndexFromTag(t *testing.T) {
	testAssert := 3
	testParam := []string{"3", "20"}

	index, err := getIndexFromTag(testParam)

	if err != nil {
		t.Fatal(err)
	}

	if index != testAssert {
		t.Fatalf("The value of the field should be %d and not %d", testAssert, index)
	}
}

func TestGetLengthFromTag(t *testing.T) {
	testAssert := 20
	testParam := []string{"3", "20"}

	length, err := getLengthFromTag(testParam)

	if err != nil {
		t.Fatal(err)
	}

	if length != testAssert {
		t.Fatalf("The value of the field should be %d and not %d", testAssert, length)
	}
}

func TestGetPaddingCharFromTag(t *testing.T) {
	testAssert := "|"
	testParam := []string{"3", "20", "|", "true"}

	paddingChar, err := getPaddingCharFromTag(testParam, " ")

	if err != nil {
		t.Fatal(err)
	}

	if paddingChar != testAssert {
		t.Fatalf("The value of the field should be %s and not %s", testAssert, paddingChar)
	}
}

func TestGetPaddingCharFromTagWithDefault(t *testing.T) {
	testAssert := " "
	testParam := []string{"3", "20", "-", "true"}

	paddingChar, err := getPaddingCharFromTag(testParam, " ")

	if err != nil {
		t.Fatal(err)
	}

	if paddingChar != testAssert {
		t.Fatalf("The value of the field should be %s and not %s", testAssert, paddingChar)
	}
}

func TestGetPaddingCharFromTagEmptyWithDefault(t *testing.T) {
	testAssert := " "
	testParam := []string{"3", "20"}

	paddingChar, err := getPaddingCharFromTag(testParam, " ")

	if err != nil {
		t.Fatal(err)
	}

	if paddingChar != testAssert {
		t.Fatalf("The value of the field should be %s and not %s", testAssert, paddingChar)
	}
}

func TestGetPaddingRightFromTag(t *testing.T) {
	testAssert := true
	testParam := []string{"3", "20", "|", "true"}

	paddingRight, err := getPaddingRightFromTag(testParam, false)

	if err != nil {
		t.Fatal(err)
	}

	if paddingRight != testAssert {
		t.Fatalf("The value of the field should be %v and not %v", testAssert, paddingRight)
	}
}

func TestGetPaddingRightFromTagWithDefault(t *testing.T) {
	testAssert := false
	testParam := []string{"3", "20", "-", "-"}

	paddingRight, err := getPaddingRightFromTag(testParam, false)

	if err != nil {
		t.Fatal(err)
	}

	if paddingRight != testAssert {
		t.Fatalf("The value of the field should be %v and not %v", testAssert, paddingRight)
	}
}

func TestGetPaddingRightFromTagEmptyWithDefault(t *testing.T) {
	testAssert := false
	testParam := []string{"3", "20"}

	paddingRight, err := getPaddingRightFromTag(testParam, false)

	if err != nil {
		t.Fatal(err)
	}

	if paddingRight != testAssert {
		t.Fatalf("The value of the field should be %v and not %v", testAssert, paddingRight)
	}
}
