package bytes

import (
	"bytes"
	"unicode/utf8"
)

// Prepares slice of type [][]byte with specified length.
//
// Typical purpose - working with a reusable buffer and/or limit the length of the
// output slice.
type Preparer func(length int) [][]byte

// Splits input slice to subslices separated by separator.
//
// If a preparer function is specified, the slice prepared by it will be used as the
// output slice. Otherwise, an individual slice will be created for each Split function
// call. The slice returned by the preparer function may be shorter than the requested
// length, in which case the splitting of the input slice will be limited by the length
// of the output slice.
func Split(input, separator []byte, preparer ...Preparer) [][]byte {
	if len(separator) == 0 {
		return splitByUTF8(input, preparer)
	}

	separatorsQuantity := bytes.Count(input, separator)

	// +1 for the remainder of slice following by the last separator
	output := prepare(separatorsQuantity+1, preparer)

	if len(output) == 0 {
		return nil
	}

	for id := 0; ; id++ {
		// in case of preparer returned output slice of length less than requested
		if id >= len(output)-1 {
			output[id] = input
			return output
		}

		match := bytes.Index(input, separator)
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
func splitByUTF8(input []byte, preparers []Preparer) [][]byte {
	output := prepare(utf8.RuneCount(input), preparers)

	if len(output) == 0 {
		// [][]byte{} is used instead of [][]byte(nil) for compatibility with the
		// standard library
		return [][]byte{}
	}

	id := 0

	for ; len(input) != 0; id++ {
		// in case of preparer returned output slice of length less than requested
		if id >= len(output)-1 {
			output[id] = input
			return output
		}

		_, size := utf8.DecodeRune(input)

		output[id] = input[:size]
		input = input[size:]
	}

	return output[:id]
}

func prepare(length int, preparers []Preparer) [][]byte {
	for _, preparer := range preparers {
		if preparer != nil {
			return preparer(length)
		}
	}

	return make([][]byte, length)
}
