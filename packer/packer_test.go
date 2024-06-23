package packer

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

// pack

func TestPackEmptyString(t *testing.T) {

	got := PackString("")
	want := []byte{0x00}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestPackSimpleString(t *testing.T) {
	got := PackString("foo")
	want := []byte{'f', 'o', 'o', 0x00}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestPackSmallPositiveInts(t *testing.T) {
	got := PackInt(1)
	want := []byte{0x01}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestPackMultiBytePositiveInts(t *testing.T) {
	got := PackInt(63)
	want := []byte{0x3F}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = PackInt(64)
	want = []byte{0x80, 0x01}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = PackInt(65)
	want = []byte{0x81, 0x01}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestPackSmallNegativeInts(t *testing.T) {
	got := PackInt(-1)
	want := []byte{0x40}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = PackInt(-2)
	want = []byte{0x41}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestPackMultiByteNegativeInts(t *testing.T) {
	got := PackInt(-63)
	want := []byte{0x7E}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = PackInt(-64)
	want = []byte{0x7F}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = PackInt(-65)
	want = []byte{0xC0, 0x01}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

// unpack

func TestUnpackSmallPositiveInts(t *testing.T) {

	require := require.New(t)
	want := 1
	got, err := UnpackInt([]byte{0x01})
	require.NoError(err)
	require.Equal(want, got)

	want = 2
	got, err = UnpackInt([]byte{0x02})
	require.NoError(err)
	require.Equal(want, got)

	want = 3
	got, err = UnpackInt([]byte{0x03})
	require.NoError(err)
	require.Equal(want, got)

}

func TestUnpackMultiBytePositiveInts(t *testing.T) {
	require := require.New(t)
	want := 63
	got, err := UnpackInt([]byte{0x3f})
	require.NoError(err)
	require.Equal(want, got)

	want = 64
	got, err = UnpackInt([]byte{0x80, 0x01})
	require.NoError(err)
	require.Equal(want, got)

	want = 65
	got, err = UnpackInt([]byte{0x81, 0x01})
	require.NoError(err)
	require.Equal(want, got)
}
