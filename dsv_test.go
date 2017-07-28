package dsv

import (
	"testing"

	"github.com/tj/assert"
)

func TestStructToDsv(t *testing.T) {
	type myStruct struct {
		FieldOne       string
		FieldNumberTwo int
		FieldThree     bool
		FieldFour      float32
	}

	assertValue := "fieldOne|123|false|0.43"

	structData := &myStruct{
		FieldOne:       "fieldOne",
		FieldNumberTwo: 123,
		FieldThree:     false,
		FieldFour:      0.43,
	}

	dsv, err := StructToDSV(structData, "|")
	if err != nil {
		t.Fatal(err)
	}

	if dsv != assertValue {
		t.Fatalf("Value of the line should be equals to \"%s\" and not \"%s\"", assertValue, dsv)
	}
}

func TestStructToDsvWithStringPaddingLeftAndRight(t *testing.T) {
	type myStruct struct {
		FieldPaddingLeft  string `length:"20" paddingChar:"1"`
		FieldPaddingRight string `length:"27" paddingChar:"7" paddingRight:"true"`
	}

	assertValue := "1111FieldPaddingLeft|FieldPaddingRight7777777777"

	structData := &myStruct{
		FieldPaddingLeft:  "FieldPaddingLeft",
		FieldPaddingRight: "FieldPaddingRight",
	}

	dsv, err := StructToDSV(structData, "|")
	if err != nil {
		t.Fatal(err)
	}

	if dsv != assertValue {
		t.Fatalf("Value of the line should be equals to \"%s\" and not \"%s\"", assertValue, dsv)
	}
}

func TestStructToDsvWithIndexes(t *testing.T) {
	type myStruct struct {
		FieldTwo       string `index:"2"`
		FieldTwoButOne string `index:"1"`
	}

	assertValue := "FieldTwoButOne|FieldTwo"

	structData := &myStruct{
		FieldTwo:       "FieldTwo",
		FieldTwoButOne: "FieldTwoButOne",
	}

	dsv, err := StructToDSV(structData, "|")
	if err != nil {
		t.Fatal(err)
	}

	if dsv != assertValue {
		t.Fatalf("Value of the line should be equals to \"%s\" and not \"%s\"", assertValue, dsv)
	}
}

func TestDsvToStruct(t *testing.T) {
	type myStruct struct {
		FieldOne   string
		FieldTwo   int
		FieldThree bool
	}

	assertStruct := &myStruct{
		FieldOne:   "fieldOneData",
		FieldTwo:   123,
		FieldThree: false,
	}

	dsv := "fieldOneData|123|false"

	data, err := ToStruct(dsv, myStruct{}, "|")
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, assertStruct, data, "Assert struct and returned struct from dsv should be equals")
}
