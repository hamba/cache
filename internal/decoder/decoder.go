package decoder

import (
	"errors"
	"strconv"
)

type StringDecoder struct{}

func (d StringDecoder) Bool(v interface{}) (bool, error) {
	b, ok := v.([]byte)
	if !ok {
		return false, errors.New("decoder: expected byte slice")
	}

	return string(b) == "1", nil
}

func (d StringDecoder) Bytes(v interface{}) ([]byte, error) {
	b, ok := v.([]byte)
	if !ok {
		return nil, errors.New("decoder: expected byte slice")
	}

	return b, nil
}

func (d StringDecoder) Int64(v interface{}) (int64, error) {
	b, ok := v.([]byte)
	if !ok {
		return 0, errors.New("decoder: expected byte slice")
	}

	return strconv.ParseInt(string(b), 10, 64)
}

func (d StringDecoder) Uint64(v interface{}) (uint64, error) {
	b, ok := v.([]byte)
	if !ok {
		return 0, errors.New("decoder: expected byte slice")
	}

	return strconv.ParseUint(string(b), 10, 64)
}

func (d StringDecoder) Float64(v interface{}) (float64, error) {
	b, ok := v.([]byte)
	if !ok {
		return 0, errors.New("decoder: expected byte slice")
	}

	return strconv.ParseFloat(string(b), 64)
}

func (d StringDecoder) String(v interface{}) (string, error) {
	b, ok := v.([]byte)
	if !ok {
		return "", errors.New("decoder: expected byte slice")
	}

	return string(b), nil
}
