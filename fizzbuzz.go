package fizzbuzz

import "strconv"

func FizzBuzz(got int) string {

	var want string = ""

	if got%3 == 0 {
		want = "Fizz"
	}

	if got%5 == 0 {
		want += "Buzz"
	}

	if want == "" {
		want = strconv.Itoa(got)
	}

	return want
}