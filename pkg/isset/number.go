package isset

import (
	"encoding/json"
	"github.com/ektoric/isset/internal"
)

// Int is a common JSON `number` primitive
type Int struct {
	Value         int
	IsNullPrivate bool // private field, but needs to be go-exported for reflection
	IsSetPrivate  bool // private field, but needs to be go-exported for reflection
}

func NewInt(v int) Int {
	return Int{Value: v, IsNullPrivate: false, IsSetPrivate: true}
}
func NewIntPtr(v int) *Int {
	return &Int{Value: v, IsNullPrivate: false, IsSetPrivate: true}
}

// IsSet if this field was set in JSON
func (i *Int) IsSet() bool {
	return i.IsSetPrivate
}

// IsNull if this field was set in JSON to json-null
func (i *Int) IsNull() bool {
	return i.IsNullPrivate
}

// Private method, used for testing
func (i *Int) nullValue() {
	i.IsNullPrivate = true
}

// UnmarshalJSON implements json.Unmarshaler.
// Unmarshal the field from a JSON primitive
func (i *Int) UnmarshalJSON(data []byte) error {
	return internal.JsonUnmarshalValue(i, data, &i.Value)
}

// MarshalJSON implements json.Marshaler.
// It will encode null if !IsSet()
func (t Int) MarshalJSON() ([]byte, error) {
	if !t.IsSet() || t.IsNull() {
		return []byte("null"), nil
	}
	return json.Marshal(t.Value)
}

// Float is a JSON `number` primitive
type Float struct {
	Value         float64
	IsNullPrivate bool // private field, but needs to be go-exported for reflection
	IsSetPrivate  bool // private field, but needs to be go-exported for reflection
}

func NewFloat(v float64) Float {
	return Float{Value: v, IsNullPrivate: false, IsSetPrivate: true}
}
func NewFloatPtr(v float64) *Float {
	return &Float{Value: v, IsNullPrivate: false, IsSetPrivate: true}
}

// IsSet if this field was set in JSON
func (i *Float) IsSet() bool {
	return i.IsSetPrivate
}

// IsNull if this field was set in JSON to json-null
func (i *Float) IsNull() bool {
	return i.IsNullPrivate
}

// Private method, used for testing
func (i *Float) nullValue() {
	i.IsNullPrivate = true
}

// UnmarshalJSON implements json.Unmarshaler.
// Unmarshal the field from a JSON primitive
func (i *Float) UnmarshalJSON(data []byte) error {
	return internal.JsonUnmarshalValue(i, data, &i.Value)
}

// MarshalJSON implements json.Marshaler.
// It will encode null if !IsSet()
func (t Float) MarshalJSON() ([]byte, error) {
	if !t.IsSet() || t.IsNull() {
		return []byte("null"), nil
	}
	return json.Marshal(t.Value)
}
