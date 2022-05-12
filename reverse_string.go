package reverse_string

import "strings"

func ReverseString(input string) (output string) {
	// solution goes here
	inputSls := strings.Split(input, "")
	outputSls := make([]string, 0, len(inputSls))
	for i := len(inputSls) - 1; i >= 0; i-- {
		outputSls = append(outputSls, inputSls[i])
	}
	output = strings.Join(outputSls, "")
	return output
}
