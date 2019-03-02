package main

func main() {
	letterCombinations("23")
}
func letterCombinations(digits string) []string {
	//edge case
	if len(digits) == 0 {
		return []string{}
	}

	//dictionary for phone number to letter mapping
	var hash = map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	//initialize
	result := []string{""}
	// loop input
	for _, digitAscii := range digits {
		digit := string(digitAscii)
		// get letters based on the digit
		// check if digit exists (for verifying the digit is not "1")
		letters, success := hash[digit]
		if success {
			var combo []string
			for _, prefix := range result {
				for _, letter := range letters {
					// keep building combination list based on prefix calculated so far and appending new character
					combo = append(combo, string(prefix)+string(letter))
				}
			}
			result = combo
		}
	}
	return result
}
