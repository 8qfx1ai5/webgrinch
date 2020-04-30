package encodetext

var m map[rune]rune

// Run function encodes content HTML text based on the specified key and add css classes
func Run(in string, keyFrom string, keyTo string) (out string, err error) {

	m = make(map[rune]rune)

	runesFrom := []rune(keyFrom)
	runesTo := []rune(keyTo)
	runesIn := []rune(in)

	// create key map
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
