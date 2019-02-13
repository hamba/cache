package decoder_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/hamba/cache/internal/decoder"
)

func TestStringDecoder_Bool(t *testing.T) {
	tests := []struct {
		name    string
		in      interface{}
		want    bool
		wantErr bool
	}{
		{
			name:    "Valid",
			in:      []byte("1"),
			want:    true,
			wantErr: false,
		},
		{
			name:    "Invalid Type",
			in:      struct{}{},
			want:    false,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dec := decoder.StringDecoder{}

			got, err := dec.Bool(tt.in)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestStringDecoder_Bytes(t *testing.T) {
	tests := []struct {
		name    string
		in      interface{}
		want    []byte
		wantErr bool
	}{
		{
			name:    "Valid",
			in:      []byte{0x01},
			want:    []byte{0x01},
			wantErr: false,
		},
		{
			name:    "Invalid Type",
			in:      struct{}{},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dec := decoder.StringDecoder{}

			got, err := dec.Bytes(tt.in)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestStringDecoder_Int64(t *testing.T) {
	tests := []struct {
		name    string
		in      interface{}
		want    int64
		wantErr bool
	}{
		{
			name:    "Valid",
			in:      []byte("1"),
			want:    1,
			wantErr: false,
		},
		{
			name:    "Invalid Type",
			in:      struct{}{},
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dec := decoder.StringDecoder{}

			got, err := dec.Int64(tt.in)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestStringDecoder_Uint64(t *testing.T) {
	tests := []struct {
		name    string
		in      interface{}
		want    uint64
		wantErr bool
	}{
		{
			name:    "Valid",
			in:      []byte("1"),
			want:    1,
			wantErr: false,
		},
		{
			name:    "Invalid Type",
			in:      struct{}{},
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dec := decoder.StringDecoder{}

			got, err := dec.Uint64(tt.in)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestStringDecoder_Float64(t *testing.T) {
	tests := []struct {
		name    string
		in      interface{}
		want    float64
		wantErr bool
	}{
		{
			name:    "Valid",
			in:      []byte("2.3"),
			want:    2.3,
			wantErr: false,
		},
		{
			name:    "Invalid Type",
			in:      struct{}{},
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dec := decoder.StringDecoder{}

			got, err := dec.Float64(tt.in)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestStringDecoder_String(t *testing.T) {
	tests := []struct {
		name    string
		in      interface{}
		want    string
		wantErr bool
	}{
		{
			name:    "Valid",
			in:      []byte("foobar"),
			want:    "foobar",
			wantErr: false,
		},
		{
			name:    "Invalid Type",
			in:      struct{}{},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dec := decoder.StringDecoder{}

			got, err := dec.String(tt.in)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
