package bytes

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSplitBase(t *testing.T) {
	dataset := []struct {
		input     []byte
		separator []byte
		expected  [][]byte
	}{
		{
			[]byte("1 2 3 4 5"),
			[]byte(" "),
			[][]byte{
				[]byte("1"),
				[]byte("2"),
				[]byte("3"),
				[]byte("4"),
				[]byte("5"),
			},
		},
		{
			[]byte("1 2 3 4 5 "),
			[]byte(" "),
			[][]byte{
				[]byte("1"),
				[]byte("2"),
				[]byte("3"),
				[]byte("4"),
				[]byte("5"),
				{},
			},
		},
		{
			[]byte("1 2 3  4 5 "),
			[]byte(" "),
			[][]byte{
				[]byte("1"),
				[]byte("2"),
				[]byte("3"),
				{},
				[]byte("4"),
				[]byte("5"),
				{},
			},
		},
	}

	for _, item := range dataset {
		require.Equal(
			t,
			item.expected,
			Split(item.input, item.separator),
		)

		require.Equal(
			t,
			bytes.Split(item.input, item.separator),
			Split(item.input, item.separator),
		)
	}
}

func TestSplitPreparer(t *testing.T) {
	input := []byte("1 2 3 4 5")
	separator := []byte(" ")

	dataset := []struct {
		preparer Preparer
		expected [][]byte
	}{
		{
			func(int) [][]byte { return make([][]byte, 0) },
			[][]byte(nil),
		},
		{
			func(int) [][]byte { return make([][]byte, 1) },
			[][]byte{
				[]byte("1 2 3 4 5"),
			},
		},
		{
			func(int) [][]byte { return make([][]byte, 2) },
			[][]byte{
				[]byte("1"),
				[]byte("2 3 4 5"),
			},
		},
		{
			func(int) [][]byte { return make([][]byte, 3) },
			[][]byte{
				[]byte("1"),
				[]byte("2"),
				[]byte("3 4 5"),
			},
		},
		{
			func(int) [][]byte { return make([][]byte, 4) },
			[][]byte{
				[]byte("1"),
				[]byte("2"),
				[]byte("3"),
				[]byte("4 5"),
			},
		},
		{
			func(int) [][]byte { return make([][]byte, 5) },
			[][]byte{
				[]byte("1"),
				[]byte("2"),
				[]byte("3"),
				[]byte("4"),
				[]byte("5"),
			},
		},
		{
			func(int) [][]byte { return make([][]byte, 6) },
			[][]byte{
				[]byte("1"),
				[]byte("2"),
				[]byte("3"),
				[]byte("4"),
				[]byte("5"),
			},
		},
		{
			func(int) [][]byte { return make([][]byte, 7) },
			[][]byte{
				[]byte("1"),
				[]byte("2"),
				[]byte("3"),
				[]byte("4"),
				[]byte("5"),
			},
		},
		{
			func(int) [][]byte { return make([][]byte, 8) },
			[][]byte{
				[]byte("1"),
				[]byte("2"),
				[]byte("3"),
				[]byte("4"),
				[]byte("5"),
			},
		},
	}

	for _, item := range dataset {
		require.Equal(
			t,
			item.expected,
			Split(input, separator, item.preparer),
		)
	}
}

func TestSplitByUTF8Base(t *testing.T) {
	input := []byte("Hello, 世界 Hello, 世界")

	require.Equal(
		t,
		[][]byte{
			[]byte("H"),
			[]byte("e"),
			[]byte("l"),
			[]byte("l"),
			[]byte("o"),
			[]byte(","),
			[]byte(" "),
			[]byte("世"),
			[]byte("界"),
			[]byte(" "),
			[]byte("H"),
			[]byte("e"),
			[]byte("l"),
			[]byte("l"),
			[]byte("o"),
			[]byte(","),
			[]byte(" "),
			[]byte("世"),
			[]byte("界"),
		},
		Split(input, nil),
	)

	require.Equal(
		t,
		bytes.Split(input, nil),
		Split(input, nil),
	)
}

func TestSplitByUTF8Preparer(t *testing.T) {
	input := []byte("Hello, 世界")

	dataset := []struct {
		preparer Preparer
		expected [][]byte
	}{
		{
			func(int) [][]byte { return make([][]byte, 0) },
			[][]byte{},
		},
		{
			func(int) [][]byte { return make([][]byte, 1) },
			[][]byte{
				[]byte("Hello, 世界"),
			},
		},
		{
			func(int) [][]byte { return make([][]byte, 2) },
			[][]byte{
				[]byte("H"),
				[]byte("ello, 世界"),
			},
		},
		{
			func(int) [][]byte { return make([][]byte, 3) },
			[][]byte{
				[]byte("H"),
				[]byte("e"),
				[]byte("llo, 世界"),
			},
		},
		{
			func(int) [][]byte { return make([][]byte, 4) },
			[][]byte{
				[]byte("H"),
				[]byte("e"),
				[]byte("l"),
				[]byte("lo, 世界"),
			},
		},
		{
			func(int) [][]byte { return make([][]byte, 5) },
			[][]byte{
				[]byte("H"),
				[]byte("e"),
				[]byte("l"),
				[]byte("l"),
				[]byte("o, 世界"),
			},
		},
		{
			func(int) [][]byte { return make([][]byte, 6) },
			[][]byte{
				[]byte("H"),
				[]byte("e"),
				[]byte("l"),
				[]byte("l"),
				[]byte("o"),
				[]byte(", 世界"),
			},
		},
		{
			func(int) [][]byte { return make([][]byte, 7) },
			[][]byte{
				[]byte("H"),
				[]byte("e"),
				[]byte("l"),
				[]byte("l"),
				[]byte("o"),
				[]byte(","),
				[]byte(" 世界"),
			},
		},
		{
			func(int) [][]byte { return make([][]byte, 8) },
			[][]byte{
				[]byte("H"),
				[]byte("e"),
				[]byte("l"),
				[]byte("l"),
				[]byte("o"),
				[]byte(","),
				[]byte(" "),
				[]byte("世界"),
			},
		},
		{
			func(int) [][]byte { return make([][]byte, 9) },
			[][]byte{
				[]byte("H"),
				[]byte("e"),
				[]byte("l"),
				[]byte("l"),
				[]byte("o"),
				[]byte(","),
				[]byte(" "),
				[]byte("世"),
				[]byte("界"),
			},
		},
		{
			func(int) [][]byte { return make([][]byte, 10) },
			[][]byte{
				[]byte("H"),
				[]byte("e"),
				[]byte("l"),
				[]byte("l"),
				[]byte("o"),
				[]byte(","),
				[]byte(" "),
				[]byte("世"),
				[]byte("界"),
			},
		},
		{
			func(int) [][]byte { return make([][]byte, 11) },
			[][]byte{
				[]byte("H"),
				[]byte("e"),
				[]byte("l"),
				[]byte("l"),
				[]byte("o"),
				[]byte(","),
				[]byte(" "),
				[]byte("世"),
				[]byte("界"),
			},
		},
		{
			func(int) [][]byte { return make([][]byte, 12) },
			[][]byte{
				[]byte("H"),
				[]byte("e"),
				[]byte("l"),
				[]byte("l"),
				[]byte("o"),
				[]byte(","),
				[]byte(" "),
				[]byte("世"),
				[]byte("界"),
			},
		},
	}

	for _, item := range dataset {
		require.Equal(
			t,
			item.expected,
			Split(input, nil, item.preparer),
		)
	}
}

func TestSplitDiff(t *testing.T) {
	dataset := []struct {
		input     []byte
		separator []byte
	}{
		{
			[]byte(nil),
			[]byte(nil),
		},
		{
			[]byte(""),
			[]byte(nil),
		},
		{
			[]byte(nil),
			[]byte(""),
		},
		{
			[]byte(""),
			[]byte(""),
		},
		{
			[]byte(""),
			[]byte(" "),
		},
		{
			[]byte(" "),
			[]byte(""),
		},
		{
			[]byte(" "),
			[]byte(" "),
		},
		{
			[]byte("\xf00"),
			[]byte(""),
		},
	}

	for _, item := range dataset {
		require.Equal(
			t,
			bytes.Split(item.input, item.separator),
			Split(item.input, item.separator),
		)
	}
}

func FuzzSplit(f *testing.F) {
	f.Add([]byte("1 2 3 4 5 "), []byte(" "))
	f.Add([]byte("Hello, 世界 Hello, 世界"), []byte(""))

	f.Fuzz(
		func(t *testing.T, input []byte, separator []byte) {
			require.Equal(t, bytes.Split(input, separator), Split(input, separator))
		},
	)
}

func BenchmarkSplitStd(b *testing.B) {
	input := []byte("1 2 3 4 5")
	separator := []byte(" ")
	expected := [][]byte{
		[]byte("1"),
		[]byte("2"),
		[]byte("3"),
		[]byte("4"),
		[]byte("5"),
	}

	var splitted [][]byte

	for range b.N {
		splitted = bytes.Split(input, separator)
	}

	b.StopTimer()

	require.Equal(b, expected, splitted)
}

func BenchmarkSplit(b *testing.B) {
	input := []byte("1 2 3 4 5")
	separator := []byte(" ")
	expected := [][]byte{
		[]byte("1"),
		[]byte("2"),
		[]byte("3"),
		[]byte("4"),
		[]byte("5"),
	}

	var splitted [][]byte

	for range b.N {
		splitted = Split(input, separator)
	}

	b.StopTimer()

	require.Equal(b, expected, splitted)
}

func BenchmarkSplitPreparer(b *testing.B) {
	input := []byte("1 2 3 4 5")
	separator := []byte(" ")
	expected := [][]byte{
		[]byte("1"),
		[]byte("2"),
		[]byte("3"),
		[]byte("4"),
		[]byte("5"),
	}

	buffer := make([][]byte, len(expected))

	preparer := func(int) [][]byte {
		return buffer
	}

	var splitted [][]byte

	for range b.N {
		splitted = Split(input, separator, preparer)
	}

	b.StopTimer()

	require.Equal(b, expected, splitted)
}
