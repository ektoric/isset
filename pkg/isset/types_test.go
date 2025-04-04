package isset

import (
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"testing"
)

type IsSetType interface {
	IsSet() bool
	IsNull() bool
	nullValue()
}

type DummyType struct {
	TInt    Int    `json:"tint"`
	TFloat  Float  `json:"tfloat"`
	TBool   Bool   `json:"tbool"`
	TString String `json:"tstring"`
}
type tCase struct {
	name     string
	field    IsSetType
	badjson  []byte
	unset    []byte
	nullset  []byte
	nilset   []byte
	nilvalue interface{}
}

func (suite *JsonTypeTester) TestJsonPrimitiveTypes() {
	var value DummyType
	var err error
	matrix := []tCase{
		{
			name:     "Int",
			field:    &value.TInt,
			badjson:  []byte(`{ "tint":true}`),
			unset:    []byte(`{}`),
			nullset:  []byte(`{"tint":null}`),
			nilset:   []byte(`{"tint":0}`),
			nilvalue: &Int{IsSetPrivate: true, Value: 0},
		},
		{
			name:     "Float",
			field:    &value.TFloat,
			badjson:  []byte(`{ "tfloat":true}`),
			unset:    []byte(`{}`),
			nullset:  []byte(`{"tfloat":null}`),
			nilset:   []byte(`{"tfloat":0.0}`),
			nilvalue: &Float{IsSetPrivate: true, Value: 0.0},
		},
		{
			name:     "Bool",
			field:    &value.TBool,
			badjson:  []byte(`{ "tbool":0.1}`),
			unset:    []byte(`{}`),
			nullset:  []byte(`{"tbool":null}`),
			nilset:   []byte(`{"tbool":false}`),
			nilvalue: &Bool{IsSetPrivate: true, Value: false},
		},
		{
			name:     "String",
			field:    &value.TString,
			badjson:  []byte(`{ "tstring":3}`),
			unset:    []byte(`{}`),
			nullset:  []byte(`{"tstring":null}`),
			nilset:   []byte(`{"tstring":""}`),
			nilvalue: &String{IsSetPrivate: true, Value: ""},
		},
	}
	for _, entry := range matrix {
		err = json.Unmarshal(entry.badjson, &value)
		suite.Require().Error(err, entry.name)
		suite.Assert().Regexp(`json: cannot unmarshal .* into Go struct field .* of type .*`, err.Error(), entry.name)

		value = DummyType{} // clear it
		err = json.Unmarshal(entry.unset, &value)
		suite.Require().NoError(err, entry.name)
		suite.Assert().False(entry.field.IsSet(), entry.name)

		value = DummyType{} // clear it
		err = json.Unmarshal(entry.nullset, &value)
		suite.Require().NoError(err, entry.name)
		suite.Assert().True(entry.field.IsSet(), entry.name)
		suite.Assert().True(entry.field.IsNull(), entry.name)

		value = DummyType{} // clear it
		err = json.Unmarshal(entry.nilset, &value)
		suite.Require().NoError(err, entry.name)
		suite.Assert().True(entry.field.IsSet(), entry.name)
		suite.Assert().False(entry.field.IsNull(), entry.name)
		suite.Assert().Equal(entry.nilvalue, entry.field, entry.name)
	}
}

type DummyEgressType struct {
	TInt    *Int    `json:"tint,omitempty"`
	TFloat  *Float  `json:"tfloat,omitempty"`
	TBool   *Bool   `json:"tbool,omitempty"`
	TString *String `json:"tstring,omitempty"`
}

type tCaseEgress struct {
	name         string
	nilvalue     func()
	field        func() IsSetType
	nilexpected  []byte
	nullexpected []byte
}

func (suite *JsonTypeTester) TestJsonMarshal() {
	var value = DummyEgressType{}
	var err error
	var msg []byte

	matrix := []tCaseEgress{
		{
			name:         "Int",
			field:        func() IsSetType { return value.TInt },
			nilvalue:     func() { value.TInt = NewIntPtr(0) },
			nilexpected:  []byte(`{"tint":0}`),
			nullexpected: []byte(`{"tint":null}`),
		},
		{
			name:         "Float",
			field:        func() IsSetType { return value.TFloat },
			nilvalue:     func() { value.TFloat = NewFloatPtr(0.0) },
			nilexpected:  []byte(`{"tfloat":0}`),
			nullexpected: []byte(`{"tfloat":null}`),
		},
		{
			name:         "Bool",
			field:        func() IsSetType { return value.TBool },
			nilvalue:     func() { value.TBool = NewBoolPtr(false) },
			nilexpected:  []byte(`{"tbool":false}`),
			nullexpected: []byte(`{"tbool":null}`),
		},
		{
			name:         "String",
			field:        func() IsSetType { return value.TString },
			nilvalue:     func() { value.TString = NewStringPtr("") },
			nilexpected:  []byte(`{"tstring":""}`),
			nullexpected: []byte(`{"tstring":null}`),
		},
	}
	for _, entry := range matrix {
		// clear state
		value = DummyEgressType{}
		// init `value`-template with field, marshal, verify
		entry.nilvalue()
		msg, err = json.Marshal(value)
		suite.Require().NoError(err, entry.name)
		suite.Assert().Equal(msg, entry.nilexpected, entry.name)
		// set `value`-template's field as `null`, marshal, verify
		entry.field().nullValue()
		msg, err = json.Marshal(value)
		suite.Require().NoError(err, entry.name)
		suite.Assert().Equal(msg, entry.nullexpected, entry.name)
	}
}

type JsonTypeTester struct {
	suite.Suite
}

// Testing entry point.
func TestJsonType(t *testing.T) {
	suite.Run(t, new(JsonTypeTester))
}
