package submatch

import "regexp"

/* FindSubmatchSpecial
 */
func FindSubmatchSpecial(string string, regex regexp.Regexp) []Submatch {
	result := []Submatch{}

	stringSubmatchIndex := regex.FindStringSubmatchIndex(string)
	curPos := 1

	var submatch Submatch

	for i := 0; i < len(stringSubmatchIndex); i = i + 2 {
		if stringSubmatchIndex[i] == -1 {
			submatch = Submatch{
				String: "",
				Start:  curPos,
				End:    curPos,
			}
		} else {
			matchString := string[stringSubmatchIndex[i]:stringSubmatchIndex[i+1]]
			start := stringSubmatchIndex[i] + 1
			end := max(stringSubmatchIndex[i]+len(matchString), start)

			submatch = Submatch{
				String: matchString,
				Start:  start,
				End:    end,
			}

		}
		curPos = submatch.End
		result = append(result, submatch)
	}

	return result
} //

/*

What I want is for the

*/
