package isset

import (
	"encoding/json"
	"github.com/ektoric/isset/internal"
)

// Bool is a JSON primitive
type Bool struct {
	Value         bool
	IsNullPrivate bool
	IsSetPrivate  bool
}

func NewBool(v bool) Bool {
	return Bool{Value: v, IsNullPrivate: false, IsSetPrivate: true}
}
func NewBoolPtr(v bool) *Bool {
	return &Bool{Value: v, IsNullPrivate: false, IsSetPrivate: true}
}

// IsSet if this field was set in JSON
func (i *Bool) IsSet() bool {
	return i.IsSetPrivate
}

// IsNull if this field was set in JSON to json-null
func (i *Bool) IsNull() bool {
	return i.IsNullPrivate
}

// Private method, used for testing
func (i *Bool) nullValue() {
	i.IsNullPrivate = true
}

// UnmarshalJSON implements json.Unmarshaler.
// Unmarshal the field from a JSON primitive
func (i *Bool) UnmarshalJSON(data []byte) error {
	return internal.JsonUnmarshalValue(i, data, &i.Value)
}

// MarshalJSON implements json.Marshaler.
// It will encode null if !IsSet()
func (t Bool) MarshalJSON() ([]byte, error) {
	if !t.IsSet() || t.IsNull() {
		return []byte("null"), nil
	}
	return json.Marshal(t.Value)
}
