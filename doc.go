// Package appvalue provides a way to unmarshal app.Value objects into Go structs, similar to how json.Unmarshal and xml.Unmarshal work.
//
// # Struct Tags
//
// The following struct tags are supported:
//
//   - `javascript:"<name>"`: The name of the property in the javascript object (such as `javascript:"favoriteColor"`).
//   - `javascript:"-"` : Skip the field.
//   - `javascript:"@this"`: Put the raw app.Value into this field.
//
// # Special Support
//
// The following special types are supported:
//
//   - app.Value: This is a raw app.Value; no processing will be done on it.
//   - appvalue.Timestamp: This is a `time.Time` value; you may cast it via `time.Time(value)`.
package appvalue
