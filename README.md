## isset

`isset` is a library with the ability to JSON unmarshal while knowing if a field existed in the JSON

## Interfaces
- Implements `json.Unmarshaler`, and `json.Marshaler` for primitives.

## Concern Addressed
There are many blogs and repos that describe the concern of golang and JSON impedance mismatch of `null` vs zero-value.
We often describe the need to distinguish between these states:
* a non-zero value
* a zero value (JSON: "", false, 0)
* JSON `null`

While the first is easily distinguished, and the second and third can be distinguished by using
pointer fields (`*string`), there is actually a fourth case that is not called out:
* the field is not specified in the JSON.

If we use the common pointer field in the struct (`Field *string`) if the field does not exist
in the JSON, the object is unmarshaled but collapsed to `null`.

### PATCH
Being able to distinguish if a "field does not exist" vs "field exists and is set to null" is
especially import for HTTP PATCH.  _Explicitly_ specifying the value as `null` may have a different
semantic meaning (e.g. "reset to default") different from _explicitly_ not including the field 
(e.g. "do not change").
