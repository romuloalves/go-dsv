package dsv

import (
	"testing"

	"github.com/tj/assert"
)

func TestStructToDsvWithNoFields(t *testing.T) {
	type myStruct struct {
		FieldOne       string  `dsv:"1"`
		FieldNumberTwo int     `dsv:"2"`
		FieldThree     bool    `dsv:"3"`
		FieldFour      float32 `dsv:"4"`
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
		FieldPaddingLeft  string `dsv:"1,20,1"`
		FieldPaddingRight string `dsv:"2,27,7,true"`
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
		FieldTwo       string `dsv:"2"`
		FieldTwoButOne string `dsv:"1"`
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

func TestStructToDsvWithSameDefaultFlags(t *testing.T) {
	type myStruct struct {
		FieldOne string `dsv:"1,10,-,true"`
		FieldTwo string `dsv:"2,11,=,false"`
	}

	assertValue := "FieldOne  ;===FieldTwo"

	structData := &myStruct{
		FieldOne: "FieldOne",
		FieldTwo: "FieldTwo",
	}

	dsv, err := StructToDSV(structData, ";")

	if err != nil {
		t.Fatal(err)
	}

	if dsv != assertValue {
		t.Fatalf("Value of the line should be equals to \"%s\" and not \"%s\"", assertValue, dsv)
	}
}

func TestDsvToStruct(t *testing.T) {
	type myStruct struct {
		FieldOne   string  `dsv:"0"`
		FieldTwo   int     `dsv:"1"`
		FieldThree bool    `dsv:"2"`
		FieldFour  float32 `dsv:"3"`
	}

	assertStruct := &myStruct{
		FieldOne:   "fieldOneData",
		FieldTwo:   123,
		FieldThree: false,
		FieldFour:  43.34,
	}

	dsv := "fieldOneData|123|false|43.34"

	data, err := ToStruct(dsv, &myStruct{}, "|")
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, assertStruct, data, "Assert struct and returned struct from dsv should be equals")
}
