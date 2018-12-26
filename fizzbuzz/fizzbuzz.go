package fizzbuzz

import "strconv"

func FizzBuzz(got int) string {

	var want string = ""

	switch {
	case (got%3 == 0) && (got%5 == 0):
		want = "FizzBuzz"
	case got%3 == 0:
		want = "Fizz"
	case got%5 == 0:
		want = "Buzz"
	default:
		want = strconv.Itoa(got)
	}

	return want
}