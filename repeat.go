package integers

func Repeat(character string, time int) string {
	var repeated string
	for i := 0; i < time; i++ {
		repeated += character
	}
	return repeated

	// return strings.Repeat(character, time)
}
