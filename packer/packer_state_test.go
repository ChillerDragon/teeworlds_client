package packer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// rest

func TestUnpackRest(t *testing.T) {
	u := NewUnpacker([]byte{0x01, 0xff, 0xaa})

	{
		got, err := u.NextInt()
		require.NoError(t, err)
		require.Equal(t, 1, got)
	}

	{
		want := []byte{0xff, 0xaa}
		got := u.Bytes()
		require.Equal(t, want, got)
	}
}

// client info

func TestUnpackClientInfo(t *testing.T) {
	require := require.New(t)
	u := NewUnpacker([]byte{
		0x24, 0x00, 0x01, 0x00, 0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x00,
		0x00, 0x40, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x73, 0x77, 0x61, 0x72,
		0x64, 0x00, 0x64, 0x75, 0x6f, 0x64, 0x6f, 0x6e, 0x6e, 0x79, 0x00,
		0x00, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x00, 0x73,
		0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x00, 0x73, 0x74, 0x61,
		0x6e, 0x64, 0x61, 0x72, 0x64, 0x00, 0x01, 0x01, 0x00, 0x00, 0x00,
		0x00, 0x80, 0xfc, 0xaf, 0x05, 0xeb, 0x83, 0xd0, 0x0a, 0x80, 0xfe,
		0x07, 0x80, 0xfe, 0x07, 0x80, 0xfe, 0x07, 0x80, 0xfe, 0x07, 0x00,
	})

	{
		// message id
		want := 36
		got, err := u.NextInt()
		require.NoError(err)
		require.Equal(want, got)

		// client id
		want = 0
		got, err = u.NextInt()
		require.NoError(err)
		require.Equal(want, got)

		_, err = u.NextBool() // Local bool
		require.NoError(err)
		_, err = u.NextInt() // Team int
		require.NoError(err)
	}

	{
		// name
		want := "gopher"
		got, err := u.NextString()
		require.NoError(err)
		require.Equal(want, got)

		// clan
		want = ""
		got, err = u.NextString()
		require.NoError(err)
		require.Equal(want, got)

	}

	{
		// country
		want := -1
		got, err := u.NextInt()
		require.NoError(err)
		require.Equal(want, got)
	}

	{
		// body
		want := "greensward"
		got, err := u.NextString()
		require.NoError(err)
		require.Equal(want, got)
	}
}

// unpack with state

func TestUnpackSimpleInts(t *testing.T) {
	require := require.New(t)
	u := NewUnpacker([]byte{0x01, 0x02, 0x03, 0x0f})

	want := 1
	got, err := u.NextInt()
	require.NoError(err)
	require.Equal(want, got)

	want = 2
	got, err = u.NextInt()
	require.NoError(err)
	require.Equal(want, got)

	want = 3
	got, err = u.NextInt()
	require.NoError(err)
	require.Equal(want, got)

	want = 15
	got, err = u.NextInt()
	require.NoError(err)
	require.Equal(want, got)
}

func TestUnpackString(t *testing.T) {
	require := require.New(t)
	u := NewUnpacker([]byte{'f', 'o', 'o', 0x00})

	want := "foo"
	got, err := u.NextString()
	require.NoError(err)
	require.Equal(want, got)
}

func TestUnpackTwoStrings(t *testing.T) {
	require := require.New(t)
	u := NewUnpacker([]byte{'f', 'o', 'o', 0x00, 'b', 'a', 'r', 0x00})

	want := "foo"
	got, err := u.NextString()
	require.NoError(err)
	require.Equal(want, got)

	want = "bar"
	got, err = u.NextString()
	require.NoError(err)
	require.Equal(want, got)
}

func TestUnpackMixed(t *testing.T) {
	require := require.New(t)
	u := NewUnpacker([]byte{0x0F, 0x0F, 'f', 'o', 'o', 0x00, 'b', 'a', 'r', 0x00, 0x01})

	// ints
	{
		want := 15
		got, err := u.NextInt()
		require.NoError(err)
		require.Equal(want, got)

		want = 15
		got, err = u.NextInt()
		require.NoError(err)
		require.Equal(want, got)
	}

	// strings
	{
		want := "foo"
		got, err := u.NextString()
		require.NoError(err)
		require.Equal(want, got)

		want = "bar"
		got, err = u.NextString()
		require.NoError(err)
		require.Equal(want, got)
	}

	// ints
	{
		want := 1
		got, err := u.NextInt()
		require.NoError(err)
		require.Equal(want, got)
	}
}
