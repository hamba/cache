package cache_test

import (
	"testing"

	"github.com/hamba/cache/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNullCache_Get(t *testing.T) {
	i := cache.Null.Get("test")
	v, err := i.Bytes()

	require.NoError(t, err)
	assert.Equal(t, []byte{}, v)
}

func TestNullCache_GetBool(t *testing.T) {
	i := cache.Null.Get("test")
	b, err := i.Bool()

	require.NoError(t, err)
	assert.Equal(t, false, b)
}

func TestNullCache_GetInt64(t *testing.T) {
	i := cache.Null.Get("test")
	b, err := i.Int64()

	require.NoError(t, err)
	assert.Equal(t, int64(0), b)
}

func TestNullCache_GetUint64(t *testing.T) {
	i := cache.Null.Get("test")
	b, err := i.Uint64()

	require.NoError(t, err)
	assert.Equal(t, uint64(0), b)
}

func TestNullCache_GetFloat64(t *testing.T) {
	i := cache.Null.Get("test")
	v, err := i.Float64()

	require.NoError(t, err)
	assert.Equal(t, float64(0), v)
}

func TestNullCache_GetString(t *testing.T) {
	i := cache.Null.Get("test")
	v, err := i.String()

	require.NoError(t, err)
	assert.Equal(t, "", v)
}

func TestNullCache_GetMulti(t *testing.T) {
	v, err := cache.Null.GetMulti("test")

	require.NoError(t, err)
	assert.Len(t, v, 0)
}

func TestNullCache_Set(t *testing.T) {
	assert.NoError(t, cache.Null.Set("test", 1, 0))
}

func TestNullCache_Add(t *testing.T) {
	assert.NoError(t, cache.Null.Add("test", 1, 0))
}

func TestNullCache_Replace(t *testing.T) {
	assert.NoError(t, cache.Null.Replace("test", 1, 0))
}

func TestNullCache_Delete(t *testing.T) {
	assert.NoError(t, cache.Null.Delete("test"))
}

func TestNullCache_Inc(t *testing.T) {
	v, err := cache.Null.Inc("test", 1)

	require.NoError(t, err)
	assert.Equal(t, int64(0), v)
}

func TestNullCache_Dec(t *testing.T) {
	v, err := cache.Null.Dec("test", 1)

	require.NoError(t, err)
	assert.Equal(t, int64(0), v)
}
