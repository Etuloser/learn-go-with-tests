package iteration

const repeatConut = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatConut; i++ {
		repeated += character
	}
	return repeated
}
