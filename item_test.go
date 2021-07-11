package cache_test

import (
	"errors"
	"testing"

	"github.com/hamba/cache/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestItem_Bool(t *testing.T) {
	dec := new(mockDecoder)
	dec.On("Bool", []byte("1")).Return(true, nil)
	item := cache.NewItem(dec, []byte("1"), nil)

	got, err := item.Bool()

	require.NoError(t, err)
	assert.Equal(t, true, got)

	dec.AssertExpectations(t)
}

func TestItem_BoolDecoderError(t *testing.T) {
	dec := new(mockDecoder)
	dec.On("Bool", []byte("1")).Return(false, errors.New("test"))
	item := cache.NewItem(dec, []byte("1"), nil)

	_, err := item.Bool()

	require.Error(t, err)

	dec.AssertExpectations(t)
}

func TestItem_BoolError(t *testing.T) {
	dec := new(mockDecoder)
	item := cache.NewItem(dec, []byte("1"), errors.New("test"))

	_, err := item.Bool()

	assert.Error(t, err)
}

func TestItem_Bytes(t *testing.T) {
	dec := new(mockDecoder)
	dec.On("Bytes", []byte{0x01}).Return([]byte{0x01}, nil)
	item := cache.NewItem(dec, []byte{0x01}, nil)

	got, err := item.Bytes()

	require.NoError(t, err)
	assert.Equal(t, []byte{0x01}, got)

	dec.AssertExpectations(t)
}

func TestItem_BytesDecoderError(t *testing.T) {
	dec := new(mockDecoder)
	dec.On("Bytes", []byte{0x01}).Return([]byte{}, errors.New("test"))
	item := cache.NewItem(dec, []byte{0x01}, nil)

	_, err := item.Bytes()

	require.Error(t, err)

	dec.AssertExpectations(t)
}

func TestItem_BytesError(t *testing.T) {
	dec := new(mockDecoder)
	item := cache.NewItem(dec, []byte{0x01}, errors.New("test"))

	_, err := item.Bytes()

	assert.Error(t, err)
}

func TestItem_Int64(t *testing.T) {
	dec := new(mockDecoder)
	dec.On("Int64", []byte("2")).Return(int64(2), nil)
	item := cache.NewItem(dec, []byte("2"), nil)

	got, err := item.Int64()

	require.NoError(t, err)
	assert.Equal(t, int64(2), got)

	dec.AssertExpectations(t)
}

func TestItem_Int64DecoderError(t *testing.T) {
	dec := new(mockDecoder)
	dec.On("Int64", []byte("2")).Return(int64(0), errors.New("test"))
	item := cache.NewItem(dec, []byte("2"), nil)

	_, err := item.Int64()

	require.Error(t, err)

	dec.AssertExpectations(t)
}

func TestItem_Int64Error(t *testing.T) {
	dec := new(mockDecoder)
	item := cache.NewItem(dec, []byte("2"), errors.New("test"))

	_, err := item.Int64()

	assert.Error(t, err)
}

func TestItem_Uint64(t *testing.T) {
	dec := new(mockDecoder)
	dec.On("Uint64", []byte("2")).Return(uint64(2), nil)
	item := cache.NewItem(dec, []byte("2"), nil)

	got, err := item.Uint64()

	require.NoError(t, err)
	assert.Equal(t, uint64(2), got)

	dec.AssertExpectations(t)
}

func TestItem_Uint64DecoderError(t *testing.T) {
	dec := new(mockDecoder)
	dec.On("Uint64", []byte("2")).Return(uint64(0), errors.New("test"))
	item := cache.NewItem(dec, []byte("2"), nil)

	_, err := item.Uint64()

	require.Error(t, err)

	dec.AssertExpectations(t)
}

func TestItem_Uint64Error(t *testing.T) {
	dec := new(mockDecoder)
	item := cache.NewItem(dec, []byte("2"), errors.New("test"))

	_, err := item.Uint64()

	assert.Error(t, err)
}

func TestItem_Float64(t *testing.T) {
	dec := new(mockDecoder)
	dec.On("Float64", []byte("2.3")).Return(float64(2.3), nil)
	item := cache.NewItem(dec, []byte("2.3"), nil)

	got, err := item.Float64()

	require.NoError(t, err)
	assert.Equal(t, float64(2.3), got)

	dec.AssertExpectations(t)
}

func TestItem_Float64DecoderError(t *testing.T) {
	dec := new(mockDecoder)
	dec.On("Float64", []byte("2.3")).Return(float64(0), errors.New("test"))
	item := cache.NewItem(dec, []byte("2.3"), nil)

	_, err := item.Float64()

	require.Error(t, err)

	dec.AssertExpectations(t)
}

func TestItem_Float64Error(t *testing.T) {
	dec := new(mockDecoder)
	item := cache.NewItem(dec, []byte("2.3"), errors.New("test"))

	_, err := item.Float64()

	assert.Error(t, err)
}

func TestItem_String(t *testing.T) {
	dec := new(mockDecoder)
	dec.On("String", []byte("test")).Return("test", nil)
	item := cache.NewItem(dec, []byte("test"), nil)

	got, err := item.String()

	require.NoError(t, err)
	assert.Equal(t, "test", got)

	dec.AssertExpectations(t)
}

func TestItem_StringDecoderError(t *testing.T) {
	dec := new(mockDecoder)
	dec.On("String", []byte("test")).Return("", errors.New("test"))
	item := cache.NewItem(dec, []byte("test"), nil)

	_, err := item.String()

	require.Error(t, err)

	dec.AssertExpectations(t)
}

func TestItem_StringError(t *testing.T) {
	dec := new(mockDecoder)
	item := cache.NewItem(dec, []byte("test"), errors.New("test"))

	_, err := item.String()

	assert.Error(t, err)
}

type mockDecoder struct {
	mock.Mock
}

func (m *mockDecoder) Bool(v interface{}) (bool, error) {
	args := m.Called(v)

	return args.Bool(0), args.Error(1)
}

func (m *mockDecoder) Bytes(v interface{}) ([]byte, error) {
	args := m.Called(v)

	return args.Get(0).([]byte), args.Error(1)
}

func (m *mockDecoder) Int64(v interface{}) (int64, error) {
	args := m.Called(v)

	return args.Get(0).(int64), args.Error(1)
}

func (m *mockDecoder) Uint64(v interface{}) (uint64, error) {
	args := m.Called(v)

	return args.Get(0).(uint64), args.Error(1)
}

func (m *mockDecoder) Float64(v interface{}) (float64, error) {
	args := m.Called(v)

	return args.Get(0).(float64), args.Error(1)
}

func (m *mockDecoder) String(v interface{}) (string, error) {
	args := m.Called(v)

	return args.String(0), args.Error(1)
}
