package submatch

import "regexp"

/* FindSubmatchSpecial
 */
func FindSubmatchSpecial(string string, regex regexp.Regexp) []Submatch {
	result := []Submatch{}

	stringSubmatchIndex := regex.FindStringSubmatchIndex(string)

	for i := 0; i < len(stringSubmatchIndex); i = i + 2 {
		matchString := string[stringSubmatchIndex[i]:stringSubmatchIndex[i+1]]
		start := stringSubmatchIndex[i] + 1
		end := max(stringSubmatchIndex[i]+len(matchString), start)

		submatch := Submatch{
			Submatch: matchString,
			Start:    start,
			End:      end,
		}
		result = append(result, submatch)
	}

	return result
} //

/*

What I want is for the

*/
