package internal

import (
	"encoding/json"
	"reflect"
)

func JsonUnmarshalValue(i interface{}, data []byte, valueIface interface{}) error {
	reflect.ValueOf(i).Elem().FieldByName("IsSetPrivate").SetBool(true)

	if string(data) == "null" {
		// The key was set to null
		reflect.ValueOf(i).Elem().FieldByName("IsNullPrivate").SetBool(true)
		return nil
	}

	if err := json.Unmarshal(data, valueIface); err != nil {
		return err
	}
	reflect.ValueOf(i).Elem().FieldByName("Value").Set(reflect.ValueOf(valueIface).Elem())
	return nil
}
