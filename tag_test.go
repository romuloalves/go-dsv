package dsv

import (
	"reflect"
	"testing"

	"github.com/tj/assert"
)

func TestGetStringTag(t *testing.T) {
	type stringTest struct {
		Abc string `type:"string"`
	}

	testAssert := "string"

	testParam := &stringTest{Abc: "value"}

	fields := reflect.ValueOf(testParam).Elem()
	typeField := fields.Type().Field(0)

	value, err := getStringTag(typeField.Tag, "type")
	if err != nil {
		t.Fatal(err)
	}

	if value != testAssert {
		t.Fatalf("The value of the field should be %s and not %s", testAssert, value)
	}
}

func TestGetIntegerTag(t *testing.T) {
	type stringTest struct {
		Abc string `type:"123"`
	}

	testAssert := 123

	testParam := &stringTest{Abc: "value"}

	fields := reflect.ValueOf(testParam).Elem()
	typeField := fields.Type().Field(0)

	value, err := getIntegerTag(typeField.Tag, "type", -1)
	if err != nil {
		t.Fatal(err)
	}

	if value != testAssert {
		t.Fatalf("The value of the field should be %d and not %d", testAssert, value)
	}
}

func TestGetIntegerTagWithDefaultValue(t *testing.T) {
	type stringTest struct {
		Abc string `type:""`
	}

	testAssert := -1

	testParam := &stringTest{Abc: "value"}

	fields := reflect.ValueOf(testParam).Elem()
	typeField := fields.Type().Field(0)

	value, err := getIntegerTag(typeField.Tag, "type", -1)
	if err != nil {
		t.Fatal(err)
	}

	if value != testAssert {
		t.Fatalf("The value of the field should be %d and not %d", testAssert, value)
	}
}

func TestGetBooleanTag(t *testing.T) {
	type stringTest struct {
		Abc string `type:"true"`
	}

	testAssert := true

	testParam := &stringTest{Abc: "value"}

	fields := reflect.ValueOf(testParam).Elem()
	typeField := fields.Type().Field(0)

	value, err := getBooleanTag(typeField.Tag, "type", true)
	if err != nil {
		t.Fatal(err)
	}

	if value != testAssert {
		t.Fatalf("The value of the field should be %v and not %v", testAssert, value)
	}
}

func TestGetBooleanTagWithDefaultValue(t *testing.T) {
	type stringTest struct {
		Abc string `type:""`
	}

	testAssert := false

	testParam := &stringTest{Abc: "value"}

	fields := reflect.ValueOf(testParam).Elem()
	typeField := fields.Type().Field(0)

	value, err := getBooleanTag(typeField.Tag, "type", false)
	if err != nil {
		t.Fatal(err)
	}

	if value != testAssert {
		t.Fatalf("The value of the field should be %v and not %v", testAssert, value)
	}
}

func TestGetFieldsWithIndex(t *testing.T) {
	type myTestStruct struct {
		Field string `index:"1"`
	}

	testAssert := []field{
		field{
			value:  "1234",
			index:  1,
			length: -1,
		},
	}

	test := &myTestStruct{Field: "1234"}

	fields, err := getFields(test)
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, testAssert, fields, "Fields and assert should be equals")
}

func TestGetFieldsWithIndexAndLength(t *testing.T) {
	type myTestStruct struct {
		Field string `index:"1" length:"50"`
	}

	testAssert := []field{
		field{
			value:  "1234",
			index:  1,
			length: 50,
		},
	}

	test := &myTestStruct{Field: "1234"}

	fields, err := getFields(test)
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, testAssert, fields, "Fields and assert should be equals")
}

func TestGetFieldsWithIndexLengthAndPaddingChar(t *testing.T) {
	type myTestStruct struct {
		Field string `index:"1" length:"50" paddingChar:"|"`
	}

	testAssert := []field{
		field{
			value:       "1234",
			index:       1,
			length:      50,
			paddingChar: "|",
		},
	}

	test := &myTestStruct{Field: "1234"}

	fields, err := getFields(test)
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, testAssert, fields, "Fields and assert should be equals")
}

func TestGetFieldsWithPaddingRight(t *testing.T) {
	type myTestStruct struct {
		Field string `paddingRight:"true"`
	}

	testAssert := []field{
		field{
			value:        "1234",
			index:        0,
			length:       -1,
			paddingRight: true,
		},
	}

	test := &myTestStruct{Field: "1234"}

	fields, err := getFields(test)
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, testAssert, fields, "Fields and assert should be equals")
}

func TestGetFieldsWithIndexLengthPaddingCharAndPaddingRight(t *testing.T) {
	type myTestStruct struct {
		Field string `index:"1" length:"50" paddingChar:"|" paddingRight:"true"`
	}

	testAssert := []field{
		field{
			value:        "1234",
			index:        1,
			length:       50,
			paddingChar:  "|",
			paddingRight: true,
		},
	}

	test := &myTestStruct{Field: "1234"}

	fields, err := getFields(test)
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, testAssert, fields, "Fields and assert should be equals")
}

func TestGetFieldsWithIndexLengthPaddingCharAndDefaultPaddingRight(t *testing.T) {
	type myTestStruct struct {
		Field string `index:"1" length:"50" paddingChar:"|"`
	}

	testAssert := []field{
		field{
			value:        "1234",
			index:        1,
			length:       50,
			paddingChar:  "|",
			paddingRight: false,
		},
	}

	test := &myTestStruct{Field: "1234"}

	fields, err := getFields(test)
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, testAssert, fields, "Fields and assert should be equals")
}
