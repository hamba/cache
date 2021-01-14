package decoder

import (
	"errors"
	"strconv"
)

// StringDecoder decodes a string into various types.
type StringDecoder struct{}

// Bool coverts a string to a boolean.
func (d StringDecoder) Bool(v interface{}) (bool, error) {
	b, ok := v.([]byte)
	if !ok {
		return false, errors.New("decoder: expected byte slice")
	}

	return string(b) == "1", nil
}

// Bytes converts a string to bytes.
func (d StringDecoder) Bytes(v interface{}) ([]byte, error) {
	b, ok := v.([]byte)
	if !ok {
		return nil, errors.New("decoder: expected byte slice")
	}

	return b, nil
}

// Int64 converts a string to an int64.
func (d StringDecoder) Int64(v interface{}) (int64, error) {
	b, ok := v.([]byte)
	if !ok {
		return 0, errors.New("decoder: expected byte slice")
	}

	return strconv.ParseInt(string(b), 10, 64)
}

// Uint64 converts a string to a uint64.
func (d StringDecoder) Uint64(v interface{}) (uint64, error) {
	b, ok := v.([]byte)
	if !ok {
		return 0, errors.New("decoder: expected byte slice")
	}

	return strconv.ParseUint(string(b), 10, 64)
}

// Float64 converts a string to a float64.
func (d StringDecoder) Float64(v interface{}) (float64, error) {
	b, ok := v.([]byte)
	if !ok {
		return 0, errors.New("decoder: expected byte slice")
	}

	return strconv.ParseFloat(string(b), 64)
}

// String converts a string to a string.
func (d StringDecoder) String(v interface{}) (string, error) {
	b, ok := v.([]byte)
	if !ok {
		return "", errors.New("decoder: expected byte slice")
	}

	return string(b), nil
}
