package memcache

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWithIdleConns(t *testing.T) {
	c := &memcache.Client{}

	WithIdleConns(12)(c)

	assert.Equal(t, 12, c.MaxIdleConns)
}

func TestWithTimeout(t *testing.T) {
	c := &memcache.Client{}

	WithTimeout(time.Second)(c)

	assert.Equal(t, time.Second, c.Timeout)
}

func TestNewMemcache(t *testing.T) {
	c := New("test", WithIdleConns(12))

	assert.Equal(t, 12, c.client.MaxIdleConns)
}

func TestEncoderError(t *testing.T) {
	c := Memcache{
		enc: func(v interface{}) ([]byte, error) {
			return nil, errors.New("test error")
		},
	}

	assert.EqualError(t, c.Add(context.Background(), "test", 1, 0), "test error")
	assert.EqualError(t, c.Set(context.Background(), "test", 1, 0), "test error")
	assert.EqualError(t, c.Replace(context.Background(), "test", 1, 0), "test error")
}

func TestByteEncode(t *testing.T) {
	tests := []struct {
		name string
		v    interface{}
		want []byte
	}{
		{
			name: "bool true",
			v:    true,
			want: []byte("1"),
		},
		{
			name: "bool false",
			v:    false,
			want: []byte("0"),
		},
		{
			name: "int64",
			v:    int64(10),
			want: []byte("10"),
		},
		{
			name: "uint64",
			v:    uint64(10),
			want: []byte("10"),
		},
		{
			name: "float64",
			v:    float64(10.34),
			want: []byte("10.340000"),
		},
		{
			name: "string",
			v:    "foobar",
			want: []byte("foobar"),
		},
		{
			name: "struct",
			v:    struct{ A int }{1},
			want: []byte(`{"A":1}`),
		},
		{
			name: "string slice",
			v:    []string{"foo", "bar"},
			want: []byte(`["foo","bar"]`),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			got, err := memcacheEncoder(test.v)

			require.NoError(t, err)
			assert.Equal(t, test.want, got)
		})
	}
}
