package encodetext

// Run function encodes text based on the specified key
func Run(in string, keyFrom string, keyTo string) (string, error) {

	runesFrom := []rune(keyFrom)
	runesTo := []rune(keyTo)
	runesIn := []rune(in)

	// create key map
	m := make(map[rune]rune)

	for position, vRune := range runesFrom {
		if len(runesTo) <= position {
			break
		}
		m[vRune] = runesTo[position]
	}

	// make conversion
	runesOut := runesIn
	for position, vRune := range runesIn {
		if val, ok := m[vRune]; ok {
			runesOut[position] = val
		}
	}

	return string(runesOut), nil
}
