package fizzbuzz

import "testing"

func assertCorrectMessage (t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestFizzBuzz1Return1(t *testing.T) {
	got := FizzBuzz(1)
	want := "1"
	assertCorrectMessage(t, got, want)
}

func TestFizzBuzz2Return2(t *testing.T) {
	got := FizzBuzz(2)
	want := "2"
	assertCorrectMessage(t, got, want)
}

func TestFizzBuzz3ReturnFizz(t *testing.T) {
	got := FizzBuzz(3)
	want := "Fizz"
	assertCorrectMessage(t, got, want)
}

func TestFizzBuzz5ReturnBuzz(t *testing.T) {
	got := FizzBuzz(5)
	want := "Buzz"
	assertCorrectMessage(t, got, want)
}

func TestFizzBuzz15ReturnFizzBuzz(t *testing.T) {
	got := FizzBuzz(15)
	want := "FizzBuzz"
	assertCorrectMessage(t, got, want)
}