package decoder_test

import (
	"testing"

	"github.com/hamba/cache/v2/internal/decoder"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStringDecoder_Bool(t *testing.T) {
	tests := []struct {
		name    string
		in      interface{}
		want    bool
		wantErr require.ErrorAssertionFunc
	}{
		{
			name:    "valid",
			in:      []byte("1"),
			want:    true,
			wantErr: require.NoError,
		},
		{
			name:    "invalid type",
			in:      struct{}{},
			want:    false,
			wantErr: require.Error,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			dec := decoder.StringDecoder{}

			got, err := dec.Bool(test.in)

			test.wantErr(t, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestStringDecoder_Bytes(t *testing.T) {
	tests := []struct {
		name    string
		in      interface{}
		want    []byte
		wantErr require.ErrorAssertionFunc
	}{
		{
			name:    "valid",
			in:      []byte{0x01},
			want:    []byte{0x01},
			wantErr: require.NoError,
		},
		{
			name:    "invalid type",
			in:      struct{}{},
			want:    nil,
			wantErr: require.Error,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			dec := decoder.StringDecoder{}

			got, err := dec.Bytes(test.in)

			test.wantErr(t, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestStringDecoder_Int64(t *testing.T) {
	tests := []struct {
		name    string
		in      interface{}
		want    int64
		wantErr require.ErrorAssertionFunc
	}{
		{
			name:    "valid",
			in:      []byte("1"),
			want:    1,
			wantErr: require.NoError,
		},
		{
			name:    "invalid type",
			in:      struct{}{},
			want:    0,
			wantErr: require.Error,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			dec := decoder.StringDecoder{}

			got, err := dec.Int64(test.in)

			test.wantErr(t, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestStringDecoder_Uint64(t *testing.T) {
	tests := []struct {
		name    string
		in      interface{}
		want    uint64
		wantErr require.ErrorAssertionFunc
	}{
		{
			name:    "valid",
			in:      []byte("1"),
			want:    1,
			wantErr: require.NoError,
		},
		{
			name:    "invalid type",
			in:      struct{}{},
			want:    0,
			wantErr: require.Error,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			dec := decoder.StringDecoder{}

			got, err := dec.Uint64(test.in)

			test.wantErr(t, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestStringDecoder_Float64(t *testing.T) {
	tests := []struct {
		name    string
		in      interface{}
		want    float64
		wantErr require.ErrorAssertionFunc
	}{
		{
			name:    "valid",
			in:      []byte("2.3"),
			want:    2.3,
			wantErr: require.NoError,
		},
		{
			name:    "invalid type",
			in:      struct{}{},
			want:    0,
			wantErr: require.Error,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			dec := decoder.StringDecoder{}

			got, err := dec.Float64(test.in)

			test.wantErr(t, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestStringDecoder_String(t *testing.T) {
	tests := []struct {
		name    string
		in      interface{}
		want    string
		wantErr require.ErrorAssertionFunc
	}{
		{
			name:    "valid",
			in:      []byte("foobar"),
			want:    "foobar",
			wantErr: require.NoError,
		},
		{
			name:    "invalid type",
			in:      struct{}{},
			want:    "",
			wantErr: require.Error,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			dec := decoder.StringDecoder{}

			got, err := dec.String(test.in)

			test.wantErr(t, err)
			assert.Equal(t, test.want, got)
		})
	}
}
