package fizzbuzz

import "testing"

func TestFizzBuzz (t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		} 
	}

	t.Run("Test FizzBuzz 1 Return 1", func(t *testing.T) {
		got := FizzBuzz(1)
		want := "1"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Test FizzBuzz 2 Return 2", func(t *testing.T) {
		got := FizzBuzz(2)
		want := "2"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Test FizzBuzz 3 Return Fizz", func(t *testing.T) {
		got := FizzBuzz(3)
		want := "Fizz"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Test FizzBuzz 5 Return Buzz", func(t *testing.T) {
		got := FizzBuzz(5)
		want := "Buzz"
		assertCorrectMessage(t, got, want)
	})
	
	t.Run("Test FizzBuzz 15 Return FizzBuzz", func(t *testing.T) {
		got := FizzBuzz(15)
		want := "FizzBuzz"
		assertCorrectMessage(t, got, want)
	})	
}