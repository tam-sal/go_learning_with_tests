package iteration

func Repeat(ch string, count int) string {
	var res string
	for i := 0; i < count; i++ {
		res += ch
	}
	return res
}
