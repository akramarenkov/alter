package strings

import (
	"strings"
	"unicode/utf8"
)

// Prepares slice of type []string with specified length.
//
// Typical purpose - working with a reusable buffer and/or limit the length of the
// output slice.
type Preparer func(length int) []string

// Splits input slice to subslices separated by separator.
//
// If a preparer function is specified, the slice prepared by it will be used as the
// output slice. Otherwise, an individual slice will be created for each Split function
// call. The slice returned by the preparer function may be shorter than the requested
// length, in which case the splitting of the input slice will be limited by the length
// of the output slice.
func Split(input, separator string, preparer ...Preparer) []string {
	if len(separator) == 0 {
		return splitByUTF8(input, preparer)
	}

	separatorsQuantity := strings.Count(input, separator)

	// +1 for the remainder of slice following by the last separator
	output := prepare(separatorsQuantity+1, preparer)

	if len(output) == 0 {
		// []string{} is used instead of []string(nil) for compatibility with the
		// standard library
		return []string{}
	}

	for id := 0; ; id++ {
		// in case of preparer returned output slice of length less than requested
		if id >= len(output)-1 {
			output[id] = input
			return output
		}

		match := strings.Index(input, separator)
		if match == -1 {
			output[id] = input

			// id+1 in case of preparer returned output slice of length more than
			// requested
			return output[:id+1]
		}

		output[id] = input[:match]
		input = input[match+len(separator):]
	}
}

// Invalid UTF-8 sequences are splitted to separated bytes.
func splitByUTF8(input string, preparers []Preparer) []string {
	output := prepare(utf8.RuneCountInString(input), preparers)

	if len(output) == 0 {
		// []string{} is used instead of []string(nil) for compatibility with the
		// standard library
		return []string{}
	}

	id := 0

	for ; len(input) != 0; id++ {
		// in case of preparer returned output slice of length less than requested
		if id >= len(output)-1 {
			output[id] = input
			return output
		}

		_, size := utf8.DecodeRuneInString(input)

		output[id] = input[:size]
		input = input[size:]
	}

	return output[:id]
}

func prepare(length int, preparers []Preparer) []string {
	for _, preparer := range preparers {
		if preparer != nil {
			return preparer(length)
		}
	}

	return make([]string, length)
}
