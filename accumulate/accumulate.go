package accumulate

const testVersion = 1

func Accumulate(s []string, fn func(string) string) []string {
	collection := s
	for i, element := range s {
		collection[i] = fn(element)
	}
	return collection
}
