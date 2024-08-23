package strings

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSplitBase(t *testing.T) {
	dataset := []struct {
		input     string
		separator string
		expected  []string
	}{
		{
			"1 2 3 4 5",
			" ",
			[]string{
				"1",
				"2",
				"3",
				"4",
				"5",
			},
		},
		{
			"1 2 3 4 5 ",
			" ",
			[]string{
				"1",
				"2",
				"3",
				"4",
				"5",
				"",
			},
		},
		{
			"1 2 3  4 5 ",
			" ",
			[]string{
				"1",
				"2",
				"3",
				"",
				"4",
				"5",
				"",
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
			strings.Split(item.input, item.separator),
			Split(item.input, item.separator),
		)
	}
}

func TestSplitPreparer(t *testing.T) {
	const (
		input     = "1 2 3 4 5"
		separator = " "
	)

	dataset := []struct {
		preparer Preparer
		expected []string
	}{
		{
			func(int) []string { return make([]string, 0) },
			[]string{},
		},
		{
			func(int) []string { return make([]string, 1) },
			[]string{
				"1 2 3 4 5",
			},
		},
		{
			func(int) []string { return make([]string, 2) },
			[]string{
				"1",
				"2 3 4 5",
			},
		},
		{
			func(int) []string { return make([]string, 3) },
			[]string{
				"1",
				"2",
				"3 4 5",
			},
		},
		{
			func(int) []string { return make([]string, 4) },
			[]string{
				"1",
				"2",
				"3",
				"4 5",
			},
		},
		{
			func(int) []string { return make([]string, 5) },
			[]string{
				"1",
				"2",
				"3",
				"4",
				"5",
			},
		},
		{
			func(int) []string { return make([]string, 6) },
			[]string{
				"1",
				"2",
				"3",
				"4",
				"5",
			},
		},
		{
			func(int) []string { return make([]string, 7) },
			[]string{
				"1",
				"2",
				"3",
				"4",
				"5",
			},
		},
		{
			func(int) []string { return make([]string, 8) },
			[]string{
				"1",
				"2",
				"3",
				"4",
				"5",
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
	input := "Hello, 世界 Hello, 世界"

	require.Equal(
		t,
		[]string{
			"H",
			"e",
			"l",
			"l",
			"o",
			",",
			" ",
			"世",
			"界",
			" ",
			"H",
			"e",
			"l",
			"l",
			"o",
			",",
			" ",
			"世",
			"界",
		},
		Split(input, ""),
	)

	require.Equal(
		t,
		strings.Split(input, ""),
		Split(input, ""),
	)
}

func TestSplitByUTF8Preparer(t *testing.T) {
	input := "Hello, 世界"

	dataset := []struct {
		preparer Preparer
		expected []string
	}{
		{
			func(int) []string { return make([]string, 0) },
			[]string{},
		},
		{
			func(int) []string { return make([]string, 1) },
			[]string{
				"Hello, 世界",
			},
		},
		{
			func(int) []string { return make([]string, 2) },
			[]string{
				"H",
				"ello, 世界",
			},
		},
		{
			func(int) []string { return make([]string, 3) },
			[]string{
				"H",
				"e",
				"llo, 世界",
			},
		},
		{
			func(int) []string { return make([]string, 4) },
			[]string{
				"H",
				"e",
				"l",
				"lo, 世界",
			},
		},
		{
			func(int) []string { return make([]string, 5) },
			[]string{
				"H",
				"e",
				"l",
				"l",
				"o, 世界",
			},
		},
		{
			func(int) []string { return make([]string, 6) },
			[]string{
				"H",
				"e",
				"l",
				"l",
				"o",
				", 世界",
			},
		},
		{
			func(int) []string { return make([]string, 7) },
			[]string{
				"H",
				"e",
				"l",
				"l",
				"o",
				",",
				" 世界",
			},
		},
		{
			func(int) []string { return make([]string, 8) },
			[]string{
				"H",
				"e",
				"l",
				"l",
				"o",
				",",
				" ",
				"世界",
			},
		},
		{
			func(int) []string { return make([]string, 9) },
			[]string{
				"H",
				"e",
				"l",
				"l",
				"o",
				",",
				" ",
				"世",
				"界",
			},
		},
		{
			func(int) []string { return make([]string, 10) },
			[]string{
				"H",
				"e",
				"l",
				"l",
				"o",
				",",
				" ",
				"世",
				"界",
			},
		},
		{
			func(int) []string { return make([]string, 11) },
			[]string{
				"H",
				"e",
				"l",
				"l",
				"o",
				",",
				" ",
				"世",
				"界",
			},
		},
		{
			func(int) []string { return make([]string, 12) },
			[]string{
				"H",
				"e",
				"l",
				"l",
				"o",
				",",
				" ",
				"世",
				"界",
			},
		},
	}

	for _, item := range dataset {
		require.Equal(
			t,
			item.expected,
			Split(input, "", item.preparer),
		)
	}
}

func TestSplitDiff(t *testing.T) {
	dataset := []struct {
		input     string
		separator string
	}{
		{
			"",
			"",
		},
		{
			"",
			"",
		},
		{
			"",
			"",
		},
		{
			"",
			"",
		},
		{
			"",
			" ",
		},
		{
			" ",
			"",
		},
		{
			" ",
			" ",
		},
		{
			"\xf00",
			"",
		},
	}

	for _, item := range dataset {
		require.Equal(
			t,
			strings.Split(item.input, item.separator),
			Split(item.input, item.separator),
		)
	}
}

func FuzzSplit(f *testing.F) {
	f.Add("1 2 3 4 5 ", " ")
	f.Add("Hello, 世界 Hello, 世界", "")

	f.Fuzz(
		func(t *testing.T, input string, separator string) {
			require.Equal(t, strings.Split(input, separator), Split(input, separator))
		},
	)
}

func BenchmarkSplitStd(b *testing.B) {
	const (
		input     = "1 2 3 4 5"
		separator = " "
	)

	expected := []string{
		"1",
		"2",
		"3",
		"4",
		"5",
	}

	var splitted []string

	for range b.N {
		splitted = strings.Split(input, separator)
	}

	b.StopTimer()

	require.Equal(b, expected, splitted)
}

func BenchmarkSplit(b *testing.B) {
	input := "1 2 3 4 5"
	separator := " "
	expected := []string{
		"1",
		"2",
		"3",
		"4",
		"5",
	}

	var splitted []string

	for range b.N {
		splitted = Split(input, separator)
	}

	b.StopTimer()

	require.Equal(b, expected, splitted)
}

func BenchmarkSplitPreparer(b *testing.B) {
	input := "1 2 3 4 5"
	separator := " "
	expected := []string{
		"1",
		"2",
		"3",
		"4",
		"5",
	}

	buffer := make([]string, len(expected))

	preparer := func(int) []string {
		return buffer
	}

	var splitted []string

	for range b.N {
		splitted = Split(input, separator, preparer)
	}

	b.StopTimer()

	require.Equal(b, expected, splitted)
}
