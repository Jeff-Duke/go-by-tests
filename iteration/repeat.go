package iteration

// Repeat takes a char and repeats it
func Repeat(char string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated = repeated + char
	}
	return repeated
}
