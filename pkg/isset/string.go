package isset

import (
	"encoding/json"
	"github.com/ektoric/isset/internal"
)

// String is a JSON primitive
type String struct {
	Value         string
	IsNullPrivate bool
	IsSetPrivate  bool
}

func NewString(v string) String {
	return String{Value: v, IsNullPrivate: false, IsSetPrivate: true}
}
func NewStringPtr(v string) *String {
	return &String{Value: v, IsNullPrivate: false, IsSetPrivate: true}
}

// IsSet if this field was set in JSON
func (i *String) IsSet() bool {
	return i.IsSetPrivate
}

// IsNull if this field was set in JSON to json-null
func (i *String) IsNull() bool {
	return i.IsNullPrivate
}

// Private method, used for testing
func (i *String) nullValue() {
	i.IsNullPrivate = true
}

// UnmarshalJSON implements json.Unmarshaler.
// Unmarshal the field from a JSON primitive
func (i *String) UnmarshalJSON(data []byte) error {
	return internal.JsonUnmarshalValue(i, data, &i.Value)
}

// MarshalJSON implements json.Marshaler.
// It will encode null if !IsSet()
func (t String) MarshalJSON() ([]byte, error) {
	if !t.IsSet() || t.IsNull() {
		return []byte("null"), nil
	}
	return json.Marshal(t.Value)
}
